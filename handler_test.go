package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/*
Cara menjalankan kode
-run test sesuai fungsi yang dipilih
-masuk ke localhost:8080 pada browser
*/

// Materi Handler
func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		//logic web
		fmt.Fprint(w, "Hello Word")
	}
	server := http.Server{
		Addr:    "Localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Materi ServeMux
// Mux -> membantu pembuatan Endpoint (url) lebih dari satu
func TestMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Saya Bima")
	})
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Image")
	})
	mux.HandleFunc("/tumb/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Tumbnails")
	})

	server := http.Server{
		Addr:    "Localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Materi Request
func TestRequst(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr:    "Localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
