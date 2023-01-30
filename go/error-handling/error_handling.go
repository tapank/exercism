package erratum

func Use(opener ResourceOpener, input string) (e error) {
	resource, err := opener()
	for {
		if err == nil {
			defer resource.Close()
			defer func() {
				if err := recover(); err != nil {
					if df, ok := err.(FrobError); ok {
						resource.Defrob(df.defrobTag)
					}
					e = err.(error)
				}
			}()
			resource.Frob(input)
			break
		} else {
			switch err.(type) {
			case TransientError:
				resource, err = opener()
				continue
			default:
				e = err
			}
			break
		}
	}
	return
}
