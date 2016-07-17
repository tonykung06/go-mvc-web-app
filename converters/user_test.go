package converters

import (
	"github.com/go-mvc-web-app/models"
	"testing"
)

func Test_ConvertUserToViewModel(t *testing.T) {
	user := models.User{}
	user.SetId(1)
	user.SetEmail("test@gmail.com")
	user.SetFirstName("test first name")
	user.SetLastName("test last name")

	userViewModel := ConvertUserToViewModel(user)

	if userViewModel.Email != user.Email() {
		t.Log("Email not converted properly")
		t.FailNow()
	}

	if userViewModel.FirstName != user.FirstName() {
		t.Log("FirstName not converted properly")
		t.FailNow()
	}

	if userViewModel.LastName != user.LastName() {
		t.Log("LastName not converted properly")
		t.FailNow()
	}

	if userViewModel.Id != user.Id() {
		t.Log("Id not converted properly")
		t.FailNow()
	}
}
