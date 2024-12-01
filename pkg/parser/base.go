package parser

import "fmt"

type Parser interface {
	parse(input string) string
}

type ParserState[T any] struct {
	Payload   T
	Err       *Error
	Remaining []rune
	Index     int
}

func (ps ParserState[T]) String() string {
	errStr := "No Error"
	if ps.Err != nil {
		errStr = (*ps.Err).Error()
	}
	return fmt.Sprintf("ParseState{\n\tPayload: %v,\n\tErr: %v,\n\tRemaining: %q}", ps.Payload, errStr, string(ps.Remaining))
}

type Func[T any] func(input []rune, idx int) ParserState[T]

func Success[T any](payload T, remaining []rune, idx int) ParserState[T] {
	return ParserState[T]{
		Payload:   payload,
		Remaining: remaining,
		Index:     idx,
	}
}

func Err[T any](err *Error, input []rune, idx int) ParserState[T] {
	return ParserState[T]{
		Err:       err,
		Remaining: input,
		Index:     idx,
	}
}

func (p *ParserState[T]) Map(fn func(T) T) {
	p.Payload = fn(p.Payload)
}

func Map[T any, R any](p ParserState[T], fn func(T) R) ParserState[R] {
	return ParserState[R]{
		Payload:   fn(p.Payload),
		Err:       p.Err,
		Remaining: p.Remaining,
		Index:     p.Index,
	}
}
