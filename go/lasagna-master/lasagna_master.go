package lasagna

func PreparationTime(layers []string, ltime int) int {
	if ltime == 0 {
		ltime = 2
	}
	return len(layers) * ltime
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}
	}
	return
}

func AddSecretIngredient(flist, mlist []string) {
	if flen, mlen := len(flist), len(mlist); flen > 0 && mlen > 0 {
		mlist[mlen-1] = flist[flen-1]
	}
}

func ScaleRecipe(f []float64, portions int) []float64 {
	quantities := []float64{}
	for _, q := range f {
		if q > 0 {
			q *= float64(portions) / 2
		}
		quantities = append(quantities, q)
	}
	return quantities
}
