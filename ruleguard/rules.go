package main

import "github.com/quasilyte/go-ruleguard/dsl"

func fortyTwo(m dsl.Matcher) {
	m.Match(`$x + $y`).
		Where(m["x"].Object.Is("Var") && m["y"].Value.Int() == 42).
		Report("adding 42 to variable $x")
}
