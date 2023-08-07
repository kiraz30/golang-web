package belajar_golang_web

import (
	"net/http"
	"testing"
)

// Materi Server
func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "Localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
