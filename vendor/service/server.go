package service

import (
	"net/http"
	"os"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/phyber/negroni-gzip/gzip"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	router := mux.NewRouter()
	initRoutes(router, formatter)
	// router.HandleFunc("/login", loginHander)

	n :=  negroni.Classic();
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(router)

	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			// fmt.Println(root)
		}
	}
	mx.HandleFunc("/login", loginHander(formatter))
	mx.HandleFunc("/unknown", unimplementHander)
	mx.HandleFunc("/", http.RedirectHandler("/login", http.StatusPermanentRedirect).ServeHTTP)
	mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
}