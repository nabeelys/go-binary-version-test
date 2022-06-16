package main

import "github.com/google/go-cmp/cmp"

func main() {
	cmp.Diff("abc", "abc")
}
