package controllers

import (
	"github.com/go-mvc-web-app/controllers/util"
	"github.com/go-mvc-web-app/viewmodels"
	"net/http"
	"text/template"
)

type profileController struct {
	template *template.Template
}

func (this *profileController) handle(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	vm := viewmodels.GetProfile()

	if req.Method == "POST" {
		vm.User.Email = req.FormValue("email")
		vm.User.FirstName = req.FormValue("firstname")
		vm.User.LastName = req.FormValue("lastname")
		vm.User.Stand.Address = req.FormValue("address")
		vm.User.Stand.City = req.FormValue("city")
		vm.User.Stand.State = req.FormValue("state")
		vm.User.Stand.Zip = req.FormValue("zip")
	}

	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}
