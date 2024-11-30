package main

import "fmt"

type Parser interface {
	parse(input string) string
}

type ParseState[T any] struct {
	Payload   T
	Err       *error
	Remaining []rune
}

func (ps ParseState[T]) String() string {
	errStr := "No Error"
	if ps.Err != nil {
		errStr = (*ps.Err).Error()
	}
	return fmt.Sprintf("ParseState{\n\tPayload: %v,\n\tErr: %v,\n\tRemaining: %q}", ps.Payload, errStr, string(ps.Remaining))
}

type Func[T any] func(input []rune) ParseState[T]

func Success[T any](payload T, remaining []rune) ParseState[T] {
	return ParseState[T]{
		Payload:   payload,
		Remaining: remaining,
	}
}

func Err[T any](err *error, remaining []rune) ParseState[T] {
	return ParseState[T]{
		Err:       err,
		Remaining: remaining,
	}
}

func Char(char rune) Func[string] {
	return func(input []rune) ParseState[string] {
		if len(input) == 0 || input[0] != char {
			err := fmt.Errorf("expected %v got %v", char, string(input))
			return Err[string](&err, input)
		}

		return Success(string(char), input[1:])
	}
}

func main() {
	println("Trying to parse a character")
	dollarParser := Char('$')
	res := dollarParser([]rune("$Testing"))
	println(res.Payload)
}
