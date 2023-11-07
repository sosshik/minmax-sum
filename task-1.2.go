package main

import (
	"fmt"
	"slices"
)

// Function takes array with length of 5 and returns array where the first element is minimum sum and second is maximum sum.
// Also it prints min and max sum to stdout
func MinMaxSum(origArr [5]int) [2]int {
	// sumsArr is an array of future sums, it's of length 5 beacause it's the most possible number of combinations
	var sumsArr [5]int

	//tmpSum is temporary variable to store sum
	var tmpSum int

	for i := 0; i < 5; i++ {

		tmpSum = 0

		for j, v := range origArr {
			if i == j {
				continue
			}
			tmpSum += v
		}

		sumsArr[i] = tmpSum
	}

	minSum := slices.Min(sumsArr[0:5])
	maxSum := slices.Max(sumsArr[0:5])

	fmt.Printf("minimum: %v, maximum: %v.\n", minSum, maxSum)

	return [2]int{minSum, maxSum}
}
