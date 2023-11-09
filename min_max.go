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
			return
		}

	}

	min, max, err := MinMaxSum(arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("minimum: %d, maximum: %d.\n", min, max)

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

	return arrSum - elemMax, arrSum - elemMin, nil
}

func MinMaxSumOld2(origArr []int) (int, int, error) {

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

	return min, max, nil
}

func MinMaxSumOld1(origArr [5]int) [2]int {

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

	return [2]int{minSum, maxSum}
}
