package proxy

import (
	"net/http"
	"net/url"
)

type OriginReverseProxy struct {
	servers []*url.URL
}

func NewOriginReverseProxy(targets []*url.URL) *OriginReverseProxy {
	return &OriginReverseProxy{
		servers: targets,
	}
}

func (proxy *OriginReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// 1. 复制了原来的请求对象
	r2 := new(http.Request)
	*r2 = *req
}
