package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {

	cmdArgs := os.Args[1:]
	arr := make([]int, len(cmdArgs))
	var err error

	for i, v := range cmdArgs {
		arr[i], err = strconv.Atoi(v)

		if err != nil {
			fmt.Println("Unable to convert cmd args to int")
		}

	}

	MinMaxSum(arr)

}

func MinMaxSum(origArr []int) (int, int, error) {

	if len(origArr) != 5 {
		return 0, 0, errors.New("array must be of length")
	}

	var arrSum int
	elemMin := origArr[0]
	elemMax := origArr[0]

	for _, v := range origArr {
		if elemMax < v {
			elemMax = v
		}
		if elemMin > v {
			elemMin = v
		}
		arrSum += v
	}

	min := arrSum - elemMax
	max := arrSum - elemMin
	fmt.Printf("minimum: %v, maximum: %v.\n", min, max)

	return min, max, nil
}

func MinMaxSumOld(origArr [5]int) [2]int {

	var sumsArr [5]int

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
