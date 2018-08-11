package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	BadRequestMethod = errors.New(http.StatusText(http.StatusMethodNotAllowed))
	InternalError    = errors.New(http.StatusText(http.StatusInternalServerError))

	NoJSONBody   = errors.New("Unable to decode JSON")
	InvalidEmail = errors.New("Invalid Email")

	VerificationNotFound = errors.New("Invalid Verification Code")
	UserNotFound         = errors.New("User does not exist")
	RouteNotFound        = errors.New("Route not found")
)

func codeMap() map[error]int {
	return map[error]int{
		BadRequestMethod:     http.StatusMethodNotAllowed,
		NoJSONBody:           http.StatusBadRequest,
		InvalidEmail:         http.StatusBadRequest,
		VerificationNotFound: http.StatusNotFound,
		UserNotFound:         http.StatusNotFound,
		RouteNotFound:        http.StatusNotFound,
		InternalError:        http.StatusInternalServerError,
	}
}

func getCode(e error) int {
	if code, ok := codeMap()[e]; ok {
		return code
	}
	return http.StatusInternalServerError
}

type errorResponse struct {
	Error string
}

func Write(w http.ResponseWriter, e error) {
	w.WriteHeader(getCode(e))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&errorResponse{Error: e.Error()})
}
