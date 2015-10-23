// The sum of the squares of the first ten natural numbers is,

// 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is,

// (1 + 2 + ... + 10)^2 = 552 = 3025
// Hence the difference between the sum of the squares of the
// first ten natural numbers and the square of the sum is
// 3025 − 385 = 2640.

// Find the difference between the sum of the squares of the
// first one hundred natural numbers and the square of the sum.

package main

import "fmt"

func main() {
	sumOfSquares := 0
	squareOfSums := 0
	sums := 0

	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		sums += i
	}

	squareOfSums = sums * sums

	fmt.Println(squareOfSums - sumOfSquares) // 25164150

}