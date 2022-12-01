package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		return &ContentTypeNotAppJsonErrorMsg
	}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&dst)
	if err != nil {
		var unmarshalTypeError *json.UnmarshalTypeError
		switch {
		case errors.Is(err, io.EOF):
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
