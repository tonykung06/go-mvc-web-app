package converters

import (
	"github.com/go-mvc-web-app/models"
	"github.com/go-mvc-web-app/viewmodels"
)

func ConvertUserToViewModel(user models.User) viewmodels.User {
	return viewmodels.User{
		Id:        user.Id(),
		Email:     user.Email(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Stand: viewmodels.Stand{
			Address: "Address 1",
			City:    "City 1",
			State:   "State 1",
			Zip:     "Zip 1",
		},
	}
}
