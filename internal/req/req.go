package req

import (
	"github.com/thearyanahmed/wallet/internal/res"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

const (
	GET = "GET"
	POST = "POST"
	PATCH = "PATCH"
	DELETE = "DELETE"
)

type Request struct {}

// Deprecated
// Use mux's router...().Method(type) instead
func ReturnIfInvalidMethod(req *http.Request,method string,w http.ResponseWriter) bool {
	if req.Method != method {
		res.SendError(w,"Invalid method",nil,422)
		return true
	}

	return false
}

func (req *Request) Validate(r *http.Request,w http.ResponseWriter,validatorCallback func(*http.Request) *govalidator.Validator ) bool {

	validator := validatorCallback(r)

	e := validator.Validate()

	if len(e) == 0 {
		return true
	}

	res.SendError(w,"Invalid data.",e,422)
	return false
}

func (req *Request) ValidatedFormData (r *http.Request, keys []string) map[string]string {
	r.ParseForm()

	return validated(r.PostForm,keys)
}

func (req *Request) ValidatedQuery (r *http.Request, keys []string) map[string]string {
	return validated(r.URL.Query(),keys)
}

func validated(requestParams map[string][]string, keys []string) map[string]string  {
	validated := map[string]string{}

	if len(keys) == 0 {
		for key, value := range requestParams {
			validated[key] = value[0]
		}
	} else {
		for _, key := range keys {
			validated[key] = requestParams[key][0]
		}
	}

	return validated
}