package urlquery

import (
	"errors"
	"fmt"
	"net/http"
)

// UrlqueryApiError represents an error returned by the API.
type UrlqueryApiError struct {
	StatusCode int
	Message    string
}

func (e *UrlqueryApiError) Error() string {
	return fmt.Sprintf("%s (%d)", e.Message, e.StatusCode)
}

var (
	ErrNotFound            = errors.New("not found")
	ErrForbidden           = errors.New("forbidden")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrUnexpectedStatus    = errors.New("unexpected status code")
)

func handleResponseError(resp *http.Response) error {

	switch resp.StatusCode {

	// Not Errors
	case http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNoContent:

		return nil

	// Errors
	case http.StatusNotFound:
		return &UrlqueryApiError{StatusCode: resp.StatusCode, Message: ErrNotFound.Error()}

	case http.StatusForbidden:
		return &UrlqueryApiError{StatusCode: resp.StatusCode, Message: ErrForbidden.Error()}

	case http.StatusUnauthorized:
		return &UrlqueryApiError{StatusCode: resp.StatusCode, Message: ErrUnauthorized.Error()}

	case http.StatusTooManyRequests:
		return &UrlqueryApiError{StatusCode: resp.StatusCode, Message: ErrTooManyRequests.Error()}

	case http.StatusUnprocessableEntity:
		return &UrlqueryApiError{StatusCode: resp.StatusCode, Message: ErrUnprocessableEntity.Error()}

	default:
		return &UrlqueryApiError{StatusCode: resp.StatusCode, Message: ErrUnexpectedStatus.Error()}
	}

}
