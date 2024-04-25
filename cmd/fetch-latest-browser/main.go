package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Commit struct {
	Sha string `json:"sha"`
}

type BrowserType string

const (
	BrowserTypeChrome BrowserType = "chrome"
)

type Browser struct {
	Version     string      `json:"version"`
	URL         string      `json:"url,omitempty"`
	Hash        string      `json:"hash,omitempty"`
	BrowserType BrowserType `json:"browser,omitempty"`
}

type Chrome struct {
	Browser
	Architecture *struct {
		X64 struct {
			URL  string `json:"url"`
			Hash string `json:"hash"`
		} `json:"64bit,omitempty"`
	} `json:"architecture"`
}

func fetch(u string, v interface{}) error {
	r, err := http.Get(u)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	log.Println(u, r.Status)

	d := json.NewDecoder(r.Body)

	return d.Decode(v)
}

func main() {
	output := flag.String("o", "browser.json", "output")
	num := flag.Uint("n", 5, "num [1, 20]")
	flag.Parse()

	if *num < 1 || *num > 20 {
		*num = 10
	}

	var err error
	var commits []Commit
	err = fetch("https://api.github.com/repos/ScoopInstaller/Extras/commits?path=/bucket/googlechrome.json&sha=master&per_page=100&page=1", &commits)
	if err != nil {
		panic(err)
	}

	if len(commits) == 0 {
		panic(len(commits))
	}

	var exists = map[string]struct{}{}
	var browsers []Browser

	for _, commit := range commits {
		log.Println("sha", commit.Sha)

		var chrome Chrome
		err = fetch(fmt.Sprintf("https://raw.githubusercontent.com/ScoopInstaller/Extras/%s/bucket/googlechrome.json", commit.Sha), &chrome)
		if err != nil {
			panic(err)
		}

		if chrome.Version == "" || chrome.Architecture == nil || chrome.Architecture.X64.URL == "" || chrome.Architecture.X64.Hash == "" {
			panic(commit.Sha)
		}

		if _, ok := exists[chrome.Version]; !ok {
			chrome.URL = chrome.Architecture.X64.URL
			chrome.Hash = chrome.Architecture.X64.Hash
			chrome.Architecture = nil
			chrome.BrowserType = BrowserTypeChrome

			exists[chrome.Version] = struct{}{}
			browsers = append(browsers, chrome.Browser)

			if len(browsers) >= int(*num) {
				break
			}
		}
	}

	b, _ := json.Marshal(browsers)
	err = os.WriteFile(*output, b, 0644)
	if err != nil {
		panic(err)
	}
}
