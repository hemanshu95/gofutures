package futures

type InterruptError struct{
	desc string
}

func (e *InterruptError) Error() string{
	return e.desc
}

