// A palindromic number reads the same both ways. The largest palindrome
// made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit
// numbers. 

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// returns the digits of a number in a slice
func isPalindrome(n int) bool {
	str := strconv.Itoa(n)

	var buffer bytes.Buffer

	for i := len(str) - 1; i >= 0; i-- {
		buffer.WriteString(string(str[i]))
	}

	if str == buffer.String() {
		return true
	} else {
		return false
	}

}

func main() {
	largest := -1

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			m := i * j
			if isPalindrome(m) && m > largest {
				largest = m
			}
		}
	}

	fmt.Println(largest) // 906609
}
