package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/hellodword/tls-fingerprinting/internal/common"
)

// https://stackoverflow.com/a/28323276
type arrayFlags []string

func (i *arrayFlags) String() string {
	return ""
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var myFlags arrayFlags
	flag.Var(&myFlags, "s", "collectorinfos to be compared")
	flag.Parse()

	if len(myFlags) != 2 {
		panic("2 collectorinfos required")
	}

	var err error
	var info1, info2 common.CollectedInfo
	err = json.Unmarshal([]byte(myFlags[0]), &info1)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(myFlags[1]), &info2)
	if err != nil {
		panic(err)
	}

	if info1.Equals(info2) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}

}
