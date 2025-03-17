package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mathrand "math/rand"
	"time"
)

type RSA struct {
	n, e, d *big.Int
}

// ----------------- Helper Functions -----------------

// Generate a random prime number of n bits
func generatePrime(bits int) *big.Int {
	prime, _ := rand.Prime(rand.Reader, bits)
	return prime
}

// Generate a random Mersenne prime number of n bits
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

// ----------------- RSA -----------------

func (rsa *RSA) createKeys() {
	bits := 512

	// Generate two random prime numbers
	p, _ := rand.Prime(rand.Reader, bits)
	q, _ := rand.Prime(rand.Reader, bits)

	for p.Cmp(q) == 0 {
		fmt.Println("p and q are equal, generating new q")
		q, _ = rand.Prime(rand.Reader, bits)
	}

	// Modul n = p * q (Part of public key)
	rsa.n = new(big.Int).Mul(p, q)

	// Euler's function phi = (p-1) * (q-1)
	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	rsa.e = big.NewInt(65537) // Common public exponent
	// Private key d = e^(-1) mod phi
	rsa.d = new(big.Int).ModInverse(rsa.e, phi)

	fmt.Println("Modulus (n): ", rsa.n)
	fmt.Println("Public Key(e) without (n): ", rsa.e)
	fmt.Println("Private Key(d) without (n): ", rsa.d)
}

func (rsa *RSA) encrypt(message *big.Int) *big.Int {
	if rsa.n.Cmp(message) == -1 {
		fmt.Println("Message is too long")
		return nil
	}

	cipher := new(big.Int).Exp(message, rsa.e, rsa.n)
	fmt.Println("Encrypted Message:", cipher)
	return cipher
}

func (rsa *RSA) decrypt(cipher *big.Int) *big.Int {
	decrypted := new(big.Int).Exp(cipher, rsa.d, rsa.n)
	fmt.Println("Decrypted Message:", decrypted)
	return decrypted
}

func (rsa *RSA) creck(cipher *big.Int) *big.Int {
	// Only can use rsa.e and rsa.n

}

// Factorize function to find the prime factors of a given number
func factorize(n *big.Int) []*big.Int {
	var factors []*big.Int
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	// Divide by 2 until n is odd
	for new(big.Int).Mod(n, two).Cmp(zero) == 0 {
		factors = append(factors, new(big.Int).Set(two))
		n.Div(n, two)
	}

	// Check for odd factors from 3 onwards
	for i := big.NewInt(3); new(big.Int).Mul(i, i).Cmp(n) <= 0; i.Add(i, two) {
		for new(big.Int).Mod(n, i).Cmp(zero) == 0 {
			factors = append(factors, new(big.Int).Set(i))
			n.Div(n, i)
		}
	}

	// If n is a prime number greater than 2
	if n.Cmp(two) > 0 {
		factors = append(factors, n)
	}

	return factors
}

func rsa_demo() {
	message := big.NewInt(1234567890)

	my_rsa := RSA{}
	my_rsa.createKeys()

	fmt.Println("Original Message:", message)
	c := my_rsa.encrypt(message)
	my_rsa.decrypt(c)
}
