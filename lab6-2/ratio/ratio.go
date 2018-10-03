package ratio

import (
	"fmt"
)

// Ratio is a rational number
type Ratio struct {
	num int
	den int
}

// New creates a Ratio
func New(num, den int) *Ratio {
	return &Ratio{num, den}
}

func (r *Ratio) String() string {
	return fmt.Sprintf("%d / %d", r.num, r.den)
}

// SumInst is the sum of two rational numbers
func (r *Ratio) SumInst(x *Ratio, y *Ratio) *Ratio {
	return x.Sum(y)
}

// Sum is the sum of two rational numbers
func (r *Ratio) Sum(other *Ratio) *Ratio {
	temp := new(Ratio)
	temp.num = r.num*other.den + other.num*r.den
	temp.den = r.den * other.den
	return temp
}

// StaticSum is the sum of two rational numbers
func StaticSum(x *Ratio, y *Ratio) *Ratio {
	temp := new(Ratio)
	return temp.SumInst(x, y)
}
