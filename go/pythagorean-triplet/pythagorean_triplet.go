package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) (t []Triplet) {
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			for k := j; k <= max; k++ {
				if i*i+j*j == k*k {
					t = append(t, Triplet{i, j, k})
				}
				if i*i+j*j < k*k {
					break
				}
			}
		}
	}
	return
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) (t []Triplet) {
	for i := 1; i <= p/2; i++ {
		for j := i; j <= p/2; j++ {
			k := p - (i + j)
			if k > 0 {
				if i*i+j*j == k*k {
					t = append(t, Triplet{i, j, k})
				}
			} else {
				break
			}
		}
	}
	return
}
