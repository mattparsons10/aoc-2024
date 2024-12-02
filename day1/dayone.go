package dayone

import (
	"sort"
)

func soltionPartOne(leftList []int, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDiff := 0
	for i := 0; i < len(leftList); i++ {
		totalDiff += absolute(leftList[i] - rightList[i])
	}
	return totalDiff
}

func soltionPartTwo(leftList []int, rightList []int) int {
	rightListMap := make(map[int]int, len(rightList))
	// Get the right list counts of existance
	for _, num := range rightList {
		rightListMap[num]++
	}
	answer := 0
	for _, num := range leftList {
		if count, ok := rightListMap[num]; ok {
			answer += num * count
		}
	}
	return answer
}

func absolute(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
