// p054.txt contains poker hands... each line in the file
// has ten cards first 5 cards are for player 1 and next 5
// for player 2... how many hands does player 1 win

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	rank int
	suit string
}

type Hand struct {
	cards []Card
}

func (h *Hand) getRanks() (ranks []int) {
	for _, c := range h.cards {
		ranks = append(ranks, c.rank)
	}
	return ranks
}

func (h *Hand) getSuits() (suits []string) {
	for _, c := range h.cards {
		suits = append(suits, c.suit)
	}
	return suits
}

func (h *Hand) String() string {
	var s string

	for i := 0; i < 5; i++ {
		if h.cards[i].rank == 10 {
			s += "10"
		} else if h.cards[i].rank == 11 {
			s += "J"
		} else if h.cards[i].rank == 12 {
			s += "Q"
		} else if h.cards[i].rank == 13 {
			s += "K"
		} else if h.cards[i].rank == 14 {
			s += "A"
		} else {
			s += strconv.Itoa(h.cards[i].rank)
		}

		if h.cards[i].suit == "S" {
			s += "♠ "
		} else if h.cards[i].suit == "H" {
			s += "♥ "
		} else if h.cards[i].suit == "D" {
			s += "♦ "
		} else {
			s += "♣ "
		}

	}

	return s
}

type ByRank []Card

func (c ByRank) Len() int {
	return len(c)
}

func (c ByRank) Less(i, j int) bool {
	return c[i].rank > c[j].rank
}

func (c ByRank) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
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

func getHands(str string) (Hand, Hand) {
	var p1Hand Hand
	var p2Hand Hand

	cardStr := strings.Split(str, " ")

	for i := 0; i < 10; i++ {
		rank := -1

		if string(cardStr[i][0]) == "T" {
			rank = 10
		} else if string(cardStr[i][0]) == "J" {
			rank = 11
		} else if string(cardStr[i][0]) == "Q" {
			rank = 12
		} else if string(cardStr[i][0]) == "K" {
			rank = 13
		} else if string(cardStr[i][0]) == "A" {
			rank = 14
		} else {
			rank, _ = strconv.Atoi(string(cardStr[i][0]))
			// if err != nil {
			// 	fmt.Println(err)
			// 	os.Exit(1)
			// }
		}

		card := Card{
			rank,
			string(cardStr[i][1]),
		}

		if i < 5 {
			p1Hand.cards = append(p1Hand.cards, card)
		} else {
			p2Hand.cards = append(p2Hand.cards, card)
		}
	}

	sort.Sort(ByRank(p1Hand.cards))
	sort.Sort(ByRank(p2Hand.cards))

	return p1Hand, p2Hand
}

func quads(h *Hand) []int {
	r := h.getRanks()
	if r[1] == r[2] && r[2] == r[3] {
		if r[0] == r[1] {
			return append([]int{7}, r...)
		}

		if r[3] == r[4] {
			return []int{7, r[1], r[2], r[3], r[4], r[0]}
		}
	}
	return nil
}

func fullHouse(h *Hand) []int {
	r := h.getRanks()
	if r[0] == r[1] && r[3] == r[4] {
		if r[1] == r[2] {
			return append([]int{6}, r...)
		}

		if r[2] == r[3] {
			return []int{6, r[2], r[3], r[4], r[0], r[1]}
		}
	}
	return nil
}

func flush(h *Hand) []int {
	s := h.getSuits()
	r := h.getRanks()
	for _, v := range s {
		if s[0] != v {
			return nil
		}
	}
	return append([]int{5}, r...)
}

// returns true if the ints in the slice are consecutive
// and in decending order
func isConsecutive(nums []int) bool {
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev-1 != nums[i] {
			return false
		}
		prev = nums[i]
	}
	return true
}

func straight(h *Hand) []int {
	r := h.getRanks()

	if isConsecutive(r[1:]) {
		if r[0] == r[1]+1 {
			return append([]int{4}, r...)
		}

		if r[0] == 14 && r[1] == 5 {
			return append([]int{4}, r[1:]...)
		}
	}

	return nil
}

// identifies a fullHouse as a set
func set(h *Hand) []int {
	r := h.getRanks()

	if r[0] == r[1] && r[1] == r[2] {
		// if c[3] == c[4] fullHouse
		return append([]int{3}, r...)
	}
	if r[1] == r[2] && r[2] == r[3] {
		return []int{3, r[1], r[2], r[3], r[0], r[4]}
	}
	if r[2] == r[3] && r[3] == r[4] {
		// if c[0] == c[1] fullHouse
		return []int{3, r[2], r[3], r[4], r[0], r[1]}
	}
	return nil
}

func twoPair(h *Hand) []int {
	c := h.getRanks()

	if c[0] == c[1] && c[2] == c[3] {
		return append([]int{2}, c...)
	}
	if c[0] == c[1] && c[3] == c[4] {
		return []int{2, c[0], c[1], c[3], c[4], c[2]}
	}
	if c[1] == c[2] && c[3] == c[4] {
		return []int{2, c[1], c[2], c[3], c[4], c[0]}
	}
	return nil
}

func pair(h *Hand) []int {
	c := h.getRanks()
	m := make(map[int]int)

	for i := 0; i < 5; i++ {
		m[c[i]] += 1
		if m[c[i]] > 1 {
			// find the kickers... not pait hands
			kickers := []int{}
			for k := 0; k < 5; k++ {
				if k != i && k != i-1 {
					kickers = append(kickers, c[k])
				}
			}

			return append([]int{1, c[i], c[i-1]}, kickers...)
		}
	}

	return nil
}

func highCard(h *Hand) []int {
	r := h.getRanks()
	return append([]int{0}, r...)
}

func determinHandValue_v2(h *Hand) []int {
	r := h.getRanks()
	m := map[int][]int{}

	for i := 0; i < 5; i++ {
		if _, ok := m[r[i]]; !ok {
			m[r[i]] = []int{i}
		} else {
			m[r[i]] = append(m[r[i]], i)
		}
	}
	// if r := []int{3, 4, 5, 6, 6} m would be
	// map[6:[3 4] 3:[0] 4:[1] 5:[2]]
	// can use m to determin quads, fullHouse, set, twoPair
	// pair, highCard by counting the value for each key
	return []int{}
}

func determinHandValue(h *Hand) []int {

	f := flush(h)
	s := straight(h)

	// straight flush
	if f != nil && s != nil {
		return append([]int{8}, f...)
	}

	// quads
	if c := quads(h); c != nil {
		return c
	}

	// fullHouse
	if c := fullHouse(h); c != nil {
		return c
	}

	// flush
	if f != nil {
		return f
	}

	// straight
	if s != nil {
		return s
	}

	// set
	if c := set(h); c != nil {
		return c
	}

	// twoPair
	if c := twoPair(h); c != nil {
		return c
	}

	// pair
	if c := pair(h); c != nil {
		return c
	}

	if c := highCard(h); c != nil {
		return c
	}

	// function shouldn't get here
	return nil
}

func winner(h1, h2 *Hand) int {
	one := determinHandValue(h1)
	two := determinHandValue(h2)

	for index, item := range one {
		if item > two[index] {
			return 1
			break
		} else if item < two[index] {
			return 2
			break
		}
	}

	// error should not get here
	// split hand
	return -1000000000

}

// you can check pair twoPair set fullHouse quads at once
// by making a map and counting like the pair function
func main() {
	lines, err := readLines("./p054_poker.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var p1Hand Hand
	var p2Hand Hand

	p1WinCount := 0

	for i := 1; i <= 1000; i++ {

		p1Hand, p2Hand = getHands(lines[i-1])

		if winner(&p1Hand, &p2Hand) == 1 {
			p1WinCount += 1
		}

		p1Hand = Hand{}
		p2Hand = Hand{}
	}

	fmt.Println(p1WinCount) // 376

}
