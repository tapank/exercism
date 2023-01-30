package proverb

import (
	"fmt"
)

// Proverb returns a proverb based on the provided strings
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return nil
	}
	p := []string{}
	for i := 0; i < len(rhyme)-1; i++ {
		p = append(p, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	return append(p, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
}
