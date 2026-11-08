package main

import (
	"fmt"
	"os"
	"os/exec"
)

func run(dir string) {
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	var choice int
	for {
		fmt.Println("\nChoose lab to run (0-4, 5 to exit):")
		fmt.Scan(&choice)
		if choice == 5 {
			break
		}
		switch choice {
		case 0:
			run("lab0")
		case 1:
			run("lab1")
		case 2:
			run("lab2")
		case 3:
			run("lab3")
		case 4:
			run("lab4")
		default:
			fmt.Println("Invalid choice")
		}
	}
}
