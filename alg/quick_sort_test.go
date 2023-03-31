package alg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

func TestInt(t *testing.T) {
	var ar = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Println(aVeryBigSum(ar))
}

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var result int64
	for _, el := range ar {
		result += el
	}
	return result
}

func TestStaircase(t *testing.T) {
	staircase(5)
}

func staircase(n int32) {
	// Write your code here
	var ar = make([]string, n)
	var last int = len(ar) - 1
	for i := 0; i < len(ar); i++ {
		ar[i] = " "
	}
	for i := 0; i < int(n); i++ {
		ar[last] = "#"
		last--
		fmt.Println(strings.Join(ar, ""))
	}
}

func miniMaxSum(arr []int32) {
	// Write your code here

}

func TestDiagonalDifference(t *testing.T) {
	var ar = [][]int32{
		{-10, 3, 0, 5, -4},
		{2, -1, 0, 2, -8},
		{9, -2, -5, 6, 0},
		{9, -7, 4, 8, -2},
		{3, 7, 8, -5, 0},
	}
	fmt.Println(diagonalDifference(ar))
}

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var firstDiagonal = 0
	var secondDiagonal = len(arr[0]) - 1
	var one int32 = 0
	var two int32 = 0
	for i := range arr {
		one += arr[i][firstDiagonal]
		firstDiagonal++
		two += arr[secondDiagonal][i]
		secondDiagonal--
	}
	if one < two {
		return two - one
	}
	return one - two
}

func TestPlusMinus(t *testing.T) {
	var ar = []int32{0, 0, -1, 1, 1}
	plusMinus(ar)
}

func plusMinus(arr []int32) {
	// Write your code here
	var countPositive int32 = 0
	var countNegative int32 = 0
	var countZero int32 = 0
	for _, element := range arr {
		if element == 0 {
			countZero++
		}
		if element > 0 {
			countPositive++
		}
		if element < 0 {
			countNegative++
		}
	}
	var size = float32(len(arr))
	fmt.Printf("%.6f\n", float32(countPositive)/size)
	fmt.Printf("%.6f\n", float32(countNegative)/size)
	fmt.Printf("%.6f\n", float32(countZero)/size)
}

func compareTriplets(a []int32, b []int32) []int32 {
	var bob int32 = 0
	var alice int32 = 0
	for i := range a {
		if a[i] > b[i] {
			alice++
		} else if a[i] < b[i] {
			bob++
		}
	}
	return []int32{alice, bob}
}

func TestScope(t *testing.T) {
	runtime.GOMAXPROCS(8)

	done := true
	fmt.Println(done)

	go func() {
		fmt.Println(done)
		done = false
		fmt.Println(done)
	}()

	for !done {
	}
	fmt.Println("finished")
}

func TestAtomi(t *testing.T) {
	var counter int
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			counter++
			wg2.Done()
		}()
	}
	wg2.Wait()
	fmt.Println(counter)

	var wg sync.WaitGroup
	var atomicCounter int32
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&atomicCounter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(atomicCounter)

}

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
