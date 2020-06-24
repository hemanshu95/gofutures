package futures

func (e *TimeoutError) Error() string{
	return e.desc
}
type TimeoutError struct{
	desc string
}