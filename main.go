package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
)

var Version = "development"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	info, header, form, body, err := dump(r)
	log.Println(info)
	log.Println(header)
	log.Println(body)
	if err != nil {
		log.Println("ERR:", err)
	}
	if len(form) > 0 {
		log.Println("FORM:", form)
	}
}

func dump(r *http.Request) (string, http.Header, url.Values, string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", nil, nil, "", fmt.Errorf("SplitHostPort err: %w", err)
	}
	info := fmt.Sprintf("%s %s %s %s%s", ip, r.Proto, r.Method, r.Host, r.URL.RequestURI())
	bodyBytes, err := io.ReadAll(r.Body)
	return info, r.Header, r.Form, string(bodyBytes), err
}
