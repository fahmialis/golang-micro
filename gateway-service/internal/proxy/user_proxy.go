package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func User() http.Handler {
	target := os.Getenv("USER_SERVICE_URL")
	if target == "" {
		target = "http://localhost:4001"
	}

	url, err := url.Parse(target)
	if err != nil {
		log.Fatalf("invalid user service url: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	return http.StripPrefix("/users", proxy)
}
