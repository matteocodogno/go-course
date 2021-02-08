package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is
sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should
merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print
the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
 */
func main() {
	reader := bufio.NewReader(os.Stdin)
	var sli []int
	goodInput := true

	for {
		fmt.Println("Insert form some integers (positive and negative)")
		line, err := reader.ReadString('\n')
		if err != nil {
			goodInput = false
			fmt.Printf("Error occurred while reading the user input %s \n", err)
		}

		args := strings.Fields(strings.TrimSpace(line))
		if len(args) < 2 {
			goodInput = false
			fmt.Println("You should insert at least 2 integers")
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

	chunks := chunkSlice(sli, 4)
	c := make(chan []int, 4)
	var sorted []int
	for _, chunk := range chunks {
		go asyncSort(chunk, c)
	}
	for i := 0; i < 4; i ++ {
		chunk := <- c
		sorted = append(sorted, chunk...)
	}
	sort.Ints(sorted)
	fmt.Println(sorted)
}

func asyncSort(slice []int, c chan []int) {
	sort.Ints(slice)
	fmt.Println(slice)
	c <- slice
}

func chunkSlice(slice []int, chunkNumber int) [][]int {
	var chunks [][]int
	var chunk []int
	chunkSize := int(math.Round(float64(len(slice))/float64(chunkNumber)))
	j := 1
	for i := 0; i < len(slice); i ++ {
		if i != 0 && i%chunkSize == 0 && j < chunkNumber {
			j ++
			chunks = append(chunks, chunk)
			chunk = nil
		}

		chunk = append(chunk, slice[i])
	}
	chunks = append(chunks, chunk)

	return chunks
}
