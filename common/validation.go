package common

type ValidationError struct {
	FieldName    string `json:"field_name"`
	ErrorMessage string `json:"error_message"`
}

type ValidationErrors struct {
	Errors []ValidationError
}

func (errors *ValidationErrors) Add(fieldName string, errorMessage string) {
	errors.Errors = append(errors.Errors, NewValidationError(fieldName, errorMessage))
}

func (errors *ValidationErrors) IsEmpty() bool {
	return len(errors.Errors) == 0
}

func NewValidationError(fieldName string, errorMessage string) ValidationError {
	return ValidationError{
		FieldName:    fieldName,
		ErrorMessage: errorMessage,
	}
}

func NewValidationErrors() ValidationErrors {
	return ValidationErrors{
		Errors: []ValidationError{},
	}
}
