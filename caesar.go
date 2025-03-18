package main

import "fmt"

const letterCount int = 26
const a int = int('a')
const A int = int('A')

type CaesarCipher struct {
	key int
}

func NewCaesarCipher(key int) *CaesarCipher {
	key = key % 26
	return &CaesarCipher{key: key}
}

func (caesar *CaesarCipher) encrypt(message string) string {
	encrypted := ""
	for _, char := range message {
		indexChar := int(char)
		if indexChar >= a && indexChar < a+letterCount {
			encrypted += string(rune((indexChar+caesar.key-a)%letterCount + a))
		} else if indexChar >= A && indexChar < A+letterCount {
			encrypted += string(rune((indexChar+caesar.key-A)%letterCount + A))
		} else {
			encrypted += string(char)
		}
	}

	fmt.Println("Encrypted message:", encrypted)
	return encrypted
}

func (caesar *CaesarCipher) decrypt(encrypted string) string {
	decrypted := ""
	for _, char := range encrypted {
		indexChar := int(char)
		if indexChar >= a && indexChar < a+letterCount {
			decrypted += string(rune((indexChar-caesar.key-a+letterCount)%letterCount + a))
		} else if indexChar >= A && indexChar < A+letterCount {
			decrypted += string(rune((indexChar-caesar.key-A+letterCount)%letterCount + A))
		} else {
			decrypted += string(char)
		}
	}

	fmt.Println("Decrypted message:", decrypted)
	return decrypted
}

func (caesar *CaesarCipher) hack(encrypted string) {
	for i := 1; i < letterCount; i++ {
		decrypted := ""
		for _, char := range encrypted {
			indexChar := int(char)
			if indexChar >= a && indexChar < a+letterCount {
				decrypted += string(rune((indexChar-i-a+letterCount)%letterCount + a))
			} else if indexChar >= A && indexChar < A+letterCount {
				decrypted += string(rune((indexChar-i-A+letterCount)%letterCount + A))
			} else {
				decrypted += string(char)
			}
		}
		fmt.Println("Key:", i, "Decrypted message:", decrypted)
	}
}

// --- DEMOS ---

func caesarDemo() {
	caesar := NewCaesarCipher(20)
	message := "Hello, World! 123"
	fmt.Println("Original message:", message)

	encrypted := caesar.encrypt(message)
	caesar.decrypt(encrypted)

}

func caesarDemoHack() {
	caesar := NewCaesarCipher(16)
	message := "Hello, World! 123"
	fmt.Println("Original message:", message)

	encrypted := caesar.encrypt(message)
	caesar.hack(encrypted)
}
