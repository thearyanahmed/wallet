package user

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func createNewUserRequest(r *http.Request) *govalidator.Validator {
	rules := govalidator.MapData{
		"first_name" : []string{"required","between:2,12"},
		"last_name" : []string{"required","between:2,12"},
		"email":    []string{"required", "email"},
	}

	messages := govalidator.MapData{
		"first_name" : []string{"required:First name is required.","between:It should be between 2 to 12 characters."},
		"last_name" : []string{"required:Last name is required.","between:It should be between 2 to 12 characters."},
		"email":    []string{"required:Email is required.", "email:Must be a valid email."},
	}
	
	options := govalidator.Options{
		Request:         r,
		RequiredDefault: false,
		Rules:           rules,
		Messages:        messages,
		FormSize:        10,
	}

	return govalidator.New(options)
}