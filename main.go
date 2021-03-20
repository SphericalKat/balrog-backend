package main

import (
	"bytes"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxy.OnResponse().DoFunc(func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		header := r.Header.Get("Content-Type")
		if r != nil && strings.Contains(header, "text/html") {
			var bodyBytes []byte

			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
			}

			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			log.Println(string(bodyBytes))
		}

		return r
	})

	log.Fatal(http.ListenAndServe(":6969", proxy))
}
