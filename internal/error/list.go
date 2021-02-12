package error

type List []error

func (e *List) ReturnOrNil() error {
	if e.Empty() {
		return nil
	}
	return e
}

func (e *List) Empty() bool {
	return len(*e) == 0
}

func (e *List) Add(err error) {
	if err != nil {
		*e = append(*e, err)
	}
}

func (e *List) Error() string {
	errString := ""
	for _, err := range *e {
		errString += err.Error()
		errString += "\n"
	}
	return errString
}
