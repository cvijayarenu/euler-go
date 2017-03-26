package main

func f2() int {
	fi := fibGenerator()
	sum := 0
	for x := range fi {
		if x > 4000000 {
			break
		}
		if x%2 == 0 {
			sum += x
		}
	}
	return sum
}
