package proxy

import (
	"io"
	"math/rand"
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

func cloneURLValues(v url.Values) url.Values {
	if v == nil {
		return nil
	}
	// http.Header and url.Values have the same representation, so temporarily
	// treat it like http.Header, which does have a clone:
	return url.Values(http.Header(v).Clone())
}

func cloneURL(u *url.URL) *url.URL {
	if u == nil {
		return nil
	}
	u2 := new(url.URL)
	*u2 = *u
	if u.User != nil {
		u2.User = new(url.Userinfo)
		*u2.User = *u.User
	}
	return u2
}

func clone(req *http.Request) *http.Request {
	r2 := new(http.Request)
	*r2 = *req
	r2.URL = cloneURL(req.URL)
	if req.Header != nil {
		r2.Header = req.Header.Clone()
	}
	if req.Trailer != nil {
		r2.Trailer = req.Trailer.Clone()
	}
	if s := req.TransferEncoding; s != nil {
		s2 := make([]string, len(s))
		copy(s2, s)
		r2.TransferEncoding = s2
	}
	r2.Form = cloneURLValues(req.Form)
	r2.PostForm = cloneURLValues(req.PostForm)
	return r2
}

func (proxy *OriginReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// 1. 复制了原来的请求对象
	r2 := clone(req)

	// 2. 修改请求的地址，换为对应的服务器地址
	target := proxy.servers[rand.Int()%len(proxy.servers)]
	r2.URL.Scheme = target.Scheme
	r2.URL.Host = target.Host

	// 3. 发送复制的新请求
	transport := http.DefaultTransport

	res, err := transport.RoundTrip(r2)

	// 4。处理响应
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	//
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}

	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)
	res.Body.Close()
}
