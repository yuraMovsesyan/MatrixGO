package STD

func MinFracArray(frac []Fraction) (id int, minFrac Fraction) {

	id, minFrac = 0, frac[0]

	for i := 1; i < len(frac); i++ {
		if frac[i].Result() < minFrac.Result() {
			id, minFrac = i, frac[i]
		}
	}

	return id, minFrac
}

func MaxFracArray(frac []Fraction) (id int, maxFrac Fraction) {

	id, maxFrac = 0, frac[0]

	for i := 1; i < len(frac); i++ {
		if frac[i].Result() > maxFrac.Result() {
			id, maxFrac = i, frac[i]
		}
	}

	return id, maxFrac
}
