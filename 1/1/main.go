package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	var input string

	fmt.Scanln(&input)

	rawBytes, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(rawBytes))
}
