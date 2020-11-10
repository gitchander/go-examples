package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"log"
)

func main() {
	exampleCTR()
}

func exampleCTR() {
	key := make([]byte, 24)
	_, err := rand.Read(key)
	checkError(err)

	b, err := aes.NewCipher(key)
	checkError(err)

	fmt.Println("blockSize:", b.BlockSize())

	iv := make([]byte, b.BlockSize())
	_, err = rand.Read(iv)
	checkError(err)

	plainText := []byte("Hello, World!")
	cipherText := make([]byte, len(plainText))

	// Encryption
	{
		stream := cipher.NewCTR(b, iv)

		stream.XORKeyStream(cipherText, plainText)
		fmt.Printf("cipherText: %x\n", cipherText)
	}

	addNoice := false
	if addNoice {
		_, err = rand.Read(cipherText[3:6])
		checkError(err)
	}

	// Decryption
	{
		stream := cipher.NewCTR(b, iv)

		plainTextDec := make([]byte, len(cipherText))
		stream.XORKeyStream(plainTextDec, cipherText)

		fmt.Printf("plainTextDec: %s\n", plainTextDec)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
