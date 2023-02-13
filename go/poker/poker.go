package poker

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards          []Card
	suites         map[Suite]int
	ranks          map[int]int
	primaryscore   int
	secondaryscore int
}

type Card struct {
	suite Suite
	rank  int
}

type Suite int

const (
	CLUBS Suite = iota
	SPADES
	HEARTS
	DIAMOND
)

type HandRank int

const (
	_ HandRank = iota
	STRAIGHT_FLUSH
	FOUR_OF_A_KIND
	FULL_HOUSE
	FLUSH
	STRAIGHT
	THREE_OF_A_KIND
	TWO_PAIR
	ONE_PAIR
	HIGH_CARD
)

func NewHand(cards []Card) (*Hand, error) {
	if len(cards) != 5 {
		return nil, errors.New("wrong number of cards in hand")
	}

	h := &Hand{cards: cards}
	h.suites = make(map[Suite]int)
	h.ranks = make(map[int]int)
	for _, card := range cards {
		h.suites[card.suite]++
		h.ranks[card.rank]++
	}
	return h, nil
}

func NewCard(s string) (Card, error) {
	card := Card{}
	chars := []rune(s)
	if len(chars) < 2 {
		return card, errors.New("unparsable card:" + s)
	}
	if r := string(chars[:len(chars)-1]); strings.ContainsAny(r, "AJQK") {
		if len(r) > 1 {
			return card, errors.New("unparsable rank:" + s)
		}
		switch r[0] {
		case 'A':
			card.rank = 1
		case 'J':
			card.rank = 11
		case 'Q':
			card.rank = 12
		case 'K':
			card.rank = 13
		default:
			return card, errors.New("unparsable rank:" + s)
		}
	} else if rank, err := strconv.Atoi(r); err != nil {
		return card, errors.New("unparsable rank:" + s)
	} else if rank < 2 || rank > 10 {
		return card, errors.New("invalid rank:" + s)
	} else {
		card.rank = rank
	}

	switch chars[len(chars)-1] {
	case '♤':
		card.suite = SPADES
	case '♡':
		card.suite = HEARTS
	case '♧':
		card.suite = CLUBS
	case '♢':
		card.suite = DIAMOND
	default:
		return card, errors.New("invalid suite:" + s)
	}
	return card, nil
}

func (h *Hand) isStraightFlush() bool {
	if len(h.suites) != 1 || len(h.ranks) != 5 {
		return false
	}
	sort.Slice(h.cards, func(i, j int) bool {
		return h.cards[i].rank < h.cards[j].rank
	})

	for i := 0; i < len(h.cards)-2; i++ {
		if h.cards[i].rank-h.cards[i+1].rank != -1 {
			return false
		}
	}
	for rank := range h.ranks {
		if rank > int(h.primaryscore) {
			h.primaryscore = rank
		}
	}
	return true
}

func BestHand(hands []string) ([]string, error) {
	allhands, err := parse(hands)
	if err != nil {
		return nil, err
	}
	// cascase down the hand types in decending order
	// and return the first set found
	if best, ok := checkStraightFlush(allhands); ok {
		return best, nil
	}
	if best, ok := checkFourOfAKind(allhands); ok {
		return best, nil
	}
	if best, ok := checkFullHouse(allhands); ok {
		return best, nil
	}

	// fmt.Println("************hurray! Here are the hands")
	// fmt.Println(allhands)
	// score all hands
	return nil, nil
}

func checkStraightFlush(hands map[*Hand]string) ([]string, bool) {
	// pick out straight flush hands and find the highest score
	besthands := []*Hand{}
	bestscore := 0
	for hand := range hands {
		if hand.isStraightFlush() {
			besthands = append(besthands, hand)
			if hand.primaryscore > bestscore {
				bestscore = hand.primaryscore
			}
		}
	}
	return pickTopHands(besthands, bestscore, hands)
}

func checkFourOfAKind(hands map[*Hand]string) ([]string, bool) {
	besthands := []*Hand{}
	pbest := 0
	sbest := 0

	// pick out four of a kind hands and find the highest score
	for hand := range hands {
		for rank, count := range hand.ranks {
			if count == 4 {
				hand.primaryscore = rank
				besthands = append(besthands, hand)
				if rank > pbest {
					pbest = rank
				}
			} else if len(hands) == 2 && count == 1 {
				hand.secondaryscore = rank
				if rank > sbest {
					sbest = rank
				}
			}
		}
	}
	return pickTopHands(besthands, pbest, hands)
}

func checkFullHouse(hands map[*Hand]string) ([]string, bool) {
	besthands := []*Hand{}
	pbest := 0
	sbest := 0

	// pick out four of a kind hands and find the highest score
	for hand := range hands {
		for rank, count := range hand.ranks {
			if len(hands) == 2 && count == 3 {
				hand.primaryscore = rank
				besthands = append(besthands, hand)
				if rank > pbest {
					pbest = rank
				}
			} else if len(hands) == 2 && count == 2 {
				hand.secondaryscore = rank
				if rank > sbest {
					sbest = rank
				}
			}
		}
	}
	return pickTopHands(besthands, pbest, hands)
}

func checkFlush(hands map[*Hand]string) ([]string, bool) {
	besthands := []*Hand{}
	pbest := 0
	sbest := 0

	// pick out four of a kind hands and find the highest score
	for hand := range hands {
		if len(hand.suites) > 1 {
			break
		}
		for rank, count := range hand.ranks {
			if len(hands) == 2 && count == 3 {
				hand.primaryscore = rank
				besthands = append(besthands, hand)
				if rank > pbest {
					pbest = rank
				}
			} else if len(hands) == 2 && count == 2 {
				hand.secondaryscore = rank
				if rank > sbest {
					sbest = rank
				}
			}
		}
	}
	return pickTopHands(besthands, pbest, hands)
}

// pick out the best scoring hand; there may be more than one
func pickTopHands(besthands []*Hand, pbest int, hands map[*Hand]string) ([]string, bool) {
	if pbest == 0 {
		return nil, false
	}
	bestbyprimary := []*Hand{}
	var sbest int
	for _, hand := range besthands {
		if hand.primaryscore == pbest {
			bestbyprimary = append(bestbyprimary, hand)
			if hand.secondaryscore > sbest {
				sbest = hand.secondaryscore
			}
		}
	}
	besthands = bestbyprimary
	bestbysecondary := []*Hand{}
	if len(bestbyprimary) > 1 {
		for _, hand := range bestbyprimary {
			if hand.secondaryscore == sbest {
				bestbysecondary = append(bestbysecondary, hand)
			}
		}
		besthands = bestbysecondary
	}

	best := []string{}
	for _, hand := range besthands {
		if hand.primaryscore == pbest && hand.secondaryscore == sbest {
			best = append(best, hands[hand])
		}
	}
	return best, true
}

func parse(hands []string) (map[*Hand]string, error) {
	allhands := make(map[*Hand]string, len(hands))
	for _, hand := range hands {
		cards := make([]Card, 0, 5)
		for _, cardname := range strings.Split(hand, " ") {
			if card, err := NewCard(cardname); err != nil {
				return nil, err
			} else {
				cards = append(cards, card)
			}
		}
		if h, err := NewHand(cards); err != nil {
			return nil, err
		} else {
			allhands[h] = hand
		}
	}
	return allhands, nil
}
