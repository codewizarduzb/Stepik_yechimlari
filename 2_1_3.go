package main

func vote(x int, y int, z int) int {
	if x+y+z <= 1 {
		return 0
	} else {
		return 1
	}
}
