package trec

type ValidationError struct {
	FieldName    string `json:"field_name"`
	ErrorMessage string `json:"error_message"`
}

type ValidationErrors struct {
	Errors []ValidationError
}

func (errors *ValidationErrors) add(fieldName string, errorMessage string) {
	errors.Errors = append(errors.Errors, newValidationError(fieldName, errorMessage))

}

func (errors *ValidationErrors) isEmpty() bool {
	return len(errors.Errors) == 0
}

func newValidationError(fieldName string, errorMessage string) ValidationError {
	return ValidationError{
		FieldName:    fieldName,
		ErrorMessage: errorMessage,
	}
}

func newValidationErrors() ValidationErrors {
	return ValidationErrors{
		Errors: []ValidationError{},
	}
}
