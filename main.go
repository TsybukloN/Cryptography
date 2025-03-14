package main

import (
	"math/big"
)

// Modulo function
func mod(a, b int) int {
	return a % b
}

// Is prime function
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

func main() {
	message := big.NewInt(12345)
	rsaEncryptDecrypt(message)
}
