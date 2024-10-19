package main

func zeroSum(nums []int, canReuse bool) [][]int {
	result := make([][]int, 0)

	occurrences := make(map[int]int)
	for _, num := range nums {
		occurrences[num]++
	}
	for _, num := range nums {
		if (num == 0 && occurrences[num] > 1) || (num != 0 && occurrences[num] > 0 && occurrences[-num] > 0) {
			result = append(result, []int{num, -num})
			if !canReuse {
				occurrences[num]--
				occurrences[-num]--
			}
		}
	}

	return result
}
