package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"log"
)

const textLoremIpsum = `Lorem ipsum dolor sit amet,
consectetur adipiscing elit, sed do eiusmod tempor
incididunt ut labore et dolore magna aliqua.
Sit amet consectetur adipiscing elit pellentesque
habitant morbi. Sed vulputate mi sit amet.
Felis bibendum ut tristique et. Egestas maecenas
pharetra convallis posuere morbi. Netus et malesuada
fames ac turpis egestas sed tempus. Donec ultrices
tincidunt arcu non sodales neque sodales. Turpis
egestas pretium aenean pharetra magna. Nisi est
sit amet facilisis magna etiam tempor orci eu.
Elit pellentesque habitant morbi tristique senectus.
Mi quis hendrerit dolor magna eget est lorem ipsum.`

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func makeRandKey(n int) ([]byte, error) {
	key := make([]byte, n)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func main() {

	key, err := makeRandKey(24)
	checkError(err)
	fmt.Printf("Key: %X\n", key)

	b, err := aes.NewCipher(key) // key len 16, 24, 32 bytes
	checkError(err)

	fmt.Println("BlockSize:", b.BlockSize())

	iv, err := makeRandKey(b.BlockSize())
	checkError(err)
	fmt.Printf("IV: %X\n", iv)

	plaintext := []byte(textLoremIpsum)

	// Encryption
	ciphertext := make([]byte, len(plaintext))
	e := cipher.NewCFBEncrypter(b, iv)
	e.XORKeyStream(ciphertext, plaintext)

	// Decryption
	plaintextDec := make([]byte, len(ciphertext))
	d := cipher.NewCFBDecrypter(b, iv)
	d.XORKeyStream(plaintextDec, ciphertext)

	fmt.Println("Decrypt result:", string(plaintextDec))
}
