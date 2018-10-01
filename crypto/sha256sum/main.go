package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no files")
	}

	filename := os.Args[1]

	sum, err := fileHashSHA256(filename)
	checkError(err)

	fmt.Printf("%x %s\n", sum, filename)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fileHashSHA256(filename string) ([sha256.Size]byte, error) {
	var sum [sha256.Size]byte
	bs, err := fileHash(filename, sha256.New())
	if err != nil {
		return sum, err
	}
	if len(bs) != sha256.Size {
		return sum, errors.New("invalid checksum size")
	}
	copy(sum[:], bs)
	return sum, nil
}

func fileHash(filename string, h hash.Hash) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	h.Reset()
	_, err = io.Copy(h, f)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
