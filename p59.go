// Each character on a computer is assigned a unique code and the
// preferred standard is ASCII (American Standard Code for Information
// Interchange). For example, uppercase A = 65, asterisk (*) = 42, and
// lowercase k = 107. 

// A modern encryption method is to take a text file, convert the bytes
// to ASCII, then XOR each byte with a given value, taken from a secret
// key. The advantage with the XOR function is that using the same
// encryption key on the cipher text, restores the plain text; for example,
// 65 XOR 42 = 107, then 107 XOR 42 = 65.

// For unbreakable encryption, the key is the same length as the plain
// text message, and the key is made up of random bytes. The user would
// keep the encrypted message and the encryption key in different
// locations, and without both "halves", it is impossible to decrypt the
// message.

// Unfortunately, this method is impractical for most users, so the
// modified method is to use a password as a key. If the password is
// shorter than the message, which is likely, the key is repeated
// cyclically throughout the message. The balance for this method is
// using a sufficiently long password key for security, but short enough
// to be memorable.

// Your task has been made easy, as the encryption key consists of three
// lower case characters. Using cipher.txt (right click and 'Save
// Link/Target As...'), a file containing the encrypted ASCII codes, and
// the knowledge that the plain text must contain common English words,
// decrypt the message and find the sum of the ASCII values in the original
// text.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var p = fmt.Println

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
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

// find the sum of the ASCII values in the original text
func calcSum(ascii []byte) (total int) {
	total = 0
	for _, v := range ascii {
		total += int(v)
	}
	return
}

// generate all possible keys
func generateKeys() [][]int {
	// 97 to 122
	keys := [][]int{}

	for msd := 97; msd <= 122; msd++ {
		for middle := 97; middle <= 122; middle++ {
			for lsd := 97; lsd <= 122; lsd++ {
				keys = append(keys, []int{msd, middle, lsd})
			}
		}
	}
	return keys
}

func main() {
	lines, err := readLines("./p059_cipher.txt")
	checkError(err)

	words, err := readLines("./wordsEn.txt")
	checkError(err)

	wordMap := make(map[string]bool)
	// inserts words to map
	for _, v := range words {
		wordMap[v] = true
	}

	keys := generateKeys()

	charSlice := strings.Split(lines[0], ",")
	cipherTxt := []int{}

	// convert charSlice elements from string to int
	for _, v := range charSlice {
		num, err := strconv.Atoi(v)
		checkError(err)
		cipherTxt = append(cipherTxt, num)
	}

	var buffer bytes.Buffer
	plain := ""

	for _, key := range keys {
		for k, v := range cipherTxt {
			buffer.WriteString(string(byte(v ^ key[k%3])))
		}

		plain = buffer.String()
		count := 0
		if len(strings.Split(plain, " ")) > 100 {
			for ii := 0; ii < 15; ii++ {
				if wordMap[strings.ToLower(strings.Split(plain, " ")[ii])] {
					count++
				}
			}

			if count > 5 {
				p(calcSum([]byte(plain))) // 107359
				break
			}
		}

		buffer.Reset()
	}

}
