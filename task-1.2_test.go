package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMaxSum(t *testing.T) {
	testCase1 := [5]int{1, 2, 3, 4, 5}
	testCase2 := [5]int{10, 9, 8, 7, 6}
	testCase3 := [5]int{10, 20, 30, 40, 50}
	testCase4 := [5]int{40, 19, 84, 84, 0}
	testCase5 := [5]int{147, 1, -1, 78, 3}
	assert.Equal(t, [2]int{10, 14}, MinMaxSum(testCase1), "should be equal")
	assert.Equal(t, [2]int{30, 34}, MinMaxSum(testCase2), "should be equal")
	assert.Equal(t, [2]int{100, 140}, MinMaxSum(testCase3), "should be equal")
	assert.Equal(t, [2]int{143, 227}, MinMaxSum(testCase4), "should be equal")
	assert.Equal(t, [2]int{81, 229}, MinMaxSum(testCase5), "should be equal")
}
