// A natural number, N, that can be written as the sum and product
// of a given set of at least two natural numbers, {a1, a2, ... ,
// ak} is called a product-sum number: N = a1 + a2 + ... + ak =
// a1 × a2 × ... × ak.

// For example, 6 = 1 + 2 + 3 = 1 × 2 × 3.

// For a given set of size, k, we shall call the smallest N with
// this property a minimal product-sum number. The minimal product
// -sum numbers for sets of size, k = 2, 3, 4, 5, and 6 are as
// follows.

// k=2: 4 = 2 × 2 = 2 + 2
// k=3: 6 = 1 × 2 × 3 = 1 + 2 + 3
// k=4: 8 = 1 × 1 × 2 × 4 = 1 + 1 + 2 + 4
// k=5: 8 = 1 × 1 × 2 × 2 × 2 = 1 + 1 + 2 + 2 + 2
// k=6: 12 = 1 × 1 × 1 × 1 × 2 × 6 = 1 + 1 + 1 + 1 + 2 + 6

// Hence for 2≤k≤6, the sum of all the minimal product-sum numbers
// is 4+6+8+12 = 30; note that 8 is only counted once in the sum.

// In fact, as the complete set of minimal product-sum numbers for
// 2≤k≤12 is {4, 6, 8, 12, 15, 16}, the sum is 61.

// What is the sum of all the minimal product-sum numbers for
// 2≤k≤12000?

package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isProductSum(nums []int) bool {
	product := 1
	sum := 0

	for _, v := range nums {
		product *= v
		sum += v
	}

	return sum == product
}

// returns slices of 'size' numbers when multiplied give num
func products(num, size int) [][]int {

	result := [][]int{}

	if isPrime(num) {
		n := []int{}
		for i := 0; i < size-1; i++ {
			n = append(n, 1)
		}
		return append(result, append(n, num))
	}

	return result
}

func main() {
	fmt.Println(isProductSum([]int{1, 1, 1, 1, 2, 6}))
	fmt.Println(products(7, 4))
}