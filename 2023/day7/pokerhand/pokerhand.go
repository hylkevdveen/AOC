package pokerhand

import (
	"../../../aoc"
	"sort"
	"strconv"
	"strings"
)

type PokerHand struct {
	Hand     []int
	HandType int
	Bet      int
}

var playingCards = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var playingCardsJoker = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func ReadHands(joker bool) []*PokerHand {
	var pokerHands []*PokerHand
	playingCardValues := makePlayingCardValues(joker)
	for _, line := range aoc.InputLines() {
		hand := strings.Split(line, " ")
		cards := make([]int, 0, 5)
		for _, card := range hand[0] {
			cards = append(cards, playingCardValues[card])
		}
		handtp := handType(cards, joker)
		value, _ := strconv.Atoi(hand[1])
		pokerHand := PokerHand{Hand: cards, HandType: handtp, Bet: value}
		pokerHands = append(pokerHands, &pokerHand)
	}
	return pokerHands
}

func makePlayingCardValues(joker bool) map[rune]int {
	cardValues := make(map[rune]int, 13)
	playingCardOrder := make([]string, 0, 13)
	if joker {
		playingCardOrder = playingCardsJoker
	} else {
		playingCardOrder = playingCards
	}
	for i, card := range playingCardOrder {
		cardValues[rune(card[0])] = i
	}
	return cardValues
}

func handType(cards []int, joker bool) int {
	cardCounts := make(map[int]int)
	for _, card := range cards {
		cardCounts[card] = cardCounts[card] + 1
	}
	jokers := 0
	if joker {
		jokers = cardCounts[0]
		delete(cardCounts, 0)
	}
	switch len(cardCounts) {
	case 0, 1:
		return FiveOfAKind
	case 2:
		for card := range cardCounts {
			if cardCounts[card] == 1 || cardCounts[card]+jokers == 4 {
				return FourOfAKind
			}
			return FullHouse
		}
	case 3:
		if jokers != 0 {
			return ThreeOfAKind
		}
		for card := range cardCounts {
			if cardCounts[card] == 2 {
				return TwoPair
			}
			if cardCounts[card] == 3 {
				return ThreeOfAKind
			}
		}
	case 4:
		return OnePair
	}
	return HighCard
}

func SortHands(pokerHands []*PokerHand) {
	sort.Slice(pokerHands, func(i, j int) bool {
		if pokerHands[i].HandType != pokerHands[j].HandType {
			return pokerHands[i].HandType < pokerHands[j].HandType
		}
		for k := 0; k < 5; k++ {
			if pokerHands[i].Hand[k] != pokerHands[j].Hand[k] {
				return pokerHands[i].Hand[k] < pokerHands[j].Hand[k]
			}
		}
		return false
	})
}
