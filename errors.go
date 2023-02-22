package GoAzam

import "fmt"

type (
	// List of errors eg: missing a field, wrong type etc
	ErrorList map[string]any

	// Bad request error returned when there's a validation error
	BadRequestError struct {
		// Status code 400
		Status int64 `json:"status"`

		// List of errors
		Errors map[string]any `json:"errors"`

		// One or more validation errors occurred
		Title string `json:"title"`

		// See https://tools.ietf.org/html/rfc7231#section-6.5.1
		Type string `json:"type"`

		// Error ID
		TraceID string `json:"traceId"`
	}

	// Status code 417: Please provide authorization
	Unauthorized struct {
		// Error status code
		Status string `json:"status"`
		// Error message
		Message string `json:"message"`
	}
)

func (bd *BadRequestError) Error() string {
	if bd != nil {
		var errorList string
		for k, v := range bd.Errors {
			errorList = errorList + fmt.Sprintf("\t\t\t%s:\t%v\n", k, v)
		}
		return fmt.Sprintf("Bad Request Error: %d\t%s\n%s", bd.Status, bd.Title, errorList)
	}
	return ""
}

func (ua *Unauthorized) Error() string {
	if ua != nil {
		return fmt.Sprintf("Unauthorized: %s", ua.Message)
	}

	return ""
}
