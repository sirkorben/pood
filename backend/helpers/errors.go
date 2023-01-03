package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"pood/models"
)

// big error handling done here
// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

type ErrorMsg struct {
	ErrorDescription string `json:"error_description"`
	ErrorType        string `json:"error_type"`
}

type InfoMsg struct {
	Info string `json:"info"`
}

func (mr *ErrorMsg) Error() string {
	return mr.ErrorDescription
}

func ErrorHandler(err error, w http.ResponseWriter) {
	if err != nil {
		if errors.Is(err, models.ErrUnauthorized) {
			ErrorResponse(w, StatusForbiddenErrorMsg, http.StatusForbidden)
		} else if errors.Is(err, models.ErrNoRecord) {
			ErrorResponse(w, BadRequestErrorMsg, http.StatusBadRequest)
		} else {
			log.Println(err)
			ErrorResponse(w, UserNotActivatedErrorMsg, http.StatusUnauthorized)
		}
	}
}

func HandleDecodeJSONBodyError(err error, w http.ResponseWriter) {
	var errMsg *ErrorMsg
	if errors.As(err, &errMsg) {
		ErrorResponse(w, *errMsg, http.StatusBadRequest)
	} else {
		log.Println("helpers.DecodeJSONBody(w, r, &u)", err)
		ErrorResponse(w, InternalServerErrorMsg, http.StatusInternalServerError)
	}
}

func ErrorResponse(w http.ResponseWriter, errorMsg ErrorMsg, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	jsonResp, _ := json.Marshal(errorMsg)
	w.Write(jsonResp)
}

func InfoResponse(w http.ResponseWriter, infoMsg InfoMsg, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	jsonResp, _ := json.Marshal(infoMsg)
	w.Write(jsonResp)
}

// server side errors
var InternalServerErrorMsg = ErrorMsg{
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

var MethodNotAllowedErrorMsg = ErrorMsg{
	ErrorDescription: "Method not allowed",
	ErrorType:        "METHOD_NOT_ALLOWED",
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
var PasswordsDoNotMatchErrorMsg = ErrorMsg{
	ErrorDescription: "Passwords do not match",
	ErrorType:        "PASSWORDS_DO_NOT_MATCH",
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

// client signin validation errors
var CredentialsDontMatchErrorMsg = ErrorMsg{
	ErrorDescription: "Email and password don't match",
	ErrorType:        "CREDENTIALS_DONT_MATCH",
}
var UserNotActivatedErrorMsg = ErrorMsg{
	ErrorDescription: "User is not activated by admin",
	ErrorType:        "NOT_ACTIVATED",
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

// api calling errors
var RequestToApiFailedErrorMsg = ErrorMsg{
	ErrorDescription: "Could not set up a connection with Price Source",
	ErrorType:        "BAD_CONNECTION",
}

//admin endpoint
var StatusForbiddenErrorMsg = ErrorMsg{
	ErrorDescription: "Forbiden to enter",
	ErrorType:        "FORBIDDEN",
}

// cart responses
var EmptyCartErrorMsg = ErrorMsg{
	ErrorDescription: "You should add products to confirm cart",
	ErrorType:        "EMPTY_CART_CONFIRMATION_ERROR",
}

var OrderConfirmedInfoMsg = InfoMsg{
	Info: "Order confirmed and will be sent to email",
}
