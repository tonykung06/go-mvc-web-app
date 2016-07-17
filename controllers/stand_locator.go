package controllers

import (
	"encoding/json"
	"github.com/go-mvc-web-app/controllers/util"
	"github.com/go-mvc-web-app/viewmodels"
	"net/http"
	"text/template"
)

type standLocatorController struct {
	template *template.Template
}

func (this *standLocatorController) get(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	vm := viewmodels.GetStandLocator()
	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}

func (this *standLocatorController) apiSearch(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	query := req.URL.Query().Get("search")
	locations := viewmodels.GetStandLocations(query)
	responseWriter.Header().Add("Content-Type", "application/json")
	data, err := json.Marshal(locations)
	if err == nil {
		responseWriter.Write(data)
	} else {
		responseWriter.WriteHeader(404)
	}
}
