package parser

import "testing"

func Expect[T comparable](t *testing.T, left T, right T) {
	if left != right {
		t.Errorf("expected:\n\t `%v` | got: %v", right, left)
	}
}

func TestChar(t *testing.T) {
	c := Char('$')
	res := c([]rune("$testing"), 0)
	Expect(t, res.Payload, "$")
}

func TestTerm(t *testing.T) {
	// add in sub test
	t.Run("Match in single string", func(t *testing.T) {
		tt := Term([]rune("let"))
		res := tt([]rune("let x = 0"), 0)
		Expect(t, res.Payload, "let")
	})

	t.Run("Match in multi string", func(t *testing.T) {
		tt := Term([]rune("tes"))
		res := tt([]rune("testing this out"), 0)
		Expect(t, res.Payload, "tes")
		Expect(t, string(res.Remaining), "ting this out")
	})

	t.Run("Map result", func(t *testing.T) {
		tt := Term([]rune("tes"))
		res := tt([]rune("testing this out"), 0)
		updated := Map(res, func(s string) int {
			return 1
		})
		Expect(t, updated.Payload, 1)
	})
}
