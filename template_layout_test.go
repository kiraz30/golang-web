package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./template/header.gohtml",
		"./template/footer.gohtml",
		"./template/layout.gohtml",
	))

	// t.ExecuteTemplate(w, "layout.gohtml", map[string]interface{}{
	// 	"Title": "Template layout",
	// 	"Name":  "bima setya",
	// })

	//menggunakan define ->menggilangkan nama type file
	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Template layout define",
		"Name":  "bima setya",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
