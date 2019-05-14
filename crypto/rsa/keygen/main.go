package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	simple()
	//testSaveKeysToPEM()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func simple() {
	alisePrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	checkError(err)

	alisePublicKey := &alisePrivateKey.PublicKey

	secretMessage := []byte("Hello, World!")
	label := []byte{}
	hash := sha256.New()

	data, err := rsa.EncryptOAEP(hash, rand.Reader, alisePublicKey, secretMessage, label)
	checkError(err)

	text, err := rsa.DecryptOAEP(hash, rand.Reader, alisePrivateKey, data, label)
	checkError(err)

	fmt.Printf("decrypt message: %s\n", text)
}

func generateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func testSaveKeysToPEM() {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	checkError(err)

	privateKeyPEM, err := EncodeRSAPrivateKey(privateKey)
	checkError(err)

	filePrefix := "name_"
	privateFilename := filePrefix + "private_key.pem"
	err = ioutil.WriteFile(privateFilename, privateKeyPEM, 0666)
	checkError(err)

	publicKeyPEM, err := EncodeRSAPublicKey(&privateKey.PublicKey)
	checkError(err)
	err = ioutil.WriteFile(filePrefix+"public_key.pem", publicKeyPEM, 0666)
	checkError(err)

	data, err := ioutil.ReadFile(privateFilename)
	checkError(err)

	b, _ := pem.Decode(data)

	prk, err := x509.ParsePKCS1PrivateKey(b.Bytes)
	checkError(err)

	fmt.Println(privateKey.PublicKey)
	fmt.Println(prk.PublicKey)
}

func EncodeRSAPrivateKey(privateKey *rsa.PrivateKey) ([]byte, error) {

	privateKeyASN1 := x509.MarshalPKCS1PrivateKey(privateKey)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyASN1,
	})

	return privateKeyPEM, nil
}

func DecodeRSAPrivateKey(data []byte) (*rsa.PrivateKey, error) {
	p, _ := pem.Decode(data)
	if p == nil {
		return nil, errors.New("invalid RSA private key")
	}
	if p.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid RSA private key")
	}
	return x509.ParsePKCS1PrivateKey(p.Bytes)
}

func EncodeRSAPublicKey(publicKey *rsa.PublicKey) ([]byte, error) {

	//publicKeyBIN := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyBIN, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBIN,
	})

	return publicKeyPEM, nil
}

var ErrInvalidRSAPublicKey = errors.New("invalid RSA public key")

func DecodeRSAPublicKey(data []byte) (*rsa.PublicKey, error) {

	p, _ := pem.Decode(data)
	if p == nil {
		return nil, ErrInvalidRSAPublicKey
	}

	if p.Type != "RSA PUBLIC KEY" {
		return nil, ErrInvalidRSAPublicKey
	}

	pub, err := x509.ParsePKIXPublicKey(p.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, ErrInvalidRSAPublicKey
	}

	return publicKey, nil
}
