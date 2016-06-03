// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package coci

// Ente64 represents a rational number where the numerator and denominator are
// int32 types.
type Ente64 struct {
	n, d int32
}

// Num returns the numerator of z.
func (z *Ente64) Num() int32 {
	return z.n
}

// Denom returns the denominator of z.
func (z *Ente64) Denom() int32 {
	return z.d
}
