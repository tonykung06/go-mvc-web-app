package models

import (
	"testing"
)

func Test_ReturnsNonEmptySlice(t *testing.T) {
	users := GetUsers()

	if len(users) == 0 {
		t.Log("Non empty slice returned")
		t.Fail()
	}
}
