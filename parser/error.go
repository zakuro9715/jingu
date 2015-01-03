package parser

import (
	"fmt"
)

type SyntaxError struct {
	message string
	line    int
	column  int
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("%d:%d %v", e.line, e.column, e.message)
}

func (e *SyntaxError) Line() int {
	return e.line
}

func (e *SyntaxError) Column() int {
	return e.column
}

func NewSyntaxError(message string, line, column int) *SyntaxError {
	return &SyntaxError{
		message: message,
		line:    line,
		column:  column,
	}
}
