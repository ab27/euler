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
	cards    []Card
	strength int // sf: 8, 4ofakind: 7, fullhouse:6...
	score    int // sum of the ranks of the hand
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

func quads(h Hand) bool {
	r := h.getRanks()
	if r[1] == r[2] && r[2] == r[3] {
		if r[0] == r[1] || r[3] == r[4] {
			return true
		}
	}
	return false
}

func fullHouse(h Hand) bool {
	r := h.getRanks()
	if r[0] == r[1] && r[3] == r[4] {
		if r[2] == r[1] || r[2] == r[3] {
			return true
		}
	}
	return false
}

func flush(h Hand) bool {
	suits := h.getSuits()
	for _, v := range suits {
		if suits[0] != v {
			return false
		}
	}
	return true
}

func straight(h Hand) bool {
	cardRank := h.getRanks()[0]
	for i := 1; i < 5; i++ {
		if cardRank-1 != h.cards[i].rank && h.cards[i].rank != 5 {
			return false
		}
		cardRank = h.cards[i].rank
	}
	return true
}

func set(h Hand) bool {
	return false
}

func twoPair(h Hand) bool {
	return false
}

func pair(h Hand) bool {
	return false
}

func main() {
	lines, err := readLines("./p054_poker.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var p1Hand Hand
	var p2Hand Hand

	for i := 1; i <= 1000; i++ {

		p1Hand, p2Hand = getHands(lines[i-1])
		// fmt.Println("b4", p1Hand)
		// fmt.Println("b4", p2Hand)

		if fullHouse(p1Hand) {
			fmt.Println(p1Hand)
		}

		if fullHouse(p2Hand) {
			fmt.Println(p2Hand)
		}

		// do hand eval here

		p1Hand = Hand{}
		p2Hand = Hand{}
	}

}
