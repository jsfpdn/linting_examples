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
}
