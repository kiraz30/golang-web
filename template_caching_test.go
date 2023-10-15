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

// template embed
//
//go:embed template/*.gohtml
var templates embed.FS

// pengcompilan sekali membuat pemanggilan oleh handler/func akan lebih cepat (global variable) dan tidak perlu parsing berulang di tiap handler
var myTemplates = template.Must(template.ParseFS(templates, "template/*.gohtml"))

// handler pemanggil
func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
