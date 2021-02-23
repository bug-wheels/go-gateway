package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	})
	http.ListenAndServe(":9090", proxy)
}

func TestGin(t *testing.T)  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9091") // listen and serve on 0.0.0.0:9091
}
