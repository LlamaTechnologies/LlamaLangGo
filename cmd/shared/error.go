package shared

import (
	"fmt"
	antlr4 "github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

// Error defines a compile-time error
type Error struct {
	Line      int
	Column    int
	FileName  string
	Message   string
	ErrorType string
}

func (err Error) Format(f fmt.State, _ rune) {
	str := err.FileName + " : [" + strconv.Itoa(err.Line) + ":" + strconv.Itoa(err.Column) + "] : " + err.ErrorType + " : " + err.Message
	_, e := f.Write([]byte(str))
	if e != nil {
		panic(e)
	}
}

func (err *Error) toString() string {
	return err.FileName + " >> [" + fmt.Sprint(err.Line) + " : " + fmt.Sprint(err.Column) + "] :: " + err.Message
}

// CreateError creates an error obj from an antlr error
func CreateError(errorNode antlr4.ErrorNode, fileName string) Error {
	err := Error{
		Line:     errorNode.GetSourceInterval().Start,
		Column:   0,
		FileName: fileName,
		Message:  errorNode.GetText(),
	}

	return err
}
