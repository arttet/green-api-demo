package handler

import (
	"log"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type GreenAPIProxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
	logger *log.Logger
}

func NewGreenAPIProxy(apiBaseURL string, logger *slog.Logger) (*GreenAPIProxy, error) {
	target, err := url.Parse(apiBaseURL)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)

		req.URL.Path = strings.TrimPrefix(
			req.URL.Path,
			"/v1/api/proxy",
		)

		req.Host = target.Host
		req.Header.Set("X-Forwarded-Host", target.Host)

		logger.Info(
			"Green API request",
			slog.String("method", req.Method),
			slog.String("url", req.URL.String()),
			slog.String("upstream", target.Host),
		)
	}

	proxy.ModifyResponse = func(res *http.Response) error {
		res.Header.Del("Access-Control-Allow-Origin")
		res.Header.Del("Access-Control-Allow-Methods")
		res.Header.Set("Access-Control-Allow-Headers", "")

		logger.Info(
			"Green API response",
			slog.Int("status", res.StatusCode),
			slog.String("url", res.Request.URL.String()),
		)
		return nil
	}

	return &GreenAPIProxy{
		target: target,
		proxy:  proxy,
	}, nil
}

func (p *GreenAPIProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.proxy.ServeHTTP(w, r)
}
