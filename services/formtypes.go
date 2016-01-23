package trec

type fieldType struct {
	name string
}

type fieldValidator interface {
	validate() (valid bool, errs []error)
}

type field struct {
	name      string
	label     string
	fieldType fieldType
	validator fieldValidator
}

type singleValueField struct {
	field
	value interface{}
}

type multiValueField struct {
	field
	values []interface{}
}

type form struct {
	name   string
	fields []field
}

func (f *form) validate() (valid bool, errs []error) {
	errs = []error{}
	valid = true

	for _, field := range f.fields {
		fieldIsValid, fieldValErrs := field.validator.validate()
		if fieldIsValid != true || len(fieldValErrs) > 0 {
			valid = false
			errs = append(errs, fieldValErrs...)
		}
	}

	return valid, errs
}
