package main

import (
	"math/big"
	mathrand "math/rand"
	"time"
)

// --- Prime number functions ---

func generateMersennePrime(bits int) *big.Int {
	mathrand.Seed(time.Now().UnixNano())
	two := big.NewInt(2)

	n := big.NewInt(int64(mathrand.Intn(100) + 5))
	for !n.ProbablyPrime(5) {
		n.Add(n, big.NewInt(1))
	}

	for {
		mersenne := new(big.Int).Exp(two, n, nil)
		mersenne.Sub(mersenne, big.NewInt(1))

		if mersenne.ProbablyPrime(10) && mersenne.BitLen() >= bits {
			return mersenne
		}

		n.Add(n, big.NewInt(int64(mathrand.Intn(5)+1)))
		for !n.ProbablyPrime(5) {
			n.Add(n, big.NewInt(1))
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// --- Factorization functions ---

// Factorize function to find the prime factors of a given number
func factorize(n big.Int) []*big.Int {
	var factors []*big.Int
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	// only big.Int class by modifying copy and original value changed, so we need to create a copy
	nCopy := new(big.Int).Set(&n)

	// Divide by 2 until n is odd
	for new(big.Int).Mod(nCopy, two).Cmp(zero) == 0 {
		factors = append(factors, new(big.Int).Set(two))
		nCopy.Div(nCopy, two)
	}

	// Divide by odd numbers starting from 3
	divisor := big.NewInt(3)
	for nCopy.Cmp(one) == 1 {
		for new(big.Int).Mod(nCopy, divisor).Cmp(zero) == 0 {
			factors = append(factors, new(big.Int).Set(divisor))
			nCopy.Div(nCopy, divisor)
		}
		divisor.Add(divisor, two) // Next odd number
	}

	return factors
}

// Pollard's Rho algorithm to find the prime factors of a given number
func pollardRho(n big.Int) *big.Int {
	one := big.NewInt(1)
	x := big.NewInt(2)
	y := big.NewInt(2)
	d := big.NewInt(1)

	// only big.Int class by modifying copy and original value changed, so we need to create a copy
	nCopy := new(big.Int).Set(&n)

	f := func(x *big.Int) *big.Int {
		res := new(big.Int).Mul(x, x)
		res.Add(res, one)
		return res.Mod(res, nCopy)
	}

	for d.Cmp(one) == 0 {
		x = f(x)
		y = f(f(y))
		d.GCD(nil, nil, new(big.Int).Abs(new(big.Int).Sub(x, y)), nCopy)
	}

	if d.Cmp(nCopy) == 0 {
		return nil
	}
	return d
}
