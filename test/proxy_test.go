package test

import (
	"fmt"
	"go-gateway/proxy"
	"net/http"
	"net/url"
	"testing"
)

func TestForwardProxy(t *testing.T) {
	fmt.Println("Serve on :8080")
	http.Handle("/", &proxy.ForwardProxy{})
	http.ListenAndServe("0.0.0.0:10000", nil)
}

func TestMultipleHostsReverseProxy(t *testing.T) {
	proxy := proxy.NewMultipleHostsReverseProxy([]*url.URL{
		{
			Scheme: "http",
			Host:   "localhost:9091",
		},
		{
			Scheme: "http",
			Host:   "localhost:9092",
		},
	})
	http.ListenAndServe(":9090", proxy)
}
