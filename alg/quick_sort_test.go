package alg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

func TestFindMin(t *testing.T) {
	assert.Equal(t, 0, findMin([]int{0, 2, 2, 2}))
	assert.Equal(t, 1, findMin([]int{1, 3, 3}))
	assert.Equal(t, 1, findMin([]int{2, 3, 4, 5, 1}))
	assert.Equal(t, 1, findMin([]int{3, 1, 2}))
	assert.Equal(t, 1, findMin([]int{1, 2}))
}

func findMin(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	if size == 2 {
		if nums[0] <= nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	index := size / 2
	if nums[index] == nums[index+1] && nums[index] == nums[index-1] {
		step := 2
		checkMiddle := true
		for {
			if index+step < size {
				if nums[index] > nums[index+step] {
					checkMiddle = false
					break
				}
			}
			if index-step >= 0 {
				if nums[index] > nums[index-step] {
					checkMiddle = false
					break
				}
			}
			if index+step > size && index-step < 0 {
				break
			}
			step++
		}
		if checkMiddle {
			return nums[index]
		}
	}
	if index+1 == size {
		index--
	}
	if nums[index] > nums[index+1] {
		return nums[index+1]
	}
	for {
		checkToSize := index - 1
		if index-1 < 0 {
			checkToSize = size - 1
		}
		if nums[index] < nums[checkToSize] {
			return nums[index]
		}
		if index-1 < 0 {
			index = size
		}
		index--
	}
}

func TestThreeConsecutiveOdds(t *testing.T) {
	assert.Equal(t, false, threeConsecutiveOdds([]int{2, 6, 4, 1}))
	assert.Equal(t, true, threeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12}))
}

func threeConsecutiveOdds(arr []int) bool {
	countOfOdds := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]&1 == 1 {
			countOfOdds++
			if countOfOdds == 3 {
				return true
			}
		} else {
			countOfOdds = 0
		}
	}
	return false
}

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

func TestMiniMaxSum(t *testing.T) {
	miniMaxSum([]int32{942381765, 627450398, 954173620, 583762094, 236817490})
}

func miniMaxSum(arr []int32) {
	// Write your code here
	var skipIndex = 0
	var min int64 = 9223372036854775807
	var max int64 = 0
	for i := 0; i < len(arr); i++ {
		var sum int64 = 0
		for y, element := range arr {
			if y != skipIndex {
				sum += int64(element)
			}
		}
		if sum <= min {
			min = sum
		}
		if sum >= max {
			max = sum
		}
		skipIndex++
	}
	fmt.Println(min, max)
}

func TestTimeConversion(t *testing.T) {
	fmt.Println(timeConversion("12:05:39AM"))
	fmt.Println(timeConversion("12:45:54PM"))
}

func timeConversion(s string) string {
	// Write your code here
	var timePeriod = s[len(s)-2:]
	var hour, _ = strconv.Atoi(s[:2])
	if timePeriod == "PM" && hour == 12 {
		var result = "12" + s[2:8]
		return result
	}
	if timePeriod == "PM" {
		hour += 12
		var result = strconv.Itoa(hour) + s[2:8]
		return result
	} else if timePeriod == "AM" && hour == 12 {
		var result = "00" + s[2:8]
		return result
	} else {
		return s[:8]
	}
}

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var max int32 = 0
	for _, candle := range candles {
		if candle > max {
			max = candle
		}
	}
	var result int32 = 0
	for _, candle := range candles {
		if candle == max {
			result++
		}
	}
	return result
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
