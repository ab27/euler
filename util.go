import (
	"bufio"
	"os"
)

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// given a 4 digit number n, returns a slice of pumutations
// of the digits
func permutations(n int) []int {
	result := []int{}
	digits := getDigits(n)

	for _, v1 := range digits {
		for _, v2 := range digits {
			if v2 == v1 {
				continue
			}
			for _, v3 := range digits {
				if v3 == v1 || v3 == v2 {
					continue
				}
				for _, v4 := range digits {
					if v4 == v1 || v4 == v2 || v4 == v3 {
						continue
					}
					num := v1*1000 + v2*100 + v3*10 + v4
					result = append(result, num)
				}
			}
		}

	}
	return result
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
