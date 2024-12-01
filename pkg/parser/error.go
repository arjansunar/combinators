package parser

import (
	"fmt"
	"strings"
)

type Error struct {
	Expected []string
	Input    []rune
}

func NewError(input []rune, expected ...string) *Error {
	return &Error{Expected: expected, Input: input}
}

func (e *Error) Error() string {
	return fmt.Sprintf("expected %v", strings.Join(e.Expected, ", "))
}

func (e *Error) ErrorAtChar(fullInput []rune) string {
	char := len(fullInput) - len(e.Input)
	return fmt.Sprintf("char %v: %v", char+1, e.Error())
}
