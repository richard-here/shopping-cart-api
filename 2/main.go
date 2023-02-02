package main

import (
	"fmt"
	"math"
)

func editableOnce(s1, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if int(math.Abs(float64(s2Len-s1Len))) > 1 {
		return false
	}
	i, j := 0, 0
	diffs := 0
	for i < s1Len && j < s2Len {
		if s1[i] != s2[j] {
			diffs++
			if s1Len > s2Len {
				i++
			} else if s2Len > s1Len {
				j++
			} else {
				i++
				j++
			}
		} else {
			i++
			j++
		}
		if diffs > 1 {
			return false
		}
	}
	if diffs == 1 && s1Len != s2Len && (i != s1Len || j != s2Len) {
		return false
	}
	return true
}

func main() {
	inputs := [][]string{{"telkom", "telecom"}, {"telkom", "tlkom"}, {"hello", "hellio"}, {"hello", "hellia"}, {"", ""}, {"a", "b"}}
	for _, input := range inputs {
		fmt.Println(editableOnce(input[0], input[1]))
	}
}
