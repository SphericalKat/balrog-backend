package main

import (
	"github.com/elazarl/goproxy"
	"github.com/gandalf/api"
	"log"
	"net/http"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnResponse().DoFunc(api.HandleResponse) // add listener on proxy response

	log.Fatal(http.ListenAndServe(":6969", proxy)) // serve proxy
}
