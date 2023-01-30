package protein

import (
	"errors"
)

var rToC = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var (
	ErrStop              = errors.New("stop error")
	ErrInvalidBase error = errors.New("invalid base error")
)

func FromRNA(rna string) (c []string, err error) {
	for len(rna) > 0 {
		codon := rna[:3]
		s, e := FromCodon(codon)
		switch e {
		case nil:
			c = append(c, s)
		case ErrStop:
			return
		case ErrInvalidBase:
			err = e
			return
		}
		if len(rna) >= 3 {
			rna = rna[3:]
		}
	}
	return
}

func FromCodon(codon string) (string, error) {
	if r, ok := rToC[codon]; ok {
		if r == "STOP" {
			return "", ErrStop
		}
		return r, nil
	}
	return "", ErrInvalidBase
}
