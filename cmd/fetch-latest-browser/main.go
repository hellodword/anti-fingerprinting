package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/dop251/goja"
	"github.com/hellodword/anti-fingerprinting/internal/common"
)

type Commit struct {
	Sha string `json:"sha"`
}

type Browser struct {
	Version     string             `json:"version"`
	URL         string             `json:"url,omitempty"`
	Hash        string             `json:"hash,omitempty"`
	BrowserType common.BrowserType `json:"browser,omitempty"`
}

type ScoopInstaller struct {
	Browser
	Architecture *struct {
		X64 struct {
			URL  string `json:"url"`
			Hash string `json:"hash"`
		} `json:"64bit,omitempty"`
	} `json:"architecture"`
}

func strBetween(src, start, end string) string {
	startIndex := strings.Index(src, start)
	if startIndex == -1 {
		return ""
	}
	endIndex := strings.Index(src[startIndex+len(start):], end)
	if endIndex == -1 {
		return ""
	}
	return src[startIndex+len(start) : startIndex+len(start)+endIndex]
}

func fetch(u string, v interface{}) ([]byte, error) {
	r, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	log.Println(u, r.Status)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if v != nil {
		return b, json.Unmarshal(b, v)
		// return json.NewDecoder(r.Body).Decode(v)
	} else {
		return b, nil
	}
}

func post(u string, body []byte) ([]byte, error) {
	r, err := http.Post(u, "application/x-www-form-urlencoded", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	log.Println(u, r.Status)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func match1st(re *regexp.Regexp, b []byte) string {
	mm := re.FindStringSubmatch(string(b))
	if len(mm) == 2 {
		return mm[1]
	}
	return ""
}

type EdgeVersion struct {
	ChannelID    string `json:"channelId,omitempty"`
	MajorVersion string `json:"majorVersion,omitempty"`
	Releases     []struct {
		FullVersion string `json:"fullVersion,omitempty"`
		PolicyURL   string `json:"policyUrl,omitempty"`
		Platforms   []struct {
			PlatformID  string `json:"platformId,omitempty"`
			DownloadURL string `json:"downloadUrl,omitempty"`
			SizeInBytes int    `json:"sizeInBytes,omitempty"`
		} `json:"platforms,omitempty"`
	} `json:"releases,omitempty"`
}

func getEdgeVersions_1(nedge uint) ([]Browser, error) {
	b, err := fetch("https://www.microsoft.com/en-us/edge/business/download", nil)
	if err != nil {
		return nil, err
	}

	s := strBetween(string(b), "<script>window.__NUXT__=", "</script>")
	if s == "" {
		return nil, errors.New("fetch edge versions")
	}

	s = "var __NUXT__=" + s

	vm := goja.New()
	_, err = vm.RunString(s)
	if err != nil {
		return nil, err
	}

	res, err := vm.RunString("JSON.stringify(__NUXT__.fetch['block-enterprise-downloads:0'].majorReleases)")
	if err != nil {
		return nil, err
	}

	var edges []EdgeVersion
	err = json.Unmarshal([]byte(res.String()), &edges)
	if err != nil {
		return nil, err
	}

	var exists = map[string]struct{}{}
	var browsers []Browser

	for i := range edges {
		if len(browsers) >= int(nedge) {
			break
		}

		if edges[i].ChannelID != "stable" {
			continue
		}

		if len(edges[i].Releases) == 0 {
			continue
		}

		if edges[i].Releases[0].FullVersion == "" {
			continue
		}

		var edge = Browser{
			BrowserType: common.BrowserTypeEdge,
			Version:     edges[i].Releases[0].FullVersion,
		}

		for _, platform := range edges[i].Releases[0].Platforms {
			if platform.PlatformID == "windows-x64" {
				edge.URL = platform.DownloadURL
				break
			}
		}

		if edge.URL == "" {
			continue
		}

		if _, ok := exists[edge.Version]; !ok {
			exists[edge.Version] = struct{}{}
			browsers = append(browsers, edge)
		}

	}

	return browsers, nil
}

func getEdgeVersions_2(nedge uint) ([]Browser, error) {

	b, err := fetch(fmt.Sprintf(
		"https://www.catalog.update.microsoft.com/Search.aspx?q=%s",
		url.QueryEscape(`"Microsoft Edge-Stable Channel Version" x64`)), nil)
	if err != nil {
		return nil, err
	}

	var re = regexp.MustCompile(`goToDetails\("([a-f\d\-]{36})"\)`)

	mm := re.FindAllStringSubmatch(string(b), -1)

	var detailIDs []string

	for i := range mm {
		if len(mm[i]) == 2 && mm[i][1] != "" {
			detailIDs = append(detailIDs, mm[i][1])
		}
	}

	if len(detailIDs) == 0 {
		return nil, errors.New("edge search updates failed")
	}

	var exists = map[string]struct{}{}
	var browsers []Browser

	for _, detailID := range detailIDs {
		if len(browsers) >= int(nedge) {
			break
		}

		b, err := post("https://www.catalog.update.microsoft.com/DownloadDialog.aspx",
			[]byte(`updateIDs=%5B%7B%22size%22%3A0%2C%22languages%22%3A%22%22%2C%22uidInfo%22%3A%22`+detailID+`%22%2C%22updateID%22%3A%22`+detailID+`%22%7D%5D&updateIDsBlockedForImport=&wsusApiPresent=&contentImport=&sku=&serverName=&ssl=&portNumber=&version=`))
		if err != nil {
			return nil, err
		}

		// fmt.Println(string(b[bytes.Index(b, []byte("downloadInformation[0].enTitle ='")) : bytes.Index(b, []byte("downloadInformation[0].enTitle ='"))+1024]))

		var edge Browser
		edge.BrowserType = common.BrowserTypeEdge

		edge.Version = match1st(regexp.MustCompile(`downloadInformation\[0\]\.enTitle ='Microsoft Edge-Stable Channel Version \d+ Update for x64 based Editions \(Build ([\d\.]+)\)'`), b)
		if edge.Version == "" {
			continue
		}

		digest := match1st(regexp.MustCompile(`downloadInformation\[0\]\.files\[0\]\.digest = '([^\n']+)'`), b)
		if digest == "" {
			continue
		}

		bdigest, err := base64.StdEncoding.DecodeString(digest)
		if err != nil {
			return nil, err
		}

		edge.Hash = hex.EncodeToString(bdigest)

		edge.URL = match1st(regexp.MustCompile(`downloadInformation\[0\]\.files\[0\]\.url = '(https:[^\n']+\.cab)'`), b)
		if edge.URL == "" {
			continue
		}

		if _, ok := exists[edge.Version]; !ok {
			exists[edge.Version] = struct{}{}
			browsers = append(browsers, edge)
		}
	}

	return browsers, nil
}

func getChromeVersions_1(nchrome uint) ([]Browser, error) {
	var commits []Commit
	_, err := fetch("https://api.github.com/repos/ScoopInstaller/Extras/commits?path=/bucket/googlechrome.json&sha=master&per_page=100&page=1", &commits)
	if err != nil {
		return nil, err
	}

	var exists = map[string]struct{}{}
	var browsers []Browser

	for _, commit := range commits {
		if len(browsers) >= int(nchrome) {
			break
		}

		log.Println("sha", commit.Sha)

		var chrome ScoopInstaller
		_, err = fetch(fmt.Sprintf("https://raw.githubusercontent.com/ScoopInstaller/Extras/%s/bucket/googlechrome.json", commit.Sha), &chrome)
		if err != nil {
			return nil, err
		}

		if chrome.Version == "" || chrome.Architecture == nil || chrome.Architecture.X64.URL == "" || chrome.Architecture.X64.Hash == "" {
			continue
		}

		if _, ok := exists[chrome.Version]; !ok {
			chrome.URL = chrome.Architecture.X64.URL
			chrome.Hash = chrome.Architecture.X64.Hash
			chrome.Architecture = nil
			chrome.BrowserType = common.BrowserTypeChrome

			exists[chrome.Version] = struct{}{}
			browsers = append(browsers, chrome.Browser)
		}
	}

	return browsers, nil
}

func getFirefoxVersions_1(nfirefox uint) ([]Browser, error) {
	var commits []Commit
	_, err := fetch("https://api.github.com/repos/ScoopInstaller/Extras/commits?path=/bucket/firefox.json&sha=master&per_page=100&page=1", &commits)
	if err != nil {
		return nil, err
	}

	var exists = map[string]struct{}{}
	var browsers []Browser

	for _, commit := range commits {
		if len(browsers) >= int(nfirefox) {
			break
		}

		log.Println("sha", commit.Sha)

		var firefox ScoopInstaller
		_, err = fetch(fmt.Sprintf("https://raw.githubusercontent.com/ScoopInstaller/Extras/%s/bucket/firefox.json", commit.Sha), &firefox)
		if err != nil {
			return nil, err
		}

		if firefox.Version == "" || firefox.Architecture == nil || firefox.Architecture.X64.URL == "" || firefox.Architecture.X64.Hash == "" {
			continue
		}

		if _, ok := exists[firefox.Version]; !ok {
			firefox.URL = firefox.Architecture.X64.URL
			firefox.Hash = firefox.Architecture.X64.Hash
			firefox.Architecture = nil
			firefox.BrowserType = common.BrowserTypeFirefox

			exists[firefox.Version] = struct{}{}
			browsers = append(browsers, firefox.Browser)
		}
	}

	return browsers, nil
}

func main() {
	output := flag.String("o", "browser.json", "output")
	nchrome := flag.Uint("nchrome", 5, "num of chrome versions [0, 20]")
	nfirefox := flag.Uint("nfirefox", 3, "num of firefox versions [0, 10]")
	nedge := flag.Uint("nedge", 5, "num of edge versions [0, 20]")
	flag.Parse()

	if *nchrome > 20 {
		*nchrome = 5
	}
	if *nfirefox > 10 {
		*nfirefox = 3
	}
	if *nedge > 20 {
		*nedge = 5
	}

	var browsers = []Browser{}

	var err error

	{
		var edges = []Browser{}
		if *nedge > 0 {
			if *nedge > 5 {
				edges, err = getEdgeVersions_2(*nedge)
			} else {
				edges, err = getEdgeVersions_1(*nedge)
			}

			if err != nil {
				panic(err)
			}
		}

		browsers = append(browsers, edges...)
	}

	{
		var chromes = []Browser{}
		if *nchrome > 0 {
			chromes, err = getChromeVersions_1(*nchrome)
			if err != nil {
				panic(err)
			}
		}

		browsers = append(browsers, chromes...)
	}
	{
		var firefoxes = []Browser{}
		if *nchrome > 0 {
			firefoxes, err = getFirefoxVersions_1(*nfirefox)
			if err != nil {
				panic(err)
			}
		}

		browsers = append(browsers, firefoxes...)
	}

	{
		b, _ := json.Marshal(browsers)
		err = os.WriteFile(*output, b, 0644)
		if err != nil {
			panic(err)
		}
	}
}
