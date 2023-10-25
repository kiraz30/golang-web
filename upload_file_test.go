package belajar_golang_web

import (
	"io"
	"net/http"
	"os"
	"testing"
	// "github.com/go-playground/locales/os"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {

	err := myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {

	//get file
	// .FormFile() memiliki return multipart.File, *multipart.fileHeader, error, sehingga instansiasinya
	file, fileHeader, err := r.FormFile("file") //file->name file di form
	if err != nil {
		panic(err)
	}
	// tempat penyimpanan file |create memiliki return valuw file dan error
	fileDestination, err := os.Create("./resource/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	//memindahkan file ke locasi destination
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	//get data non file
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})

}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resource/"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
