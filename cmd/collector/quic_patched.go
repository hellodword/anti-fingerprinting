//go:build patched
// +build patched

package main

import (
	"net"
	"time"

	"github.com/quic-go/quic-go"
)

func applyQuicPatch(quicConf *quic.Config) *quic.Config {
	quicConf.FirstPacketHijacker = func(quicFirst []byte, remoteAddr net.Addr, rcvTime time.Time) {
		quicFirstClone := make([]byte, len(quicFirst))
		copy(quicFirstClone, quicFirst)
		quicFirstPacketPool.Store(remoteAddr.String(), quicFirstClone)
	}
	return quicConf
}
