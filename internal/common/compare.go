package common

func (src CollectedInfo) Equals(dst CollectedInfo) bool {
	if src.UserAgent != dst.UserAgent {
		return false
	}

	if src.Proto != dst.Proto {
		return false
	}

	// don't compare illegal info
	{
		if src.TLS == nil || dst.TLS == nil {
			return false
		}
		if src.FingerProxy == nil || dst.FingerProxy == nil {
			return false
		}
		if src.Clienthellod == nil || dst.Clienthellod == nil {
			return false
		}
	}

	{
		if src.TLS.Version != dst.TLS.Version {
			return false
		}
		if src.TLS.HandshakeComplete != dst.TLS.HandshakeComplete {
			return false
		}
		if src.TLS.DidResume != dst.TLS.DidResume {
			return false
		}
		if src.TLS.CipherSuite != dst.TLS.CipherSuite {
			return false
		}
		if src.TLS.NegotiatedProtocol != dst.TLS.NegotiatedProtocol {
			return false
		}
		if src.TLS.ServerName != dst.TLS.ServerName {
			return false
		}
	}

	{
		if src.FingerProxy.Detail.JA3.HandshakeType != dst.FingerProxy.Detail.JA3.HandshakeType {
			return false
		}
		if src.FingerProxy.Detail.JA3.HandshakeVersion != dst.FingerProxy.Detail.JA3.HandshakeVersion {
			return false
		}
		if src.FingerProxy.Detail.JA3.CipherSuiteLen != dst.FingerProxy.Detail.JA3.CipherSuiteLen {
			return false
		}
		if src.FingerProxy.Detail.JA3.ExtensionLen != dst.FingerProxy.Detail.JA3.ExtensionLen {
			return false
		}
		if src.FingerProxy.Detail.JA3.SNI != dst.FingerProxy.Detail.JA3.SNI {
			return false
		}

		if len(src.FingerProxy.Detail.JA3.CipherSuites) != len(dst.FingerProxy.Detail.JA3.CipherSuites) {
			return false
		}
		for i := range src.FingerProxy.Detail.JA3.CipherSuites {
			if src.FingerProxy.Detail.JA3.CipherSuites[i] != dst.FingerProxy.Detail.JA3.CipherSuites[i] {
				return false
			}
		}

		if len(src.FingerProxy.Detail.JA3.SupportedGroups) != len(dst.FingerProxy.Detail.JA3.SupportedGroups) {
			return false
		}
		for i := range src.FingerProxy.Detail.JA3.SupportedGroups {
			if src.FingerProxy.Detail.JA3.SupportedGroups[i] != dst.FingerProxy.Detail.JA3.SupportedGroups[i] {
				return false
			}
		}

		if len(src.FingerProxy.Detail.JA3.SupportedPoints) != len(dst.FingerProxy.Detail.JA3.SupportedPoints) {
			return false
		}
		for i := range src.FingerProxy.Detail.JA3.SupportedPoints {
			if src.FingerProxy.Detail.JA3.SupportedPoints[i] != dst.FingerProxy.Detail.JA3.SupportedPoints[i] {
				return false
			}
		}

		if len(src.FingerProxy.Detail.JA3.AllExtensions) != len(dst.FingerProxy.Detail.JA3.AllExtensions) {
			return false
		}
		for i := range src.FingerProxy.Detail.JA3.AllExtensions {
			if src.FingerProxy.Detail.JA3.AllExtensions[i] != dst.FingerProxy.Detail.JA3.AllExtensions[i] {
				return false
			}
		}
	}

	{
		if src.FingerProxy.Detail.JA4.Protocol != dst.FingerProxy.Detail.JA4.Protocol {
			return false
		}
		if src.FingerProxy.Detail.JA4.TLSVersion != dst.FingerProxy.Detail.JA4.TLSVersion {
			return false
		}
		if src.FingerProxy.Detail.JA4.SNI != dst.FingerProxy.Detail.JA4.SNI {
			return false
		}
		if src.FingerProxy.Detail.JA4.NumberOfCipherSuites != dst.FingerProxy.Detail.JA4.NumberOfCipherSuites {
			return false
		}
		if src.FingerProxy.Detail.JA4.NumberOfExtensions != dst.FingerProxy.Detail.JA4.NumberOfExtensions {
			return false
		}
		if src.FingerProxy.Detail.JA4.FirstALPN != dst.FingerProxy.Detail.JA4.FirstALPN {
			return false
		}

		if len(src.FingerProxy.Detail.JA4.CipherSuites) != len(dst.FingerProxy.Detail.JA4.CipherSuites) {
			return false
		}
		for i := range src.FingerProxy.Detail.JA4.CipherSuites {
			if src.FingerProxy.Detail.JA4.CipherSuites[i] != dst.FingerProxy.Detail.JA4.CipherSuites[i] {
				return false
			}
		}

		if len(src.FingerProxy.Detail.JA4.Extensions) != len(dst.FingerProxy.Detail.JA4.Extensions) {
			return false
		}
		for i := range src.FingerProxy.Detail.JA4.Extensions {
			if src.FingerProxy.Detail.JA4.Extensions[i] != dst.FingerProxy.Detail.JA4.Extensions[i] {
				return false
			}
		}

		if len(src.FingerProxy.Detail.JA4.SignatureAlgorithms) != len(dst.FingerProxy.Detail.JA4.SignatureAlgorithms) {
			return false
		}
		for i := range src.FingerProxy.Detail.JA4.SignatureAlgorithms {
			if src.FingerProxy.Detail.JA4.SignatureAlgorithms[i] != dst.FingerProxy.Detail.JA4.SignatureAlgorithms[i] {
				return false
			}
		}
	}

	{
		if src.FingerProxy.Detail.HTTP2.WindowUpdateIncrement != dst.FingerProxy.Detail.HTTP2.WindowUpdateIncrement {
			return false
		}

		if len(src.FingerProxy.Detail.HTTP2.Settings) != len(dst.FingerProxy.Detail.HTTP2.Settings) {
			return false
		}
		for i := range src.FingerProxy.Detail.HTTP2.Settings {
			if src.FingerProxy.Detail.HTTP2.Settings[i] != dst.FingerProxy.Detail.HTTP2.Settings[i] {
				return false
			}
		}

		if len(src.FingerProxy.Detail.HTTP2.Priorities) != len(dst.FingerProxy.Detail.HTTP2.Priorities) {
			return false
		}
		for i := range src.FingerProxy.Detail.HTTP2.Priorities {
			if src.FingerProxy.Detail.HTTP2.Priorities[i] != dst.FingerProxy.Detail.HTTP2.Priorities[i] {
				return false
			}
		}
	}

	return true
}
