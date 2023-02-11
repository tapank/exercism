package poker

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards    []Card
	suites   map[Suite]int
	ranks    map[int]int
	handrank int
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
		if rank > int(h.handrank) {
			h.handrank = rank
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

	// fmt.Println("************hurray! Here are the hands")
	// fmt.Println(allhands)
	// score all hands
	return nil, nil
}

func checkStraightFlush(hands map[*Hand]string) ([]string, bool) {
	best := []string{}
	ok := false

	// pick out straight flush hands and find the highest score
	bestHands := []*Hand{}
	bestScore := 0
	for hand := range hands {
		if hand.isStraightFlush() {
			bestHands = append(bestHands, hand)
			if hand.handrank > bestScore {
				bestScore = hand.handrank
			}
		}
	}

	// pick out the best scoring hand; there may be more than one
	for _, hand := range bestHands {
		if hand.handrank == bestScore {
			best = append(best, hands[hand])
			ok = true
		}
	}
	return best, ok
}

func parse(hands []string) (map[*Hand]string, error) {
	allhands := make(map[*Hand]string, len(hands))
	allcards := make(map[Card]struct{})
	for _, hand := range hands {
		cards := make([]Card, 0, 5)
		for _, cardname := range strings.Split(hand, " ") {
			if card, err := NewCard(cardname); err != nil {
				return nil, err
			} else {
				if _, ok := allcards[card]; ok {
					return nil, errors.New("duplicate card")
				} else {
					allcards[card] = struct{}{}
				}
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
