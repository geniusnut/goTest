package main

import (
	"encoding/json"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/garyburd/redigo/redis"
	"io"
	"log"
	"net/http"
	"os"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/users/:id", func(w rest.ResponseWriter, req *rest.Request) {
			id := req.PathParam("id")
			_, err := c.Do("HGET", id, "name")
			user := &User{
				Id:   "id1",
				Name: "name2",
			}
			userJson, _ := json.Marshal(user)
			// nameByte := []byte(name)
			if err != nil {
				w.WriteJson(userJson)
			} else {
				w.Header().Set("Content-type", "text/plain")
				w.(http.ResponseWriter).Write([]byte("error"))
			}
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	htt
	
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func UploadPicture(w rest.ResponseWriter, r *rest.Request) {
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
