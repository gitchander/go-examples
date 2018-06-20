package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file, err := ioutil.TempFile("", "test_tmp_")
	checkError(err)
	defer func() {
		checkError(file.Close())
		checkError(os.Remove(file.Name()))
	}()

	fmt.Println("create temp file:", file.Name())

	body := []byte("temp file body")
	_, err = file.Write(body)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
