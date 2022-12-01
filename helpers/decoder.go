//TODO: Handle Errors !!!!!!!!!!!!!!!!
package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (mr *ErrorMsg) Error() string {
	return mr.ErrorDescription
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		// errDescription := "Content Type is not application/json"
		// errType := "WRONG_CONTENCT_TYPE"
		// return &ErrorMsg{ErrorDescription: errDescription, ErrorType: errType}
		return &ContentTypeNotAppJsonErrorMsg
	}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&dst)
	if err != nil {
		// big error handling done here
		// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
		var unmarshalTypeError *json.UnmarshalTypeError
		switch {

		case errors.Is(err, io.EOF):
			// errDescription := "Request body must not be empty."
			// errType := "REQUEST_BODY_EMPTY"
			// return &ErrorMsg{ErrorDescription: errDescription, ErrorType: errType}
			return &RequestBodyIsEmptyErrorMsg

		case errors.As(err, &unmarshalTypeError):
			errDescription := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			errType := "INVALID_VALUE_FOR_FIELD"
			return &ErrorMsg{ErrorDescription: errDescription, ErrorType: errType}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			errDescription := "Request body contains unknown field " + fieldName
			errType := "UNKNOWN_FIELD"
			return &ErrorMsg{ErrorDescription: errDescription, ErrorType: errType}

		default:
			return err
		}

	}
	return nil
}
