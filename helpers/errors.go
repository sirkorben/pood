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

// server side errors
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

// client signup validation errors
var FirstNameMissingErrorMsg = ErrorMsg{
	ErrorDescription: "Firstname is missing",
	ErrorType:        "FIRSTNAME_FIELD_EMPTY",
}
var LastNameMissingErrorMsg = ErrorMsg{
	ErrorDescription: "Lastname is missing",
	ErrorType:        "LASTNAME_FIELD_EMPTY",
}
var EmailMissingErrorMsg = ErrorMsg{
	ErrorDescription: "Email is missing",
	ErrorType:        "EMAIL_FIELD_EMPTY",
}
var PasswordMissingErrorMsg = ErrorMsg{
	ErrorDescription: "Password is missing",
	ErrorType:        "PASSWORD_FIELD_EMPTY",
}
var PasswordTooShortErrorMsg = ErrorMsg{
	ErrorDescription: "Password is too short - 6 chars min",
	ErrorType:        "PASSWORD_TOO_SHORT",
}
var EmailIsNotValidErrorMsg = ErrorMsg{
	ErrorDescription: "Email is not valid",
	ErrorType:        "EMAIL_INVALID",
}
var EmailAlreadyTakenErrorMsg = ErrorMsg{
	ErrorDescription: "Email already taken",
	ErrorType:        "EMAIL_ALREADY_TAKEN",
}

// decoder errors
var ContentTypeNotAppJsonErrorMsg = ErrorMsg{
	ErrorDescription: "Content Type is not application/json",
	ErrorType:        "WRONG_CONTENCT_TYPE",
}

var RequestBodyIsEmptyErrorMsg = ErrorMsg{
	ErrorDescription: "Request body must not be empty.",
	ErrorType:        "REQUEST_BODY_EMPTY",
}

func ErrorResponse(w http.ResponseWriter, errorMsg ErrorMsg, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	jsonResp, _ := json.Marshal(errorMsg)
	w.Write(jsonResp)
}