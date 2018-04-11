package main

import (
	"context"
	"crypto/tls"
	"expvar"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	envstruct "code.cloudfoundry.org/go-envstruct"
	logcache "code.cloudfoundry.org/go-log-cache"
	"code.cloudfoundry.org/loggregator-tools/log-cache-forwarders/cmd/syslog/internal/egress"
	"code.cloudfoundry.org/loggregator-tools/log-cache-forwarders/pkg/expvarfilter"
	"code.cloudfoundry.org/loggregator-tools/log-cache-forwarders/pkg/groupmanager"
	"code.cloudfoundry.org/loggregator-tools/log-cache-forwarders/pkg/metrics"
	"code.cloudfoundry.org/loggregator-tools/log-cache-forwarders/pkg/sourceidprovider"
)

const metricsNamespace = "SyslogForwarder"

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := LoadConfig()
	envstruct.WriteReport(&cfg)

	m := metrics.New(expvar.NewMap(metricsNamespace))

	mh := expvarfilter.NewHandler(expvar.Handler(), []string{metricsNamespace})
	go func() {
		// health endpoints (expvar)
		log.Printf("Health: %s", http.ListenAndServe(":"+os.Getenv("PORT"), mh))
	}()

	client := logcache.NewClient(
		cfg.LogCacheHTTPAddr,
		logcache.WithHTTPClient(newOauth2HTTPClient(cfg)),
	)

	groupClient := logcache.NewShardGroupReaderClient(
		cfg.LogCacheHTTPAddr,
		logcache.WithHTTPClient(newOauth2HTTPClient(cfg)),
	)

	provider := sourceidprovider.NewRegex(
		false,
		cfg.SourceID,
		client,
	)

	groupmanager.Start(
		cfg.GroupName,
		time.Tick(30*time.Second),
		provider,
		groupClient,
		groupmanager.WithMetrics(m),
	)

	netConf := egress.NetworkConfig{
		Keepalive:      cfg.KeepAlive,
		DialTimeout:    cfg.DialTimeout,
		WriteTimeout:   cfg.IOTimeout,
		SkipCertVerify: cfg.SkipCertVerify,
	}
	writer := egress.NewWriter(cfg.SourceID, cfg.SourceHostname, cfg.SyslogURL, netConf)

	reader := groupClient.BuildReader(rand.Uint64())
	logcache.Walk(
		context.Background(),
		cfg.GroupName,
		egress.NewVisitor(writer, m),
		reader,
		logcache.WithWalkStartTime(time.Now()),
		logcache.WithWalkBackoff(logcache.NewAlwaysRetryBackoff(250*time.Millisecond)),
		logcache.WithWalkLogger(log.New(os.Stderr, "", log.LstdFlags)),
	)
}

func newOauth2HTTPClient(cfg Config) *logcache.Oauth2HTTPClient {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: cfg.SkipCertVerify,
			},
		},
		Timeout: 5 * time.Second,
	}

	return logcache.NewOauth2HTTPClient(
		cfg.UAAAddr,
		cfg.ClientID,
		cfg.ClientSecret,
		logcache.WithOauth2HTTPClient(client),
		logcache.WithUser(cfg.Username, cfg.Password),
	)
}
