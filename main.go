package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gomicro/blockit/dbblocker"
	"github.com/gomicro/ledger"
	"github.com/gomicro/steward"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

const (
	appName = "superman"
)

var (
	appVersion string

	log    *ledger.Ledger
	config configuration

	db *sql.DB
)

type configuration struct {
	Host    string `default:"0.0.0.0"`
	Port    string `default:"4567"`
	SSLCert string
	SSLKey  string

	LogLevel string `default:"debug"`

	DBHost     string `default:"database"`
	DBPort     string `default:"5432"`
	DBName     string `default:"superman"`
	DBUser     string `default:"superman"`
	DBPassword string `default:"superman"`
	DBSSL      string `default:"disable"`
}

type statusResponse struct {
	Application string    `json:"app"`
	Version     string    `json:"version"`
	SSLStatus   sslStatus `json:"ssl"`
}

type sslStatus struct {
	ServingSSL bool `json:"serving_ssl"`
	DBSSL      bool `json:"db_ssl"`
}

func main() {
	configure()

	err := startService()
	if err != nil {
		log.Errorf("Something went wrong: %v", err.Error())
		return
	}

	log.Info("Server stopping")
}

func configure() {
	err := envconfig.Process(appName, &config)
	if err != nil {
		fmt.Printf("Fatal: Unable to configure service: %v\n", err.Error())
		os.Exit(1)
	}

	log = ledger.New(os.Stdout, ledger.ParseLevel(config.LogLevel))
	log.Debug("Logger configured")

	k, err := base64.StdEncoding.DecodeString(config.SSLKey)
	if err != nil {
		log.Warnf("Failed to decode base64 data from ssl key: %v", err.Error())
	} else {
		log.Debug("Decoding base64 encoded ssl key")
		config.SSLKey = string(k)
	}

	c, err := base64.StdEncoding.DecodeString(config.SSLCert)
	if err != nil {
		log.Warnf("Failed to decode base64 data from ssl cert: %v", err.Error())
	} else {
		log.Debug("Decoding base64 encoded ssl cert")
		config.SSLCert = string(c)
	}

	steward.SetStatusEndpoint("/v1/status")
	steward.SetStatusResponse(&statusResponse{
		Application: appName,
		Version:     appVersion,
		SSLStatus: sslStatus{
			ServingSSL: (config.SSLKey != "" && config.SSLCert != ""),
			DBSSL:      (config.DBSSL != "disable"),
		},
	})
	log.Debug("Status endpoint configured")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSL,
	)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Errorf("Cannot open sql connection: %v", err.Error())
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	dbb := dbblocker.New(db)
	log.Debug("Database connection configured")

	log.Debug("Waiting for dependencies to stablize")
	<-dbb.Blockit()
	log.Debug("Dependencies have stablized")

	log.Info("Configuration complete")
}

func startService() error {
	log.Infof("Listening on %v:%v", config.Host, config.Port)

	http.Handle("/", registerEndpoints())

	if config.SSLKey != "" && config.SSLCert != "" {
		cert, err := tls.X509KeyPair([]byte(config.SSLCert), []byte(config.SSLKey))
		if err != nil {
			log.Fatalf("failed to create ssl cert/key pair: %v", err.Error())
			os.Exit(1)
		}

		cfg := &tls.Config{
			MinVersion:               tls.VersionTLS12,
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
			Certificates: []tls.Certificate{cert},
		}

		srv := &http.Server{
			Addr:      net.JoinHostPort(config.Host, config.Port),
			TLSConfig: cfg,
		}

		log.Info("Serving with SSL")
		return srv.ListenAndServeTLS("", "")
	}

	log.Info("Serving without SSL")
	return http.ListenAndServe(net.JoinHostPort(config.Host, config.Port), nil)
}

func registerEndpoints() http.Handler {
	r := mux.NewRouter()

	r.Use(log.EndpointInfo)

	r.HandleFunc("/v1/event", handleSomeEndpoint)

	return r
}
