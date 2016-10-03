package core

// PanicOnError panics if error is not nil
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// ReturnOnError returns the err if it's not nil
func ReturnOnError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
