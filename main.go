package main

import (
	"fmt"
	"go-gateway/proxy"
	"net/http"
	"net/url"
)

func main() {
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
	err := http.ListenAndServe(":19090", proxy)
	fmt.Println(err)
}
