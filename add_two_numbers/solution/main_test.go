package main_test

import (
	"fmt"
	"testing"

	main "github.com/YashYash/leetcode-go/add_two_numbers/solution"
	"github.com/stretchr/testify/assert"
)

type Ints []int

func getListNode(ints Ints) main.LinkedListNode {
	var ll main.LinkedListNode
	for i := 0; i < len(ints); i++ {
		ll.Append(ints[i])
	}
	return ll
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1     main.LinkedListNode
		l2     main.LinkedListNode
		tc     string
		expect main.LinkedListNode
		err    error
	}{
		{
			l1:     getListNode(Ints{0, 0, 1}),
			l2:     getListNode(Ints{0, 1, 1}),
			tc:     "1",
			expect: getListNode(Ints{0, 1, 2}),
			err:    main.ErrInput,
		},
		{
			l1:     getListNode(Ints{1}),
			l2:     getListNode(Ints{1}),
			tc:     "2",
			expect: getListNode(Ints{2}),
			err:    main.ErrInput,
		},
		{
			l1:     getListNode(Ints{2, 4, 3}),
			l2:     getListNode(Ints{5, 6, 4}),
			tc:     "3",
			expect: getListNode(Ints{7, 0, 8}),
			err:    main.ErrInput,
		},
		{
			l1:     getListNode(Ints{0, 1, 0}),
			l2:     getListNode(Ints{0, 1, 2}),
			tc:     "4",
			expect: getListNode(Ints{0, 2, 2}),
			err:    main.ErrInput,
		},
		{
			l1:     getListNode(Ints{9, 9, 9}),
			l2:     getListNode(Ints{9, 9, 9}),
			tc:     "5",
			expect: getListNode(Ints{8, 9, 9}),
			err:    main.ErrInput,
		},
	}

	for _, test := range tests {
		output, err := main.AddTwoNumbers(test.l1, test.l2, test.tc)
		if err != nil {
			fmt.Println("AddTwoNumbers - Test Case: " + test.tc + " Failed")
			fmt.Println(err)
		}
		assert.Equal(t, test.expect, output)
	}
}
