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
	BrowserTypeChrome  BrowserType = "chrome"
	BrowserTypeFirefox BrowserType = "firefox"
)

type Browser struct {
	Version     string      `json:"version"`
	URL         string      `json:"url,omitempty"`
	Hash        string      `json:"hash,omitempty"`
	BrowserType BrowserType `json:"browser,omitempty"`
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
	nchrome := flag.Uint("nchrome", 5, "num of chrome versions [1, 20]")
	nfirefox := flag.Uint("nfirefox", 3, "num of firefox versions [1, 10]")
	flag.Parse()

	if *nchrome < 1 || *nchrome > 20 {
		*nchrome = 5
	}
	if *nfirefox < 1 || *nfirefox > 10 {
		*nfirefox = 3
	}

	var exists = map[string]struct{}{}
	var browsers []Browser

	var err error

	// Chrome
	{
		var commits []Commit
		var count = 0
		err = fetch("https://api.github.com/repos/ScoopInstaller/Extras/commits?path=/bucket/googlechrome.json&sha=master&per_page=100&page=1", &commits)
		if err != nil {
			panic(err)
		}

		if len(commits) == 0 {
			panic(len(commits))
		}

		for _, commit := range commits {
			log.Println("sha", commit.Sha)

			var chrome ScoopInstaller
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
				count++

				if count >= int(*nchrome) {
					break
				}
			}
		}
	}

	// Firefox
	{
		var commits []Commit
		var count = 0
		err = fetch("https://api.github.com/repos/ScoopInstaller/Extras/commits?path=/bucket/firefox.json&sha=master&per_page=100&page=1", &commits)
		if err != nil {
			panic(err)
		}

		if len(commits) == 0 {
			panic(len(commits))
		}

		for _, commit := range commits {
			log.Println("sha", commit.Sha)

			var firefox ScoopInstaller
			err = fetch(fmt.Sprintf("https://raw.githubusercontent.com/ScoopInstaller/Extras/%s/bucket/firefox.json", commit.Sha), &firefox)
			if err != nil {
				panic(err)
			}

			if firefox.Version == "" || firefox.Architecture == nil || firefox.Architecture.X64.URL == "" || firefox.Architecture.X64.Hash == "" {
				panic(commit.Sha)
			}

			if _, ok := exists[firefox.Version]; !ok {
				firefox.URL = firefox.Architecture.X64.URL
				firefox.Hash = firefox.Architecture.X64.Hash
				firefox.Architecture = nil
				firefox.BrowserType = BrowserTypeFirefox

				exists[firefox.Version] = struct{}{}
				browsers = append(browsers, firefox.Browser)
				count++

				if count >= int(*nfirefox) {
					break
				}
			}
		}
	}

	{
		b, _ := json.Marshal(browsers)
		err = os.WriteFile(*output, b, 0644)
		if err != nil {
			panic(err)
		}
	}
}
