package api

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func HandleResponse(r *http.Response, _ *goproxy.ProxyCtx) *http.Response {
	// get content type
	header := r.Header.Get("Content-Type")
	if r != nil && strings.Contains(header, "text/html") { // check if response is html
		var bodyBytes []byte

		bodyBytes, err := ioutil.ReadAll(r.Body) // read all bytes from body
		if err != nil {
			log.Println(err)
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // recreate reader from saved byte slice

		doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(bodyBytes)) // create a goquery doc from response body
		if err != nil {
			log.Println(err)
			return r
		}

		elemList := make([]string, 0)
		doc.Find("h1, h2, h3, h4, h5, h6, a, p, div, article, tr, li"). // find tags which typically contain text
			Each(func(i int, selection *goquery.Selection) {
				elemList = append(elemList, strings.TrimSpace(selection.Text())) // add cleaned text to slice
			})

		log.Println(elemList)
	}

	return r
}
