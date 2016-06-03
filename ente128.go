// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package coci

// Ente128 represents a rational number where the numerator and denominator are
// int64 types.
type Ente128 struct {
	n, d int64
}

// Num returns the numerator of z.
func (z *Ente128) Num() int64 {
	return z.n
}

// Denom returns the denominator of z.
func (z *Ente128) Denom() int64 {
	return z.d
}
