package belajar_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		//http://localhost:8080/?name=bima
		http.ServeFile(w, r, "./resource/ok.html")
	} else {
		http.ServeFile(w, r, "./resource/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resource/ok.html
var resourceOk string

//go:embed resource/notfound.html
var resourceNotfound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		//http://localhost:8080/?name=bima
		fmt.Fprint(w, resourceOk)
	} else {
		fmt.Fprint(w, resourceNotfound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
