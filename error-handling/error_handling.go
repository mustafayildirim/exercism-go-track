package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource
	var openErr error
	for {
		resource, openErr = opener()
		if openErr == nil {
			break
		}
		if _, isTransient := openErr.(TransientError); !isTransient {
			return openErr
		}
	}
	defer resource.Close()
	defer func() {
		recErr := recover()
		if recErr != nil {
			if frobErr, isFrobError := recErr.(FrobError); isFrobError {
				resource.Defrob(frobErr.defrobTag)
			}
			err, _ = recErr.(error)
		}
	}()
	resource.Frob(input)
	return nil
}
