package trec

type validateable interface {
	validate() (errs []error)
}

func errorMessages(errs []error) []string {
	errorMessages := []string{}

	for _, err := range errs {
		errorMessages = append(errorMessages, err.Error())
	}

	return errorMessages
}
