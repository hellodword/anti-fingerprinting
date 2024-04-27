//go:build !patched
// +build !patched

package main

import (
	"log"

	"github.com/quic-go/quic-go"
)

func applyQuicPatch(quicConf *quic.Config) *quic.Config {
	log.Println("quic not patched")
	return quicConf
}
