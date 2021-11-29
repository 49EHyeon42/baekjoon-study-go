package main

func sum(a []int) int {
	var r int
	for _, item := range a {
		r += item
	}
	return r
}
