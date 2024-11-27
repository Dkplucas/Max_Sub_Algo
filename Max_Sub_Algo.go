package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Create a reader to read input from the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a list of numbers separated by commas (e.g., 1,2,3):")
	input, _ := reader.ReadString('\n') // Read the first input
	input = strings.TrimSpace(input)    // Trim any whitespace

	fmt.Println("Enter a subsequence sum k:")
	input1, _ := reader.ReadString('\n') // Read the second input
	input1 = strings.TrimSpace(input1)   // Trim any whitespace

	// Split the input into parts based on commas
	parts := strings.Split(input, ",")

	// Create a slice to store the numbers
	var numbers []int

	// Convert each part to an integer and store it in the slice
	for _, part := range parts {
		part = strings.TrimSpace(part) // Trim whitespace around each number
		num, err := strconv.Atoi(part) // Convert to integer
		if err != nil {
			fmt.Printf("Error converting '%s' to a number. Please enter valid numbers only.\n", part)
			return
		}
		numbers = append(numbers, num)
	}

	// Convert input1 to an integer for subsequence sum
	k, err := strconv.Atoi(input1)
	if err != nil {
		fmt.Printf("Error converting '%s' to a number. Please enter a valid subsequence sum.\n", input1)
		return
	}

	// If the count of elements is less than k
	N := len(numbers)
	if k > N {
		fmt.Printf("Error: Subsequence size k is greater than the number of elements in the list.\n")
		return
	}

	// Separate even and odd numbers
	var evenNumbers, oddNumbers []int
	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		} else {
			oddNumbers = append(oddNumbers, num)
		}
	}

	// Sort both slices in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(evenNumbers)))
	sort.Sort(sort.Reverse(sort.IntSlice(oddNumbers)))

	// Calculate the maximum subsequence sum
	maxSum := 0
	j := 0 // Pointer for evenNumbers
	l := 0 // Pointer for oddNumbers

	for k > 0 {
		if k%2 == 1 {
			// Pick one even number if k is odd
			if j < len(evenNumbers) {
				maxSum += evenNumbers[j]
				j++
				k--
			} else {
				fmt.Println("Not enough even numbers to complete the subsequence.")
				break
			}
		} else {
			// Compare two evens with two odds
			evenPairSum := 0
			if j+1 < len(evenNumbers) {
				evenPairSum = evenNumbers[j] + evenNumbers[j+1]
			}
			var oddPairSum int
			if l+1 < len(oddNumbers) {
				oddPairSum = oddNumbers[l] + oddNumbers[l+1]
			}

			if evenPairSum >= oddPairSum && j+1 < len(evenNumbers) {
				maxSum += evenPairSum
				j += 2
			} else if l+1 < len(oddNumbers) {
				maxSum += oddPairSum
				l += 2
			} else {
				fmt.Println("Not enough pairs to complete the subsequence.")
				break
			}
			// Decrement K
			k -= 2
		}
	}

	// Print the maximum subsequence sum
	fmt.Printf("Maximum subsequence sum: %d\n", maxSum)
}
