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

