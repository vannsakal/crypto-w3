package main

import (
	"bufio"
	"fmt"
	"os"

	"lab3/utils/crack"
)

func main() {
	target := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

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
		hash := crack.HashSHA512(password)

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
