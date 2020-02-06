package main

import (
	"fmt"
)

type ErrNegSqrt float64

func (e ErrNegSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt %v negative number", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x<0 {
		return 0, ErrNegSqrt(x)
	}
	z := 1.0
	for i := 0; i<10 ; i++ {
		z -= (z*z - x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
