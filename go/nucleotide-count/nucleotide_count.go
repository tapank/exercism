package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nuc := range d {
		if _, ok := h[nuc]; ok {
			h[nuc]++
		} else {
			return nil, fmt.Errorf("invalid nucleotide:%c", nuc)
		}
	}
	return h, nil
}
