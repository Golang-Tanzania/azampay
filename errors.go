package GoAzam

import "fmt"

type (
	//Invalid detail error returned when credentials are incorrect
	InvalidDetail struct {
		// Status code 423
		Status int64 `json:"statusCode"`
		// The error message
		Message string `json:"message"`
		// This will be false
		Success bool `json:"success"`
		// Data returned from the server
		Data string `json:"null"`
	}

	// List of errors eg: missing a field, wrong type etc
	ErrorList map[string]any

	// Badrequest error returned when there's a validation error
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

// Will format the Invalid Detail Error
func (id *InvalidDetail) Error() string {
	if id != nil {
		return fmt.Sprintf("Invalid Detail Error: %d:\t%s", id.Status, id.Message)
	}

	return ""
}

// Will format the Badrequest Error
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

// Will format the Unauthorized Error
func (ua *Unauthorized) Error() string {
	if ua != nil {
		return fmt.Sprintf("Unauthorized: %s", ua.Message)
	}

	return ""
}
