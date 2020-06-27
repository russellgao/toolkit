package exception

type ParameterError struct {
	ErrorCode int
	ErrorMsg  string
}

func (pe *ParameterError) Error() string {
	return pe.ErrorMsg
}
