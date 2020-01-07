package main

import (
	"errors"
	"fmt"
	"time"
)

// ErrInput - If inputs are invalid
var ErrInput = errors.New("Not items in the array add up to target")

// Nums - a single node that composes the list
type Nums []int
type Targets []int
type NumToIndex map[int]int

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println("Run Time: " + elapsed.String())
	fmt.Println("------------------------------")
}

// Solution function
func Solution(tc string, nums Nums, target int) (Targets, error) {
	defer timeTrack(time.Now(), "Test Case: "+tc)
	targets := Targets{}
	numToIndex := NumToIndex{}
	for i1, num := range nums {
		compliment := target - num
		if _, ok := numToIndex[compliment]; ok {
			return Targets{numToIndex[compliment], i1}, nil
		}
		numToIndex[num] = i1
	}
	if len(targets) < 1 {
		fmt.Println(ErrInput)
		return Targets{}, ErrInput
	}
	return targets, nil
}

func main() {
	fmt.Println("Two Sum")
}
