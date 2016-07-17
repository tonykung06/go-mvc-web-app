package controllers

import (
	"github.com/go-mvc-web-app/controllers/util"
	"github.com/go-mvc-web-app/models"
	"github.com/go-mvc-web-app/viewmodels"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"text/template"
)

type homeController struct {
	template      *template.Template
	loginTemplate *template.Template
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

func (this *homeController) login(w http.ResponseWriter, req *http.Request) {
	wr := util.GetResponseWriter(w, req)
	defer wr.Close()

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")
		member, err := models.GetMember(email, password)

		if err == nil {
			session, err := models.CreateSession(member)
			if err == nil {
				var cookie http.Cookie
				cookie.Name = "sessionId"
				cookie.Value = session.SessionId()

				wr.Header().Add("Set-Cookie", cookie.String())
			}
		}
	}

	wr.Header().Add("Content-Type", "text/html")
	vm := viewmodels.GetLogin()
	this.loginTemplate.Execute(wr, vm)
}
