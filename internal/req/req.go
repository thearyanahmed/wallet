package req

import (
	"github.com/thearyanahmed/wallet/internal/res"
	"net/http"
)

const (
	GET = "GET"
	POST = "POST"
	PATCH = "PATCH"
	DELETE = "DELETE"
)

func ReturnIfInvalidMethod(req *http.Request,method string,w http.ResponseWriter) bool {
	if req.Method != method {
		res.SendError(w,"Invalid method",nil,422)
		return true
	}

	return false
}
