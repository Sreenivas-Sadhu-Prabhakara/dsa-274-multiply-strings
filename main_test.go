package main

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestMultiply(t *testing.T) {
	cases := []struct{ a, b, want string }{
		{"2", "3", "6"},
		{"123", "456", "56088"},
		{"0", "52", "0"},
		{"99", "0", "0"},
		{"999", "999", "998001"},
		{"1", "1", "1"},
		{"123456789", "987654321", "121932631112635269"},
	}
	for _, c := range cases {
		if got := Multiply(c.a, c.b); got != c.want {
			t.Errorf("Multiply(%q,%q) = %q, want %q", c.a, c.b, got, c.want)
		}
	}
}

// Property: agree with math/big over random large numbers.
func TestAgainstBig(t *testing.T) {
	rng := rand.New(rand.NewSource(5))
	for i := 0; i < 100; i++ {
		a := new(big.Int).Rand(rng, new(big.Int).Exp(big.NewInt(10), big.NewInt(25), nil))
		b := new(big.Int).Rand(rng, new(big.Int).Exp(big.NewInt(10), big.NewInt(25), nil))
		want := new(big.Int).Mul(a, b).String()
		if got := Multiply(a.String(), b.String()); got != want {
			t.Fatalf("Multiply(%s,%s) = %s, want %s", a, b, got, want)
		}
	}
}
