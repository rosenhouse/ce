package main

import (
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

func display(data []byte) {
	fmt.Printf("%s\n", string(data))
}

var (
	plaintext1 = []byte{
		0x43, 0x72, 0x79, 0x70, 0x74, 0x6F, 0x67, 0x72,
		0x61, 0x70, 0x68, 0x79, 0x20, 0x43, 0x72, 0x79,
		0x70, 0x74, 0x6F, 0x67, 0x72, 0x61, 0x70, 0x68,
		0x79, 0x20, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6F,
	}

	ciphertext1 = []byte{
		0x46, 0x64, 0xdc, 0x06, 0x97, 0xbb, 0xfe, 0x69,
		0x33, 0x07, 0x15, 0x07, 0x9b, 0xa6, 0xc2, 0x3d,
		0x2b, 0x84, 0xde, 0x4f, 0x90, 0x8d, 0x7d, 0x34,
		0xaa, 0xce, 0x96, 0x8b, 0x64, 0xf3, 0xdf, 0x75,
	}

	ciphertext2 = []byte{
		0x51, 0x7e, 0xcc, 0x05, 0xc3, 0xbd, 0xea, 0x3b,
		0x33, 0x57, 0x0e, 0x1b, 0xd8, 0x97, 0xd5, 0x30,
		0x7b, 0xd0, 0x91, 0x6b, 0x8d, 0x82, 0x6b, 0x35,
		0xb7, 0x8b, 0xbb, 0x8d, 0x74, 0xe2, 0xc7, 0x3b,
	}
)

func xor(a, b []byte) []byte {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	dest := make([]byte, n)
	for i := 0; i < n; i++ {
		dest[i] = a[i] ^ b[i]
	}
	return dest
}

func main() {
	display(plaintext1)

	diff := xor(ciphertext1, ciphertext2)

	display(xor(diff, plaintext1))
}
