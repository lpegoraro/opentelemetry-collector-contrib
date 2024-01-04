// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package httpmetricsreceiver // import "github.com/lpegoraro/opentelemetry-collector-contrib/receiver/httpmetricsreceiver"

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/lpegoraro/opentelemetry-collector-contrib/receiver/httpmetricsreceiver/internal/metadata"
)

var (
	errClientNotInit    = errors.New("client not initialized")
	httpResponseClasses = map[string]int{"1xx": 1, "2xx": 2, "3xx": 3, "4xx": 4, "5xx": 5}
)

type httpmetricScraper struct {
	clients  []*http.Client
	cfg      *Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder
}

// start starts the scraper by creating a new HTTP Client on the scraper
func (h *httpmetricScraper) start(_ context.Context, host component.Host) (err error) {
	for targetIndex, target := range h.cfg.Targets {
		client, clientErr := target.ToClient(host, h.settings)
		if clientErr != nil {
			err = multierr.Append(err, clientErr)
		}
		if h.cfg.Targets[targetIndex].FollowRedirects {
			client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}
		}
		h.clients = append(h.clients, client)
	}
	return
}

// scrape connects to the endpoint and produces metrics based on the response
func (h *httpmetricScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	if h.clients == nil || len(h.clients) == 0 {
		return pmetric.NewMetrics(), errClientNotInit
	}

	var wg sync.WaitGroup
	wg.Add(len(h.clients))
	var mux sync.Mutex

	for idx, client := range h.clients {
		go func(targetClient *http.Client, targetIndex int) {
			defer wg.Done()

			now := pcommon.NewTimestampFromTime(time.Now())
			req, err := http.NewRequestWithContext(ctx, h.cfg.Targets[targetIndex].Method, h.cfg.Targets[targetIndex].Endpoint, http.NoBody)
			if err != nil {
				h.settings.Logger.Error("failed to create request", zap.String("target endpoint", h.cfg.Targets[targetIndex].Endpoint), zap.Error(err))
				return
			}

			start := time.Now()
			resp, err2 := targetClient.Do(req)
			mux.Lock()
			h.mb.RecordHttpmetricDurationDataPoint(now, time.Since(start).Milliseconds(), h.cfg.Targets[targetIndex].Endpoint)

			statusCode := 0
			if err2 != nil {
				h.mb.RecordHttpmetricErrorDataPoint(now, int64(1), h.cfg.Targets[targetIndex].Endpoint, err2.Error())
			} else {
				statusCode = resp.StatusCode
			}

			if resp != nil {
				if resp.TLS != nil {
					if resp.TLS.HandshakeComplete {
						h.mb.RecordHttpmetricTLSDataPoint(now, int64(1), h.cfg.Targets[targetIndex].Endpoint)
					} else {
						h.mb.RecordHttpmetricTLSDataPoint(now, int64(0), h.cfg.Targets[targetIndex].Endpoint)
					}
				} else {
					h.mb.RecordHttpmetricTLSDataPoint(now, int64(0), h.cfg.Targets[targetIndex].Endpoint)
				}
			}

			for class, intVal := range httpResponseClasses {
				if statusCode/100 == intVal {
					h.mb.RecordHttpmetricStatusDataPoint(now, int64(1), h.cfg.Targets[targetIndex].Endpoint, int64(statusCode), req.Method, class)
				} else {
					h.mb.RecordHttpmetricStatusDataPoint(now, int64(0), h.cfg.Targets[targetIndex].Endpoint, int64(statusCode), req.Method, class)
				}
			}

			if err2 == nil && len(h.cfg.Targets[targetIndex].ContainsText) > 0 {
				if body, err := io.ReadAll(resp.Body); err != nil {
					h.mb.RecordHttpmetricErrorDataPoint(now, int64(1), h.cfg.Targets[targetIndex].Endpoint, err.Error())
				} else {
					bodyAsStr := string(body)
					for _, text := range h.cfg.Targets[targetIndex].ContainsText {
						count := strings.Count(bodyAsStr, text)
						h.mb.RecordHttpmetricContentCountDataPoint(now, int64(count), h.cfg.Targets[targetIndex].Endpoint, []any{text})
					}
				}
			}
			if err2 == nil && len(h.cfg.Targets[targetIndex].Tags) > 0 {
				h.mb.NewResourceBuilder().SetTags(h.cfg.Targets[targetIndex].Tags)
			}
			mux.Unlock()
		}(client, idx)
	}

	wg.Wait()

	return h.mb.Emit(), nil
}

func newScraper(conf *Config, settings receiver.CreateSettings) *httpmetricScraper {
	return &httpmetricScraper{
		cfg:      conf,
		settings: settings.TelemetrySettings,
		mb:       metadata.NewMetricsBuilder(conf.MetricsBuilderConfig, settings),
	}
}
