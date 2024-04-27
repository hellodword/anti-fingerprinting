module github.com/hellodword/anti-fingerprinting

go 1.22

require (
	github.com/dop251/goja v0.0.0-20240220182346-e401ed450204
	github.com/dreadl0ck/tlsx v1.0.1-google-gopacket
	github.com/gaukas/clienthellod v0.4.2
	github.com/go-chi/chi/v5 v5.0.12
	github.com/go-chi/cors v1.2.1
	github.com/google/uuid v1.3.1
	github.com/quic-go/quic-go v0.42.0
	github.com/refraction-networking/utls v1.6.4
	github.com/wi1dcard/fingerproxy v0.5.0
	golang.org/x/net v0.24.0
	gorm.io/driver/sqlite v1.5.5
	gorm.io/gorm v1.25.9
)

require (
	github.com/andybalholm/brotli v1.0.6 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/dlclark/regexp2 v1.7.0 // indirect
	github.com/gaukas/godicttls v0.0.4 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/pprof v0.0.0-20231212022811-ec68065c825e // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/onsi/ginkgo/v2 v2.13.2 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	go.uber.org/mock v0.4.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.16.1 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace golang.org/x/net v0.24.0 => github.com/hellodword/http2-custom-fingerprint v0.24.1-0.20240421163422-e10ffbc4cee8
