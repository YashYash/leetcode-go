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
type TrackedIndex map[float64]bool

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println("Run Time: " + elapsed.String())
	fmt.Println("------------------------------")
}

// Solution function
func Solution(tc string, nums Nums, target int) (Targets, error) {
	defer timeTrack(time.Now(), "Test Case: "+tc)
	fmt.Println(nums)
	targets := Targets{}
	fmt.Println("---------------")
	for i1, num := range nums {
		for i2, num2 := range nums {
			fmt.Println(num, num2)
			fmt.Println(i1, i2)
			if i1 != i2 && num+num2 == target {
				targets = append(targets, Targets{i1, i2}...)
				return targets, nil
			}
		}
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
