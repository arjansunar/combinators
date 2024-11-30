package main

import "fmt"

type Parser interface {
	parse(input string) string
}

type ParseState struct {
	Payload   interface{}
	Err       *error
	Remaining []rune
}

func (ps ParseState) String() string {
	errStr := "No Error"
	if ps.Err != nil {
		errStr = (*ps.Err).Error()
	}
	return fmt.Sprintf("ParseState{\n\tPayload: %v,\n\tErr: %v,\n\tRemaining: %q}", ps.Payload, errStr, string(ps.Remaining))
}

type Func func(input []rune) ParseState

func Success(payload interface{}, remaining []rune) ParseState {
	return ParseState{
		Payload:   payload,
		Remaining: remaining,
	}
}

func Err(err *error, remaining []rune) ParseState {
	return ParseState{
		Err:       err,
		Remaining: remaining,
	}
}

func Char(char rune) Func {
	return func(input []rune) ParseState {
		if len(input) == 0 || input[0] != char {
			err := fmt.Errorf("expected %v got %v", char, string(input))
			return Err(&err, input)
		}

		return Success(string(char), input[1:])
	}
}

func main() {
	println("Trying to parse a character")
	dollarParser := Char('$')
	println(dollarParser([]rune("$Testing")).String())
}
