package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
)

func hashMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	query := "%62%36%63%63%62%34%65%63%65%35%34%35%34%64%63%61%65%35%31%37%37%38%62%33%65%32%33%39%65%62%63%32"
	decoded, _ := url.QueryUnescape(query)

	fmt.Println("Decoded Hash:", decoded)

	// words related to cats
	words := []string{"cat", "cats", "kitty", "kitten", "meow", "meowmeow"}

	var password string
	for _, w := range words {
		if hashMD5(w) == decoded {
			password = w
			break
		}
	}

	fmt.Println("Password found:", password)
	fmt.Printf("Flag: cryptoCTF{%s}\n", password)
}
