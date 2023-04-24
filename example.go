package main

import "fmt"

func do_something_else(accumulator int) int {
	accumulator = accumulator + 43
	return accumulator + 42
}

// do_something calls do_something_else with 'before + 42' as argument.
func do_something() {
	before := 100
	fmt.Printf("did something: %d\n", do_something_else(before+42))

	some_string := "this is a string!"
	// Type system won't allow this and the code won't compile, but parsing does not type check,
	// so our linter has no idea that this is an invalid construct:
	some_string + 42
}
