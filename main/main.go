package main

import (
	"github.com/go-mvc-web-app/controllers"
	"net/http"
	"os"
	"text/template"
)

func main() {
	//template cache
	templates := populateTemplates()
	controllers.Register(templates)
	http.ListenAndServe(":8000", nil)
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
