package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type FactorizeRequest struct {
	Number string `json:"number"`
}

func NewFactorizeRequest(r *http.Request) (*FactorizeRequest, error) {
	fr := FactorizeRequest{}
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		fr.Number = vars["number"]
	case "POST":
		if fillParametersFromBody(r, &fr) != nil {
			return nil, ErrInvalidRequestParams
		}
	default:
		return nil, fmt.Errorf("unsupported type: %s", r.Method)
	}
	return &fr, nil
}
