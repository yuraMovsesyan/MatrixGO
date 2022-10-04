package STD

func NOD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func NOK(a, b int) int {
	return (a / NOD(a, b)) * b
}
