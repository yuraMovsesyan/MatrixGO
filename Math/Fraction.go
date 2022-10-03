package Math

import (
	"fmt"
	"strconv"
)

type Fraction struct {
	Numerator   int
	Denominator int
}

func NewFrac(numerator int, denominator int) (frac Fraction) {
	frac.Numerator = numerator
	frac.Denominator = denominator

	return
}

func NewFracNum(numerator int) Fraction {
	return NewFrac(numerator, 1)
}

func FracAddFrac(a, b Fraction) (c Fraction) {
	c.Denominator = NOK(a.Denominator, b.Denominator)

	c.Numerator = a.Numerator*(c.Denominator/a.Denominator) + b.Numerator*(c.Denominator/b.Denominator)

	return
}

func FracMulNum(a Fraction, b int) (c Fraction) {
	c.Numerator = a.Numerator * b
	c.Denominator = a.Denominator

	return
}

func FracMulFrac(a, b Fraction) (c Fraction) {
	c.Numerator = a.Numerator * b.Numerator
	c.Denominator = a.Denominator * b.Denominator

	return
}

func FracDivNum(a Fraction, b int) (c Fraction) {
	c.Numerator = a.Numerator
	c.Denominator = a.Denominator * b

	return
}

func FracDivFrac(a, b Fraction) (c Fraction) {
	c.Numerator = b.Denominator
	c.Denominator = b.Numerator

	c = FracMulFrac(a, c)

	return
}

func (frac Fraction) FracCut() Fraction {
	nod := NOD(frac.Numerator, frac.Denominator)

	if nod == 1 {
		return frac
	}

	frac.Numerator /= nod
	frac.Denominator /= nod

	return frac
}

func (frac Fraction) Result() float64 {
	return float64(frac.Numerator) / float64(frac.Denominator)
}

func (frac Fraction) ToString() (line string) {

	line += strconv.Itoa(frac.Numerator)

	if frac.Denominator != 1 {
		line += "\\" + strconv.Itoa(frac.Denominator)
	}

	return
}

func (frac Fraction) Print() Fraction {
	fmt.Print(frac.ToString())

	return frac
}

func (frac Fraction) Println() Fraction {
	fmt.Println(frac.ToString())

	return frac
}
