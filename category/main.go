package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type Wallpaper struct {
	url string
	id  int
}

type Wallpapers struct {
	sync.RWMutex
	Store map[string]*Wallpaper
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	//strCategories := []string{"recommand", "natural", "animal", "architecture", "tech", "comic", "sport", "color"}
	router, err := rest.MakeRouter(
		rest.Get("/category/all", func(w rest.ResponseWriter, req *rest.Request) {
			buf, _ := ioutil.ReadFile("category.txt")
			w.Header().Set("Content-Type", "text/plain")
			w.(http.ResponseWriter).Write(buf)
		}),
		rest.Get("/category/list", func(w rest.ResponseWriter, req *rest.Request) {
			id := req.URL.Query().Get("categoryid")
			absPath, _ := filepath.Abs("page" + id + ".txt")
			buf, err := ioutil.ReadFile(absPath)
			if err != nil {
				rest.NotFound(w, req)
				return
			}
			w.Header().Set("Content-Type", "text/plain")
			w.(http.ResponseWriter).Write(buf)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func (ws *Wallpapers) GetWallpapersByCategory(w rest.ResponseWriter, req *rest.Request) {
	ws.RLock()
	wallpapers := make([]Wallpaper, len(ws.Store))
	i := 0
	for _, wallpaper := range ws.Store {
		wallpapers[i] = *wallpaper
		i++
	}
	ws.RUnlock()
	buf, _ := ioutil.ReadFile("test.txt")
	str := string(buf)
	w.WriteJson(str)
}
