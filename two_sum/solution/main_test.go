package main_test

import (
	"fmt"
	"testing"

	main "github.com/YashYash/leetcode-go/two_sum/solution"
	"github.com/stretchr/testify/assert"
)

type Ints []int

func TestSolution(t *testing.T) {
	tests := []struct {
		tc     string
		nums   main.Nums
		target int
		expect main.Targets
		err    error
	}{
		{
			tc:     "1",
			nums:   main.Nums{5, 7, 8},
			target: 10,
			expect: main.Targets{},
			err:    main.ErrInput,
		},
		{
			tc:     "2",
			nums:   main.Nums{3, 2, 3},
			target: 6,
			expect: main.Targets{0, 2},
			err:    main.ErrInput,
		},
		{
			tc:     "3",
			nums:   main.Nums{2, 7, 8, 3},
			target: 11,
			expect: main.Targets{2, 3},
			err:    main.ErrInput,
		},
	}

	for _, test := range tests {
		output, err := main.Solution(test.tc, test.nums, test.target)
		if err != nil {
			fmt.Println("Test Case: " + test.tc + " Failed")
			fmt.Println(err)
		}
		assert.Equal(t, test.expect, output)
	}
}
