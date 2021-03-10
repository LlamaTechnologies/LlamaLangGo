package errors

type IrError struct {
	message string
}

func (err *IrError) Error() string {
	return err.message
}
