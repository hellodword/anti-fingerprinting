package common

import (
	"log"
)

func (src CollectedInfo) Equals(dst CollectedInfo) bool {
	if src.UserAgent != dst.UserAgent {
		log.Println("UserAgent")
		return false
	}

	if src.Proto != dst.Proto {
		log.Println("Proto")
		return false
	}

	// don't compare illegal info
	{
		if src.FingerProxy == nil || dst.FingerProxy == nil {
			log.Println("nil FingerProxy")
			return false
		}
		if src.Clienthellod == nil || dst.Clienthellod == nil {
			log.Println("nil Clienthellod")
			return false
		}
	}

	if src.TLS == nil && dst.TLS != nil {
		log.Println("dst TLS")
		return false
	}
	if src.TLS != nil && dst.TLS == nil {
		log.Println("src TLS")
		return false
	}

	if src.TLS != nil && dst.TLS != nil {
		if src.TLS.Version != dst.TLS.Version {
			log.Println("TLS.Version")
			return false
		}
		if src.TLS.HandshakeComplete != dst.TLS.HandshakeComplete {
			log.Println("TLS.HandshakeComplete")
			return false
		}
		if src.TLS.DidResume != dst.TLS.DidResume {
			log.Println("TLS.DidResume")
			return false
		}
		if src.TLS.CipherSuite != dst.TLS.CipherSuite {
			log.Println("TLS.CipherSuite")
			return false
		}
		if src.TLS.NegotiatedProtocol != dst.TLS.NegotiatedProtocol {
			log.Println("TLS.NegotiatedProtocol")
			return false
		}
		if src.TLS.ServerName != dst.TLS.ServerName {
			log.Println("TLS.ServerName")
			return false
		}
	}

	// chromium contains shuffle, do not compare ja3
	if false {
		if src.FingerProxy.Detail.JA3.HandshakeType != dst.FingerProxy.Detail.JA3.HandshakeType {
			log.Println("FingerProxy.Detail.JA3.HandshakeType")
			return false
		}
		if src.FingerProxy.Detail.JA3.HandshakeVersion != dst.FingerProxy.Detail.JA3.HandshakeVersion {
			log.Println("FingerProxy.Detail.JA3.HandshakeVersion")
			return false
		}
		if src.FingerProxy.Detail.JA3.CipherSuiteLen != dst.FingerProxy.Detail.JA3.CipherSuiteLen {
			log.Println("FingerProxy.Detail.JA3.CipherSuiteLen")
			return false
		}
		if src.FingerProxy.Detail.JA3.ExtensionLen != dst.FingerProxy.Detail.JA3.ExtensionLen {
			log.Println("FingerProxy.Detail.JA3.ExtensionLen")
			return false
		}
		if src.FingerProxy.Detail.JA3.SNI != dst.FingerProxy.Detail.JA3.SNI {
			log.Println("FingerProxy.Detail.JA3.SNI")
			return false
		}

		if len(src.FingerProxy.Detail.JA3.CipherSuites) != len(dst.FingerProxy.Detail.JA3.CipherSuites) {
			log.Println("FingerProxy.Detail.JA3.CipherSuites")
			return false
		}
		for i := range src.FingerProxy.Detail.JA3.CipherSuites {
			if src.FingerProxy.Detail.JA3.CipherSuites[i] != dst.FingerProxy.Detail.JA3.CipherSuites[i] {
				log.Println("FingerProxy.Detail.JA3.CipherSuites")
				return false
			}
		}

		if len(src.FingerProxy.Detail.JA3.SupportedGroups) != len(dst.FingerProxy.Detail.JA3.SupportedGroups) {
			log.Println("FingerProxy.Detail.JA3.SupportedGroups")
			return false
		}
		for i := range src.FingerProxy.Detail.JA3.SupportedGroups {
			if src.FingerProxy.Detail.JA3.SupportedGroups[i] != dst.FingerProxy.Detail.JA3.SupportedGroups[i] {
				log.Println("FingerProxy.Detail.JA3.SupportedGroups")
				return false
			}
		}

		if len(src.FingerProxy.Detail.JA3.SupportedPoints) != len(dst.FingerProxy.Detail.JA3.SupportedPoints) {
			log.Println("FingerProxy.Detail.JA3.SupportedPoints")
			return false
		}
		for i := range src.FingerProxy.Detail.JA3.SupportedPoints {
			if src.FingerProxy.Detail.JA3.SupportedPoints[i] != dst.FingerProxy.Detail.JA3.SupportedPoints[i] {
				log.Println("FingerProxy.Detail.JA3.SupportedPoints")
				return false
			}
		}

		if len(src.FingerProxy.Detail.JA3.AllExtensions) != len(dst.FingerProxy.Detail.JA3.AllExtensions) {
			log.Println("FingerProxy.Detail.JA3.AllExtensions")
			return false
		}
		for i := range src.FingerProxy.Detail.JA3.AllExtensions {
			if src.FingerProxy.Detail.JA3.AllExtensions[i] != dst.FingerProxy.Detail.JA3.AllExtensions[i] {
				log.Println("FingerProxy.Detail.JA3.AllExtensions")
				return false
			}
		}
	}

	{
		if src.FingerProxy.Detail.JA4.Protocol != dst.FingerProxy.Detail.JA4.Protocol {
			log.Println("FingerProxy.Detail.JA4.Protocol")
			return false
		}
		if src.FingerProxy.Detail.JA4.TLSVersion != dst.FingerProxy.Detail.JA4.TLSVersion {
			log.Println("FingerProxy.Detail.JA4.TLSVersion")
			return false
		}
		if src.FingerProxy.Detail.JA4.SNI != dst.FingerProxy.Detail.JA4.SNI {
			log.Println("FingerProxy.Detail.JA4.SNI")
			return false
		}
		if src.FingerProxy.Detail.JA4.NumberOfCipherSuites != dst.FingerProxy.Detail.JA4.NumberOfCipherSuites {
			log.Println("FingerProxy.Detail.JA4.NumberOfCipherSuites")
			return false
		}
		if src.FingerProxy.Detail.JA4.NumberOfExtensions != dst.FingerProxy.Detail.JA4.NumberOfExtensions {
			log.Println("FingerProxy.Detail.JA4.NumberOfExtensions")
			return false
		}
		if src.FingerProxy.Detail.JA4.FirstALPN != dst.FingerProxy.Detail.JA4.FirstALPN {
			log.Println("FingerProxy.Detail.JA4.FirstALPN")
			return false
		}

		if len(src.FingerProxy.Detail.JA4.CipherSuites) != len(dst.FingerProxy.Detail.JA4.CipherSuites) {
			log.Println("FingerProxy.Detail.JA4.CipherSuites")
			return false
		}
		for i := range src.FingerProxy.Detail.JA4.CipherSuites {
			if src.FingerProxy.Detail.JA4.CipherSuites[i] != dst.FingerProxy.Detail.JA4.CipherSuites[i] {
				log.Println("FingerProxy.Detail.JA4.CipherSuites")
				return false
			}
		}

		if len(src.FingerProxy.Detail.JA4.Extensions) != len(dst.FingerProxy.Detail.JA4.Extensions) {
			log.Println("FingerProxy.Detail.JA4.Extensions")
			return false
		}
		for i := range src.FingerProxy.Detail.JA4.Extensions {
			if src.FingerProxy.Detail.JA4.Extensions[i] != dst.FingerProxy.Detail.JA4.Extensions[i] {
				log.Println("FingerProxy.Detail.JA4.Extensions")
				return false
			}
		}

		if len(src.FingerProxy.Detail.JA4.SignatureAlgorithms) != len(dst.FingerProxy.Detail.JA4.SignatureAlgorithms) {
			log.Println("FingerProxy.Detail.JA4.SignatureAlgorithms")
			return false
		}
		for i := range src.FingerProxy.Detail.JA4.SignatureAlgorithms {
			if src.FingerProxy.Detail.JA4.SignatureAlgorithms[i] != dst.FingerProxy.Detail.JA4.SignatureAlgorithms[i] {
				log.Println("FingerProxy.Detail.JA4.SignatureAlgorithms")
				return false
			}
		}
	}

	{
		if uint32(src.FingerProxy.Detail.HTTP2.WindowUpdateIncrement/100) != uint32(dst.FingerProxy.Detail.HTTP2.WindowUpdateIncrement/100) {
			log.Println("FingerProxy.Detail.HTTP2.WindowUpdateIncrement")
			return false
		}

		if len(src.FingerProxy.Detail.HTTP2.Settings) != len(dst.FingerProxy.Detail.HTTP2.Settings) {
			log.Println("FingerProxy.Detail.HTTP2.Settings")
			return false
		}
		for i := range src.FingerProxy.Detail.HTTP2.Settings {
			if src.FingerProxy.Detail.HTTP2.Settings[i] != dst.FingerProxy.Detail.HTTP2.Settings[i] {
				log.Println("FingerProxy.Detail.HTTP2.Settings")
				return false
			}
		}

		if len(src.FingerProxy.Detail.HTTP2.Priorities) != len(dst.FingerProxy.Detail.HTTP2.Priorities) {
			log.Println("FingerProxy.Detail.HTTP2.Priorities")
			return false
		}
		for i := range src.FingerProxy.Detail.HTTP2.Priorities {
			if src.FingerProxy.Detail.HTTP2.Priorities[i] != dst.FingerProxy.Detail.HTTP2.Priorities[i] {
				log.Println("FingerProxy.Detail.HTTP2.Priorities")
				return false
			}
		}
	}

	if !((src.Clienthellod.TLS != nil && dst.Clienthellod.TLS != nil) ||
		(src.Clienthellod.QUIC != nil && dst.Clienthellod.QUIC != nil)) {
		log.Println("Clienthellod mismatch")
		return false
	}

	if src.Clienthellod.TLS != nil && dst.Clienthellod.TLS != nil {
		if src.Clienthellod.TLS.TLSRecordVersion != dst.Clienthellod.TLS.TLSRecordVersion {
			log.Println("Clienthellod.TLS.TLSRecordVersion")
			return false
		}

		if src.Clienthellod.TLS.TLSHandshakeVersion != dst.Clienthellod.TLS.TLSHandshakeVersion {
			log.Println("Clienthellod.TLS.TLSHandshakeVersion")
			return false
		}

		if src.Clienthellod.TLS.ServerName != dst.Clienthellod.TLS.ServerName {
			log.Println("Clienthellod.TLS.ServerName")
			return false
		}

		if src.Clienthellod.TLS.UserAgent != dst.Clienthellod.TLS.UserAgent {
			log.Println("Clienthellod.TLS.UserAgent")
			return false
		}

		if src.Clienthellod.TLS.FingerprintNID(true) != dst.Clienthellod.TLS.FingerprintNID(true) {
			log.Println("Clienthellod.TLS.FingerprintNID(true)")
			return false
		}

		if src.Clienthellod.TLS.FingerprintID(true) != dst.Clienthellod.TLS.FingerprintID(true) {
			log.Println("Clienthellod.TLS.FingerprintID(true)")
			return false
		}
	}

	// if src.Clienthellod.QUIC != nil && dst.Clienthellod.QUIC != nil {	}

	return true
}
