// The fraction 49/98 is a curious fraction, as an inexperienced
// mathematician in attempting to simplify it may incorrectly
// believe that 49/98 = 4/8, which is correct, is obtained by
// cancelling the 9s.

// We shall consider fractions like, 30/50 = 3/5, to be trivial
// examples.

// There are exactly four non-trivial examples of this type of
// fraction, less than one in value, and containing two digits
// in the numerator and denominator.

// If the product of these four fractions is given in its lowest
// common terms, find the value of the denominator.

package main

import "fmt"

var p = fmt.Println

func getDigits(n int) []int {
	digits := []int{n / 10, n % 10}
	return digits
}

func commonDigits(n, d int) bool {
	num := getDigits(n)
	denom := getDigits(d)

	if num[0] == denom[0] || num[0] == denom[1] {
		return true
	}

	if num[1] == denom[0] || num[1] == denom[1] {
		return true
	}
	return false
}

func getAllFractions() [][]int {
	fractions := [][]int{}

	for i := 10; i < 99; i++ {
		for j := i + 1; j < 100; j++ {
			if commonDigits(i, j) {
				fractions = append(fractions, []int{i, j})
			}
		}

	}

	return fractions
}

func flip(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 0
	}
	return -1
}

func filter(fractions [][]int) [][]int {
	filtered := [][]int{}

	for _, v := range fractions {
		numerator := getDigits(v[0])
		denominator := getDigits(v[1])
		simplified := float64(v[0]) / float64(v[1])

		if numerator[1] == 0 && denominator[1] == 0 {
			continue
		}

		cancelIndex := [][]int{
			[]int{0, 0},
			[]int{0, 1},
			[]int{1, 0},
			[]int{1, 1}}

		for _, v := range cancelIndex {
			if numerator[v[0]] == denominator[v[1]] {
				if simplified == (float64(numerator[flip(v[0])]) /
					float64(denominator[flip(v[1])])) {
					filtered = append(filtered,
						[]int{numerator[flip(v[0])], denominator[flip(v[1])]})
				}
			}
		}
	}

	return filtered
}

func product(fractions [][]int) []int {
	p := []int{1, 1}

	for _, v := range fractions {
		p[0] *= v[0]
		p[1] *= v[1]
	}

	return p
}

func simplify(fraction []int) []int {
	frac := make([]int, len(fraction))
	copy(frac, fraction)
loop:
	for {
		for i := 2; i <= frac[0]; i++ {
			if i >= frac[0] {
				break loop
			}
			if frac[0]%i == 0 && frac[1]%i == 0 {
				frac[0] = frac[0] / i
				frac[1] = frac[1] / i
				break
			}
		}
	}

	if frac[1]%frac[0] == 0 {
		frac[1] = frac[1] / frac[0]
		frac[0] = 1
	}

	return frac
}

func main() {
	allFraction := getAllFractions()
	filtered := filter(allFraction)
	prod := product(filtered)
	simplified := simplify(prod)
	p(simplified[1])
}
