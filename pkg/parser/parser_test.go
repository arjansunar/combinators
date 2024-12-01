package parser

import "testing"

func TestChar(t *testing.T) {
	c := Char('$')
	res := c([]rune("testing"), 0)
	if res.Err == nil {
		t.Error("should have errored")
	}
	res = c([]rune("$testing"), 0)
	if res.Payload != "$" {
		t.Errorf("expected:\n\t `%v` | got: %v", "$", res.Payload)
	}
}

func TestTerm(t *testing.T) {
	// add in sub test
	t.Run("Match in single string", func(t *testing.T) {
		tt := Term([]rune("test"))
		res := tt([]rune("testing"), 0)
		if res.Payload != "test" {
			t.Errorf("expected:\n\t `%v` | got: %v", "test", res.Payload)
		}
	})

	t.Run("Match in multi string", func(t *testing.T) {
		tt := Term([]rune("tes"))
		res := tt([]rune("testing this out"), 0)
		if res.Payload != "tes" {
			t.Errorf("expected:\n\t `%v` | got: %v", "test", res.Payload)
		}
		if string(res.Remaining) != "ting this out" {
			t.Errorf("expected:\n\t `%v` | got: %v", "ting this out", string(res.Remaining))
		}
	})
}
