package speedtest

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkSpeedMapWithAllocateNeedSize(b *testing.B) {
	var size = []int{100, 1000, 10000, 100000, 100000, 1000000, 10000000}
	for _, sizeEl := range size {
		b.Run(strconv.Itoa(sizeEl), func(b *testing.B) {
			number := GenerateNumber(sizeEl)
			var mapRes = make(map[int]int, len(number))
			for _, el := range number {
				mapRes[el] = el
			}
		})
	}
}

func BenchmarkSpeedMapWithoutAllocateNeedSize(b *testing.B) {
	var size = []int{100, 1000, 10000, 100000, 100000, 1000000, 10000000}
	for _, sizeEl := range size {
		b.Run(strconv.Itoa(sizeEl), func(b *testing.B) {
			number := GenerateNumber(sizeEl)
			var mapRes = make(map[int]int, 0)
			for _, el := range number {
				mapRes[el] = el
			}
		})
	}
}

func GenerateNumber(value int) []int {
	var res = make([]int, value)
	for i := 0; i < value; i++ {
		res[i] = rand.Int()
	}
	return res
}
