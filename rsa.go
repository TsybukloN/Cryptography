package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// ----------------- Helper Functions -----------------

// Generate a random prime number of n bits
func generatePrime(bits int) *big.Int {
	prime, _ := rand.Prime(rand.Reader, bits)
	return prime
}

// Compute modular inverse using Extended Euclidean Algorithm
func modInverse(e, phi *big.Int) *big.Int {
	return new(big.Int).ModInverse(e, phi)
}

// ----------------- RSA -----------------

// Encrypt and decrypt a message (int) using RSA
func rsa_encrypt(message *big.Int) (*big.Int, *big.Int) {
	bits := 512

	// Generate two random prime numbers
	p := generatePrime(bits)
	q := generatePrime(bits)

	// Modul n = p * q (Part of public key)
	n := new(big.Int).Mul(p, q)

	// Euler's function phi = (p-1) * (q-1)
	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	e := big.NewInt(65537) // Common public exponent
	// Private key d = e^(-1) mod phi
	d := modInverse(e, phi)

	fmt.Println("Modulus: ", n)
	fmt.Println("Public Key: ", e)
	fmt.Println("Private Key: ", d)

	// C = M^e mod n
	cipher := new(big.Int).Exp(message, e, n)
	// M = C^d mod n
	decrypted := new(big.Int).Exp(cipher, d, n)

	fmt.Println("Original Message:", message)
	fmt.Println("Encrypted Message:", cipher)
	fmt.Println("Decrypted Message:", decrypted)

	return d, e
}
