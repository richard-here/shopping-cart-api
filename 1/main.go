package main

import (
	"fmt"
	"math"
)

func getDenominations() []int {
	return []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
}

func calculateDenominationsNeeded(money int) map[int]int {
	denominations := getDenominations()
	length := len(denominations)
	last := float64(denominations[length-1])
	moddedM := int(math.Ceil(float64(money)/last) * last)
	answer := make(map[int]int)
	for i := 0; i < length; i++ {
		q := moddedM / denominations[i]
		r := moddedM % denominations[i]
		if q > 0 {
			answer[denominations[i]] = q
		}
		moddedM = r
	}
	return answer
}

func main() {
	inputs := []int{145000, 2050, 2001, 1353412, 253999, 3120673}
	for _, input := range inputs {
		fmt.Println(calculateDenominationsNeeded(input))
	}
}
