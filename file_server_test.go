package belajar_golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// Fileserever untuk memanggil data/ file
func TestFileServer(t *testing.T) {
	directory := http.Dir("./resource")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	//stripPrefix untuk menghapus perfix di url
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//===============Golang_Embed=================

//go:embed resource
var resource embed.FS

func TestFileServerGolangEmbed(t *testing.T) {

	directory, _ := fs.Sub(resource, "resource")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	//stripPrefix untuk menghapus perfix di url
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
