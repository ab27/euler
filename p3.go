// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ? 

package main

import "fmt"

var p = fmt.Println

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// takes a prime number and returns the next prime number
func nextPrime(n int) int {
	for i := n + 2; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}

func main() {
	primes := []int{2}

	// total of multiplying all the prime factors
	total := 1

	x := 600851475143

	// largest prime factor
	largest := -1

	// current index for the primes slice
	i := 0

	for {
		// check if x is divisible by a prime
		if x%primes[i] == 0 {
			x = x / primes[i]
			if primes[i] > largest {
				largest = primes[i]
			}
			total *= primes[i]
			i = 0
		} else {
			if len(primes) <= i+1 {
				primes = append(primes, nextPrime(i))
			}
			i++
		}

		if total == 600851475143 {
			break
		}
	}

	fmt.Println(largest) // 6857

}
