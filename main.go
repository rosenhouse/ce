package main

import (
	"crypto/aes"
	"fmt"
	"os"
)

const (
	blockSize = 16 // bytes
	keySize   = 32 // bytes
)

func Fatalf(format string, a ...interface{}) {
	os.Stderr.WriteString(fmt.Sprintf(format, a) + "\n")
	os.Exit(1)
}

func FatalError(err error) {
	Fatalf("%s", err)
}

func main() {
	ciphertext := []byte{
		0x53, 0x9b, 0x33, 0x3b, 0x39, 0x70, 0x6d, 0x14,
		0x90, 0x28, 0xcf, 0xe1, 0xd9, 0xd4, 0xa4, 0x07,
	}
	aesKey := make([]byte, keySize)
	aesKey[0] = 0x80
	aesKey[keySize-1] = 1

	blockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		FatalError(err)
	}

	plaintext := make([]byte, blockSize)
	blockCipher.Decrypt(plaintext, ciphertext)
	fmt.Printf("% #x\n", plaintext)
}
