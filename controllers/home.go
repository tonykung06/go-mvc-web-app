package controllers

import (
	"github.com/go-mvc-web-app/controllers/util"
	"github.com/go-mvc-web-app/viewmodels"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"text/template"
)

type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idRaw := vars["id"]
	var vm viewmodels.Home
	if idRaw != "" {
		id, err := strconv.Atoi(idRaw)
		if err == nil {
			vm = viewmodels.GetHomeById(id)
		} else {
			w.WriteHeader(404)
			return
		}
	} else {
		vm = viewmodels.GetHome()
	}
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}
