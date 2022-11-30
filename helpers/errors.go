//TODO: needs proper ERROR HANDLING how to described client or server errors
package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorMsg struct {
	ErrorDescription string `json:"error_description"`
	ErrorType        string `json:"error_type"`
}

var InternalErrorMsg = ErrorMsg{
	ErrorDescription: "Internal server error",
	ErrorType:        "INTERNAL_SERVER_ERROR",
}

var NotFoundErrorMsg = ErrorMsg{
	ErrorDescription: "Page not found",
	ErrorType:        "NOT_FOUND_ERROR",
}

var BadRequestErrorMsg = ErrorMsg{
	ErrorDescription: "Bad Request",
	ErrorType:        "BAD_REQUEST_ERROR",
}

var UnauthorizedErrorMsg = ErrorMsg{
	ErrorDescription: "Restricted, because of non authorization",
	ErrorType:        "UNAUTHORIZED_ERROR",
}

var NoRecordsErrorMsg = ErrorMsg{
	ErrorDescription: "No records found",
	ErrorType:        "EMPTY_FIELD",
}

var error_responses_map = map[int]ErrorMsg{
	400: NotFoundErrorMsg,
	404: BadRequestErrorMsg,
	500: InternalErrorMsg,
}

func ErrorResponse(w http.ResponseWriter, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	jsonResp, _ := json.Marshal(error_responses_map[httpStatusCode])
	w.Write(jsonResp)
}
