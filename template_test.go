package belajar_golang_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// template dari string
func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }
	//bentuk ringkas
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "HELLO WORLD")
}
func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

// template dengan menggunakan file -> file menggunakan ekstensi gohtml
func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML template")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

// template file dengan memamnggil langsung semua file directory (PARSE GRUP)
func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./template/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML template")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

// template embed
//
//go:embed template/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "template/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML template")
}
func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
