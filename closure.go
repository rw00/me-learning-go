package main

func runningSum() func(int) int {
	sum := 0
	return func(num int) int { // closure
		sum += num
		return sum
	}
}
