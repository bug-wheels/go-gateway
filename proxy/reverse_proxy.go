package proxy

import (
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewMultipleHostsReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}

type LoadBalanceRoute interface {
	ObtainInstance(path string) url.URL
}

type DiscoveryLoadBalanceRoute struct {
}

func NewLoadBalanceReverseProxy(lb LoadBalanceRoute) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := lb.ObtainInstance(req.URL.Path)
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}
	return &httputil.ReverseProxy{Director: director}
}
