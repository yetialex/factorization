package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yetialex/factorization/internal/evaluate"
)

func FactorizeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("FactorizeHandler", time.Now())

	factorizeRequest, err := NewFactorizeRequest(r)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, []ErrorItem{{
			Code:          "number",
			Message:       "parse request parameter error",
			MessageParams: nil,
		}}, err.Error())
		return
	}
	if factorizeRequest.Number == "" {
		WriteErrorResponse(w, http.StatusBadRequest, []ErrorItem{{
			Code:          "number",
			Message:       "field is empty",
			MessageParams: nil,
		}}, ErrInvalidRequestParams.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 50*time.Second)
	defer cancel()

	factors, execTime, err := evaluate.PrimeFactorizationBigInt(ctx, factorizeRequest.Number)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, []ErrorItem{{
			Code:          "function_result",
			Message:       "Execution time: %v",
			MessageParams: []interface{}{execTime},
		}}, err.Error())
		return
	}

	WriteSuccessResponse(w, fmt.Sprintf("Execution time: %v", execTime),
		fmt.Sprintf("%s = %s", factorizeRequest.Number, factors))
}
