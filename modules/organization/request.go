package organization

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func createNewOrganizationRequest(r *http.Request) *govalidator.Validator {
	rules := govalidator.MapData{
		"user_id" : []string{"required","numeric"},
		"name" : []string{"required","between:2,50"},
		"currency_code": []string{"required","min:2","max:5","in:USD,SGD,AUD,INR,EUR,GBP"},
	}

	messages := govalidator.MapData{
		"user_id" : []string{"required:User id is required.","numeric:User id must be a unsigned numeric value."},
		"name" : []string{"required:The name of the organization is also required.","between:Name's length can not be samller than 2 or larger than 50."},
		"currency_code": []string{"required:Currency code is required.","min:","max:","in:Supported currencies are the following.USD,SGD,AUD,INR,EUR,GBP"},
	}

	options := govalidator.Options{
		Request:         r,
		RequiredDefault: false,
		Rules:           rules,
		Messages:        messages,
		FormSize:        1000,
	}

	return govalidator.New(options)
}