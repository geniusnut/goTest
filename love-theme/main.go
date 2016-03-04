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
		rest.Get("/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"msg" : "Hello, world"})
		}),
		rest.Get("/user/:id", func(w rest.ResponseWriter, req *rest.Request) {
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
		rest.Get("/login", func(w rest.ResponseWriter, req *rest.Request) {
			name := req.URL.Query().Get("name")
			token := req.URL.Query().Get("token")
			c.Do("MULTI")
			c.Do("SADD", "users", token)
			c.Do("HMSET", "user:" + token, "name", name)
			c.Do("EXEC")
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	http.Handle("/upload", http.StripPrefix("/upload", http.HandlerFunc(UploadPicture)))
	
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func UploadPicture(w http.ResponseWriter, r *http.Request) {
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
