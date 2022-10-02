package errormodel

type UnprocessableEntityError struct {
	Message string
}

func (ue *UnprocessableEntityError) Error() string {
	return ue.Message
}
