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

// 5, 4, 4, 2, 2, 2, 1
func findGiverReceiver(p []int) ([]int, bool) {
	// start from back and look for intersection
	s := p[len(p)-1] // smallest
	si := len(p) - 1 // smallest index

	for i := len(p) - 2; i >= 1; i-- {
		if p[i] != s {
			si = i + 1
			break
		}
	}

	for i := (si - 1); i >= 1; i-- {
		if (p[i] - p[si]) >= 2 {
			return []int{i, si}, true
		}
	}

	return []int{}, false
}

func nextPart(p []int, sum int) []int {
	// can something be given in p[1:]
	//   if so who should give to whom
	// if nothing can be given take 1 from p[0]

	next := make([]int, len(p))
	copy(next, p)

	if pair, ok := findGiverReceiver(p); ok {
		next[pair[0]] -= 1
		next[pair[1]] += 1
		return next
	}

	leftOver := sum - (p[0] - 1) - (len(p) - 2)
	fmt.Println("leftOver", leftOver)

	if p[0]-1 >= leftOver {
		next[0] = p[0] - 1
		next[1] = leftOver
		for i := 2; i < len(p); i++ {
			next[i] = 1
		}
		return next
	} else {
		next[0] = p[0] - 1
		next[1] = p[0] - 1

		for i := 2; i < len(p); i++ {
			next[i] = 1
		}

		left := sum - (2 * next[0]) - (len(p) - 2)
		// fmt.Println("left", left)

		// give away the left over
		for i := 2; i < len(p); i++ {
			if (left + next[i]) <= next[0] {
				next[i] += left
				break
			} else {
				next[i] = next[0]
				left = left - (next[0] - 1)
			}
		}

		return next
	}

	// return done
	return []int{}
}

func partitions(num, size int) [][]int {
	result := [][]int{}

	part := []int{num - size + 1}
	for i := 0; i < size-1; i++ {
		part = append(part, 1)
	}

	result = append(result, part)

	return result
}

func main() {
	// fmt.Println(isProductSum([]int{1, 1, 1, 1, 2, 6}))
	// fmt.Println(products(7, 4))
	// fmt.Println(partitions(20, 7))

	// ret, ok := findGiverReceiver([]int{19, 3, 3, 2, 1, 1, 1})
	// fmt.Println(ret, ok)

	// ret, ok := findGiverReceiver([]int{5, 4, 4, 2, 2, 2, 1})
	// fmt.Println(ret, ok)

	fmt.Println(nextPart([]int{5, 3, 3, 3, 2, 2, 2}, 20))
}

/*
[5 5 5 2 1 1 1]
[5 5 4 2 2 1 1]
[5 5 3 2 2 2 1]
[5 5 2 2 2 2 2]
[5 4 3 2 2 2 2]
[5 3 3 3 2 2 2]
*/
