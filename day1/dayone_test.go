package dayone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolutionPartOne(t *testing.T) {

	testCases := []struct {
		desc           string
		leftList       []int
		rightList      []int
		expectedOutput int
	}{
		{
			desc:           "example",
			leftList:       []int{3, 4, 2, 1, 3, 3},
			rightList:      []int{4, 3, 5, 3, 9, 3},
			expectedOutput: 11,
		},
		{
			desc:           "main question",
			leftList:       LeftList,
			rightList:      RightList,
			expectedOutput: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := soltionPartOne(tC.leftList, tC.rightList)

			assert.Equal(t, tC.expectedOutput, output, "solution output confirmtion")
		})
	}
}
