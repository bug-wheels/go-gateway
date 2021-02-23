package main

import (
	"go-gateway/proxy"
	"net/http"
)

func main() {
	registry, _ := proxy.NewConsulServiceRegistry("127.0.0.1", 8500, "")
	reverseProxy := proxy.NewLoadBalanceReverseProxy(&proxy.DiscoveryLoadBalanceRoute{
		DiscoveryClient: registry,
		Routes: []proxy.Route{
			{
				Path: "abc",
				ServiceName: "abc",
			},
		},
	})
	http.ListenAndServe(":19090", reverseProxy)
}
