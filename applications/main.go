package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"Body": "Hello World!"})
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	http.Handle("/upload", http.StripPrefix("/upload", http.HandlerFunc(UploadFile)))

	log.Fatal(http.ListenAndServe(":9090", nil))
}

func UploadFile(writer http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

}
