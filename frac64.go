// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package frac

// Frac64 represents a rational number where the numerator and denominator are
// int32 types.
type Frac64 struct {
	n, d int32
}

// Num returns the numerator of z.
func (z *Frac64) Num() int32 {
	return z.n
}

// Denom returns the denominator of z.
func (z *Frac64) Denom() int32 {
	return z.d
}
