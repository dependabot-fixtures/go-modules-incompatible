package examplelib

import (
	"rsc.io/quote"
)

const version = "__VERSION__"

func Greet() string {
	return quote.Hello() + " (" + version + ")"
}
