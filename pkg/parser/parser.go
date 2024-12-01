package parser

func Char(char rune) Func[string] {
	return func(input []rune, idx int) ParserState[string] {
		if len(input) == 0 || input[0] != char {
			return Err[string](NewError(input, string(char)), input, idx)
		}
		return Success(string(char), input[1:], idx+1)
	}
}

// Term
func Term(t []rune) Func[string] {
	return func(input []rune, idx int) ParserState[string] {
		if len(t) > len(input) {
			return Err[string](NewError(input, string(t)), input, idx)
		}
		for i, c := range t {
			if c != input[i] {
				return Err[string](NewError(input, string(t)), input, idx)
			}
		}
		return Success(string(t), input[len(t):], idx+len(t))
	}
}

// Digit

// OneOf
// Sequence
