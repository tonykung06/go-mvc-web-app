package main

import (
	// "fmt"
	"bufio"
	"github.com/go-mvc-web-app/viewmodels"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func main() {
	//template cache
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:] //strip prepended forward slash
		template := templates.Lookup(requestedFile + ".html")

		var context interface{} = nil
		switch requestedFile {
		case "home":
			context = viewmodels.GetHome()
		}

		if template != nil {
			template.Execute(w, context)
		} else {
			w.WriteHeader(404)
		}
	})
	//more specific patterns take precedence
	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/js/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.ListenAndServe(":8000", nil)
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

func populateTemplates() *template.Template {
	result := template.New("templates")
	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := *new([]string) //the same as []string{}
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			templatePaths = append(templatePaths, basePath+"/"+pathInfo.Name())
		}
	}
	result.ParseFiles(templatePaths...)
	return result
}
