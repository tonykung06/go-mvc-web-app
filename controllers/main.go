package controllers

import (
	"bufio"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Register(templates *template.Template) {
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	requestedFile := req.URL.Path[1:] //strip prepended forward slash
	// 	template := templates.Lookup(requestedFile + ".html")

	// 	var context interface{} = nil
	// 	switch requestedFile {
	// 	case "home":
	// 		context = viewmodels.GetHome()
	// 	}

	// 	if template != nil {
	// 		template.Execute(w, context)
	// 	} else {
	// 		w.WriteHeader(404)
	// 	}
	// })

	router := mux.NewRouter()

	hc := new(homeController) //return *homeController
	hc.template = templates.Lookup("home.html")
	router.HandleFunc("/home", hc.get)
	router.HandleFunc("/home/{id}", hc.get)

	pc := new(profileController)
	pc.template = templates.Lookup("profile.html")
	router.HandleFunc("/profile", pc.handle)

	pc2 := new(profileController)
	pc2.template = templates.Lookup("users.html")
	router.HandleFunc("/users", pc2.showUsers)

	sc := new(standLocatorController)
	sc.template = templates.Lookup("stand_locator.html")
	router.HandleFunc("/stand_locator", sc.get)
	router.HandleFunc("/api/stand_locator", sc.apiSearch)

	http.Handle("/", router)

	//more specific patterns take precedence
	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/js/", serveResource)
	http.HandleFunc("/css/", serveResource)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".mp4") {
		contentType = "video/mp4"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
