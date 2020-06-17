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
		"currency_code": []string{"required","min:2","max:5","in:USD,SGD,AUD,INR,EUR,GBP"},
		"org_id" : []string{"numeric"},
		"type" :[]string{"required","in:1,2,3,4"},
	}

	messages := govalidator.MapData{
		"first_name" : []string{"required:First name is required.","between:It should be between 2 to 12 characters."},
		"last_name" : []string{"required:Last name is required.","between:It should be between 2 to 12 characters."},
		"email":    []string{"required:Email is required.", "email:Must be a valid email."},
		"currency_code": []string{"required:Currency code is required.","min:","max:","in:Supported currencies are the following.USD,SGD,AUD,INR,EUR,GBP"},
		"org_id" : []string{"numeric:Should be numeric."},
		"type" :[]string{"required:Account type is required.","in:Should be in RegularAccount (1),MerchantAccount(2),OrganizationAccount(3),AdminAccount(4)"},
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