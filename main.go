package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type ProxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	log.Println(r.URL)

	ph.p.ServeHTTP(w, r)
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*req).Host = (*req).URL.Host
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

type DebugTransport struct{}

func (DebugTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	b, err := httputil.DumpRequestOut(r, false)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))
	return http.DefaultTransport.RoundTrip(r)
}

func handleRequests() {
	proxy_url := os.Getenv("PROXY_URL")

	// exit if empty
	if proxy_url == "" {
		log.Fatal("PROXY_URL undefined, please read README!")
		return
	}
	remote, err := url.Parse(proxy_url)
	log.Printf("forwarding to -> %s\n", remote)
	if err != nil {
		log.Panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	if os.Getenv("VERBOSE_DEBUG") == "true" {
		log.Println("verbose output enabled.")
		proxy.Transport = DebugTransport{}
	}

	http.Handle("/", &ProxyHandler{proxy})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
