package strand

var d2r = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	d := []rune(dna)
	r := make([]rune, len(d))
	for i := 0; i < len(d); i++ {
		r[i] = d2r[d[i]]
	}
	return string(r)
}
