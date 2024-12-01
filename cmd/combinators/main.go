package main

import "github.com/arjansunar/combinators/pkg/parser"

func main() {
	println("Trying to parse a character")
	dollarParser := parser.Char('$')
	res := dollarParser([]rune("$Testing"), 0)
	println(res.Payload)
}
