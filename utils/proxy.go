package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func generateProxy(proxyurl string) *handler {
	url, err := url.Parse(proxyurl)
	if err != nil {
		panic(err)
	}

	director := func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
	}

	reverseProxy := &httputil.ReverseProxy{Director: director}
	return &handler{proxy: reverseProxy}
}

func setupProxy() {
	// setup the proxies
	http.Handle("/app/", generateProxy("http://localhost:3000"))
	http.Handle("/api/", generateProxy("http://localhost:3001"))
	fmt.Println(GREEN + "reverse proxy is running!" + RESET)
	http.ListenAndServe("localhost:3002", nil)
}

type handler struct {
	proxy *httputil.ReverseProxy
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.proxy.ServeHTTP(w, r)
}
