package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument
and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent
elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should
take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap()
function should return nothing, but it should swap the contents of the slice in position i with the contents in
position i+1.
*/
func main() {
	const sliceSize = 10
	var sli []int
	goodInput := true
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Insert form 1 to 10 integers (positive and negative)")
		line, err := reader.ReadString('\n')
		if err != nil {
			goodInput = false
			fmt.Printf("Error occurred while reading the user input %s \n", err)
		}

		args := strings.Fields(strings.TrimSpace(line))
		if len(args) < 1 || len(args) > 10 {
			goodInput = false
			fmt.Println("You should insert from 1 to 10 integers")
		}

		for _, arg := range args {
			num, err := strconv.Atoi(arg)
			if err != nil {
				goodInput = false
				fmt.Printf("%s is not a number, Try again\n", arg)
				break
			}

			sli = append(sli, num)
		}

		if goodInput {
			break
		}

		sli = []int{}
		goodInput = true
	}

	BubbleSort(sli)
	fmt.Println(sli)
}

func BubbleSort(sli []int) {
	length := len(sli)

	for i := 0; i < length; i ++ {
		for j := 1; j < (length-i); j ++ {
			if sli[j-1] > sli[j] {
				Swap(sli, j)
			}
		}
	}
}

func Swap(sli []int, index int) {
	temp := sli[index-1]
	sli[index-1] = sli[index]
	sli[index] = temp
}