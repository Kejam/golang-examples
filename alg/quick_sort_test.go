package alg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var unsortedArray = []int{5, 6, 7, 3, 2, 6, 43, 4, 56, 6}
	var countOfSwap = 0
	QuickSort(unsortedArray, 0, len(unsortedArray)-1, &countOfSwap)
	assert.Equal(t, 2, unsortedArray[0])
	assert.Equal(t, 56, unsortedArray[len(unsortedArray)-1])
	println("Value count of swap: ", countOfSwap)
}

func QuickSort(array []int, start int, end int, countOfSwap *int) {
	if start < end {
		var p = Partition(array, start, end, countOfSwap)
		QuickSort(array, start, p-1, countOfSwap)
		QuickSort(array, p+1, end, countOfSwap)
	}
}

func Partition(array []int, start int, end int, countOfSwap *int) int {
	var pivot = array[end]
	var i = start - 1
	for j := start; j < end; j++ {
		if array[j] <= pivot {
			i++
			var temp = array[i]
			array[i] = array[j]
			array[j] = temp
			*countOfSwap++
		}
	}
	var temp = array[i+1]
	array[i+1] = array[end]
	array[end] = temp
	*countOfSwap++
	return i + 1
}
