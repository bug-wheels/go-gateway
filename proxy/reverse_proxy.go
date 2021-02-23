package proxy

import (
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
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

type Route struct {
	Path string
	ServiceName string
}

type DiscoveryLoadBalanceRoute struct {

	DiscoveryClient DiscoveryClient

	Routes []Route

}

func (d DiscoveryLoadBalanceRoute) ObtainInstance(path string) *url.URL {
	for _, route := range d.Routes {
		if strings.Index(path, route.Path) == 0 {
			instances, _ := d.DiscoveryClient.GetInstances(route.ServiceName)
			instance := instances[rand.Int()%len(instances)]
			scheme := "http"
			return &url.URL{
				Scheme: scheme,
				Host: instance.GetHost(),
			}
		}
	}
	return nil
}

func NewLoadBalanceReverseProxy(lb LoadBalanceRoute) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := lb.ObtainInstance(req.URL.Path)
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}
	return &httputil.ReverseProxy{Director: director}
}
