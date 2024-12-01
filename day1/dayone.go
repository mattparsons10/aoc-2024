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

func absolute(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
