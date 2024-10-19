package main

func uniqueSum(nums []int) int {
	sum := 0

	set := make(map[int]bool)
	for _, num := range nums {
		found := set[num]
		if !found {
			sum += num
			set[num] = true
		}
	}

	return sum
}
