package parser

func Char(char rune) Func[string] {
	return func(input []rune, idx int) ParserState[string] {
		if len(input) == 0 || input[0] != char {
			return Fail[string](NewError(input, string(char)), input, idx)
		}
		return Success(string(char), input[1:], idx+1)
	}
}

// Term
func Term(t []rune) Func[string] {
	return func(input []rune, idx int) ParserState[string] {
		if len(t) > len(input) {
			return Fail[string](NewError(input, string(t)), input, idx)
		}
		for i, c := range t {
			if c != input[i] {
				return Fail[string](NewError(input, string(t)), input, idx)
			}
		}
		return Success(string(t), input[len(t):], idx+len(t))
	}
}

// Digit

// OneOf
func OneOf[T any](parsers ...Func[T]) Func[T] {
	return func(input []rune, idx int) ParserState[T] {
		var res ParserState[T]
		for _, p := range parsers {
			res = p(input, idx)
			if res.Err == nil {
				return res
			}
		}
		return Fail[T](NewError(input, "OneOf"), input, idx)
	}
}

// Sequence
