package main

import (
	"bufio"
	"fmt"
	"os"

	"lab2/utils/crack"
)

func main() {
	target := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

	file, err := os.Open("nord_vpn.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	outFile, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating verbose.txt:", err)
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password := scanner.Text()
		hash := crack.HashSHA1(password)

		msg := fmt.Sprintf("Trying: %s -> %s\n", password, hash)
		fmt.Print(msg)
		outFile.WriteString(msg)

		if hash == target {
			found := fmt.Sprintf("\nPassword found: %s\n", password)
			fmt.Print(found)
			outFile.WriteString(found)
			break
		}
	}
}
