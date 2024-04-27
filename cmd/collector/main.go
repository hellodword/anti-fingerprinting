package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/dreadl0ck/tlsx"
	"github.com/gaukas/clienthellod"
	"github.com/google/uuid"
	"github.com/hellodword/anti-fingerprinting/internal/common"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/wi1dcard/fingerproxy/pkg/fingerprint"
	"github.com/wi1dcard/fingerproxy/pkg/http2"
	"github.com/wi1dcard/fingerproxy/pkg/ja3"
	"github.com/wi1dcard/fingerproxy/pkg/ja4"
	"github.com/wi1dcard/fingerproxy/pkg/metadata"
	"github.com/wi1dcard/fingerproxy/pkg/proxyserver"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Record struct {
	gorm.Model

	UUID      string `gorm:"type:varchar;UNIQUE" json:"uuid"`
	UserAgent string `gorm:"type:varchar" json:"user_agent"`
	Proto     string `gorm:"type:varchar" json:"proto"`
	Path      string `gorm:"type:varchar" json:"path"`
	Query     string `gorm:"type:varchar" json:"query"`
	Data      string `gorm:"type:varchar" json:"data"`
}

var quicFirstPacketPool sync.Map

func main() {
	flagListenAddr := flag.String(
		"addr",
		"0.0.0.0:8443",
		"Listening address",
	)
	flagCert := flag.String(
		"cert",
		"certs/tls.crt",
		"TLS certificate file",
	)
	flagKey := flag.String(
		"key",
		"certs/tls.key",
		"TLS certificate key file",
	)

	flagDB := flag.String(
		"db",
		"db/collector.db",
		"SQLite3 db file",
	)

	flagVerboseLogs := flag.Bool("verbose", true, "Enable verbose logs")
	flag.Parse()

	// load TLS certs
	tlsConf := &tls.Config{
		NextProtos: []string{"h2", "http/1.1", "hq-interop"},
	}
	if tlsCert, err := tls.LoadX509KeyPair(*flagCert, *flagKey); err != nil {
		log.Fatal(err)
	} else {
		tlsConf.Certificates = []tls.Certificate{tlsCert}
	}

	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s?_mutex=full&cache=shared&mode=rwc&_journal_mode=WAL", *flagDB)),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	rawDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	rawDB.SetMaxOpenConns(1)

	db.AutoMigrate(&Record{})

	// enable verbose logs
	fingerprint.VerboseLogs = *flagVerboseLogs
	http2.VerboseLogs = *flagVerboseLogs

	// shutdown on interrupt signal (ctrl + c)
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(cors.AllowAll().Handler)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/all", echoServer(db))
		r.Get("/id/{id}", getByID(db))
	})

	// create proxyserver
	server := proxyserver.NewServer(ctx, r, tlsConf)
	server.VerboseLogs = *flagVerboseLogs

	quicTlsConf := http3.ConfigureTLSConfig(tlsConf)
	quicConf := applyQuicPatch(&quic.Config{
		Allow0RTT:       true,
		EnableDatagrams: false,
	})
	quicServer := &http3.Server{
		Addr:       *flagListenAddr,
		Handler:    r,
		TLSConfig:  quicTlsConf,
		QuicConfig: quicConf,
	}

	// listen and serve
	var errC = make(chan error, 2)

	go func() {
		log.Printf("server https listening on %s", *flagListenAddr)
		errC <- server.ListenAndServe(*flagListenAddr)
	}()

	go func() {
		log.Printf("server http3 listening on %s", *flagListenAddr)
		errC <- quicServer.ListenAndServe()
	}()

	log.Fatal(<-errC)
}

func echoServer(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if len(req.URL.String()) > 1<<12 {
			http.Error(w, "long url", http.StatusBadRequest)
			return
		}

		_uuid := req.URL.Query().Get("id")
		if _uuid == "" {
			_uuid = uuid.New().String()
		}

		if len(_uuid) != 36 {
			http.Error(w, "uuid", http.StatusBadRequest)
			return
		}

		var res = common.CollectedInfo{
			Date:      time.Now(),
			URL:       req.URL.String(),
			UserAgent: req.UserAgent(),
			Headers:   req.Header,
			Proto:     req.Proto,
			TLS:       req.TLS,
			ID:        _uuid,
		}

		data, ok := metadata.FromContext(req.Context())
		if ok {

			if len(data.ClientHelloRecord) > 1<<16 {
				http.Error(w, "long client hello", http.StatusInternalServerError)
				return
			}

			_ja3 := &tlsx.ClientHelloBasic{}
			err := _ja3.Unmarshal(data.ClientHelloRecord)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_ja4 := &ja4.JA4Fingerprint{}
			err = _ja4.UnmarshalBytes(data.ClientHelloRecord, 't')
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_http2, err := fingerprint.HTTP2Fingerprint(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			ch, err := clienthellod.ReadClientHello(bytes.NewReader(data.ClientHelloRecord))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = ch.ParseClientHello()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res.FingerProxy = &common.CollectedInfoFingerProxy{
				ClientHello: data.ClientHelloRecord,
				JA3:         ja3.DigestHex(_ja3),
				JA4:         _ja4.String(),
				HTTP2:       _http2,
			}

			res.FingerProxy.Detail.JA3 = _ja3
			res.FingerProxy.Detail.JA4 = _ja4
			res.FingerProxy.Detail.HTTP2 = data.HTTP2Frames
			res.FingerProxy.Detail.MetaData = data

			res.Clienthellod = &common.CollectedInfoClienthellod{
				TLS: ch,
				Raw: ch.Raw(),
			}

		} else {
			defer quicFirstPacketPool.Delete(req.RemoteAddr)

			var quicFirst []byte
			if v, ok := quicFirstPacketPool.Load(req.RemoteAddr); ok {
				quicFirst = v.([]byte)
			}

			if len(quicFirst) > 1<<16 {
				http.Error(w, "long quic", http.StatusInternalServerError)
				return
			}

			cip, err := clienthellod.ParseQUICCIP(quicFirst)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Clienthellod = &common.CollectedInfoClienthellod{
				QUIC: cip,
				Raw:  quicFirst,
			}
		}

		var buf bytes.Buffer

		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(&buf)
		// if req.Header.Get("beautify") != "" {
		// 	encoder.SetIndent("", "  ")
		// }
		if err := encoder.Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(buf.Bytes())

		var record = Record{
			UserAgent: req.UserAgent(),
			Proto:     req.Proto,
			Data:      buf.String(),
			Path:      req.URL.Path,
			Query:     req.URL.RawQuery,
			UUID:      _uuid,
		}
		go func() {
			err := db.Model(&Record{}).Save(&record).Error
			if err != nil {
				log.Println("db error", err)
			}
		}()
	}

}

func getByID(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		_uuid := chi.URLParam(req, "id")

		if len(_uuid) != 36 {
			http.Error(w, "uuid", http.StatusBadRequest)
			return
		}

		var record Record

		err := db.Model(&Record{}).Where("uuid = ?", _uuid).First(&record).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(record.Data))
	}
}
