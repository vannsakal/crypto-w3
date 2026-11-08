package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/sha3"
)

// hashFunc defines a signature for hash computation functions
type hashFunc func(string) string

func hashMD5(input string) string {
	sum := md5.Sum([]byte(input))
	return hex.EncodeToString(sum[:])
}

func hashSHA1(input string) string {
	sum := sha1.Sum([]byte(input))
	return hex.EncodeToString(sum[:])
}

func hashSHA256(input string) string {
	sum := sha256.Sum256([]byte(input))
	return hex.EncodeToString(sum[:])
}

func hashSHA512(input string) string {
	sum := sha512.Sum512([]byte(input))
	return hex.EncodeToString(sum[:])
}

func hashSHA3(input string) string {
	sum := sha3.Sum256([]byte(input))
	return hex.EncodeToString(sum[:])
}

func proofMe(txt1, txt2 string) {
	algorithms := []struct {
		name string
		fn   hashFunc
	}{
		{"MD5", hashMD5},
		{"SHA1", hashSHA1},
		{"SHA256", hashSHA256},
		{"SHA512", hashSHA512},
		{"SHA3-256", hashSHA3},
	}

	fmt.Println("\n======== Comparison Results ========")
	for _, algo := range algorithms {
		outA := algo.fn(txt1)
		outB := algo.fn(txt2)

		fmt.Printf("Hash (%s):\n", algo.name)
		fmt.Printf("  Output A = %s\n", outA)
		fmt.Printf("  Output B = %s\n", outB)

		if outA == outB {
			fmt.Println("  => Match!")
		} else {
			fmt.Println("  => No Match!")
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("========Name + Hashing Program========")
	fmt.Print("Please input value 1: ")
	scanner.Scan()
	txt1 := strings.TrimRight(scanner.Text(), "\r\n")

	fmt.Print("Please input value 2: ")
	scanner.Scan()
	txt2 := strings.TrimRight(scanner.Text(), "\r\n")

	proofMe(txt1, txt2)
}
