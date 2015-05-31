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
	plaintext := []byte{
		0x29, 0x6c, 0x93, 0xfd, 0xf4, 0x99, 0xaa, 0xeb,
		0x41, 0x94, 0xba, 0xbc, 0x2e, 0x63, 0x56, 0x1d,
	}
	aesKey := make([]byte, keySize)
	aesKey[0] = 0x80
	aesKey[keySize-1] = 1

	blockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		FatalError(err)
	}

	ciphertext := make([]byte, blockSize)
	blockCipher.Encrypt(ciphertext, plaintext)
	fmt.Printf("% #x\n", ciphertext)
}
