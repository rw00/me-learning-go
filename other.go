package main

type cost struct {
	day   int
	value float64
}

// pointer receivers are more common
// no need to create a pointer to call the function
// Example:

// c := cost{day: 1, value: 10.0}
// func (c *cost) updateCost() { c.value = c.value * 1.1 }
// c.updateCost() // will work with a pointer or value

func getCostsByDay(costs []cost) []float64 {
	costsByDay := make([]float64, 0)
	for _, cost := range costs {
		for len(costsByDay) <= cost.day {
			costsByDay = append(costsByDay, 0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for r := 0; r < rows; r++ {
		matrix[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			matrix[r][c] = r * c
		}
	}
	return matrix
}

func getNameCountsByFirstChar(names []string) map[string]map[string]int {
	nameCounts := make(map[string]map[string]int)
	for _, name := range names {
		firstChar := string(name[0])
		if nameCounts[firstChar] == nil {
			nameCounts[firstChar] = make(map[string]int)
		}
		nameCounts[firstChar][name]++
	}
	return nameCounts
}
