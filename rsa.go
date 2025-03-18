package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type RSA struct {
	n, e, d *big.Int
}

func (rsa *RSA) createKeys() {
	// Number of bits for the prime numbers (p, q)
	// Not secure : 8, 16, 32
	// Secured : 512, 1024, 2048
	bits := 32

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

	fmt.Println(phi)

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

func (rsa *RSA) hack(cipher *big.Int) *big.Int {
	// Only can use rsa.e and rsa.n
	fmt.Println("-----Hacking the message-----")

	// factors := factorize(*rsa.n)
	//p, q := factors[0], factors[1]

	p := pollardRho(*rsa.n)
	q := new(big.Int).Div(rsa.n, p)

	fmt.Println("Hacked factors of the (p, q):", p, q)

	hackedPhi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	hackedD := new(big.Int).ModInverse(rsa.e, hackedPhi)

	fmt.Println("Hacked Private Key(d):", hackedD)
	decrypted := new(big.Int).Exp(cipher, hackedD, rsa.n)

	fmt.Println("Hacked Original Message:", decrypted)
	return decrypted
}

// --- DEMOS ---

func rsaDemo() {
	message := big.NewInt(12345)

	myRsa := RSA{}
	myRsa.createKeys()

	fmt.Println("Original Message:", message)
	c := myRsa.encrypt(message)
	myRsa.decrypt(c)
}

func rsaDemoHack() {
	message := big.NewInt(12345)

	myRsa := RSA{}
	myRsa.createKeys()

	fmt.Println("Original Message:", message)
	c := myRsa.encrypt(message)
	myRsa.hack(c)
}
