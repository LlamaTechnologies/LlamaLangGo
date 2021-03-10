package shared

import "fmt"

// Error defines a compile-time error
type Error struct {
	line     int
	column   int
	fileName string
	message  string
}

func (err *Error) toString() string {
	return err.fileName + " >> [" + fmt.Sprint(err.line) + " : " + fmt.Sprint(err.column) + "] :: " + err.message
}

// CreateError creates an error obj from an antlr error
func CreateError(antlrError error) *Error {
	err := Error{
		line:     0,
		column:   0,
		fileName: "File",
		message:  "ErrorMessage",
	}

	return &err
}
