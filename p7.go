// By listing the first six prime numbers:
// 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?

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

func main() {
	count := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			count++
		}

		if count == 10001 {
			fmt.Println(i) // 104743
			break
		}
	}
}
