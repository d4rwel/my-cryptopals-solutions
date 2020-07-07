package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	var hexInput1 string
	var hexInput2 string

	fmt.Scanln(&hexInput1)
	fmt.Scanln(&hexInput2)

	byteInput1, err := hex.DecodeString(hexInput1)
	if err != nil {
		log.Fatal(err)
	}
	byteInput2, err := hex.DecodeString(hexInput2)
	if err != nil {
		log.Fatal(err)
	}

	if len(byteInput1) != len(byteInput2) {
		log.Fatal("input strings have different length")
	}

	byteOutput := make([]byte, len(byteInput1))
	for i := 0; i < len(byteInput1); i++ {
		byteOutput[i] = byteInput1[i] ^ byteInput2[i]
	}

	fmt.Printf("%s\n", hex.EncodeToString(byteOutput))
}
