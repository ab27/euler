// The arithmetic sequence, 1487, 4817, 8147, in which each of
// the terms increases by 3330, is unusual in two ways:
//   (i) each of the three terms are prime, and,
//   (ii) each of the 4-digit numbers are permutations of one another.

// There are no arithmetic sequences made up of three 1-, 2-, or
// 3-digit primes, exhibiting this property, but there is one other
// 4-digit increasing sequence.

// What 12-digit number do you form by concatenating the three terms
// in this sequence?

package main

import (
	"fmt"
	"sort"
)

var p = fmt.Println

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// given a 4 digit number returns the digits in a slice
func getDigits(n int) []int {
	d1 := n / 1000
	d2 := (n - (d1 * 1000)) / 100
	d3 := (n - (d1*1000 + d2*100)) / 10
	d4 := n % 10

	return []int{d1, d2, d3, d4}
}

func convertToInt(n []int) int {
	return n[0]*1000 + n[1]*100 + n[2]*10 + n[3]
}

func group(primes map[int]struct{}) map[int][]int {
	g := make(map[int][]int)

	for k, _ := range primes {
		digits := getDigits(k)
		sort.Ints(digits)
		sorted := convertToInt(digits)

		if _, ok := g[sorted]; ok {
			g[sorted] = append(g[sorted], k)
		} else {
			g[sorted] = []int{k}
		}
	}

	return g
}

func findDifferences(nums []int) string {
	sort.Ints(nums)
	count := map[int]int{}
	difference := 0
	diffMap := make(map[int][][]int)

	for i := 0; i < len(nums)-1; i++ {
		for k := i + 1; k < len(nums); k++ {
			difference = nums[i] - nums[k]
			if difference < 0 {
				difference *= -1
			}
			if _, ok := count[difference]; ok {
				count[difference]++
				diffMap[difference] = append(diffMap[difference], []int{i, k})
			} else {
				count[difference] = 1
				diffMap[difference] = append(diffMap[difference], []int{i, k})
			}
		}
	}

	for _, v := range diffMap {
		if len(v) > 1 {
			for i := 0; i < len(v)-1; i++ {
				for j := i + 1; j < len(v); j++ {
					if v[i][1] == v[j][0] && nums[0] != 1487 {
						return fmt.Sprintf("%d%d%d", nums[v[i][0]], nums[v[i][1]],
							nums[v[j][1]])
					}
				}
			}

		}
	}

	return ""

}

func main() {
	primes := map[int]struct{}{}
	for i := 1000; i < 10000; i++ {
		if isPrime(i) {
			primes[i] = struct{}{}
		}
	}

	grouped := group(primes)
	result := ""
	for _, v := range grouped {
		result = findDifferences(v)
		if result != "" {
			break
		}
	}

	fmt.Println(result) // 296962999629
}
