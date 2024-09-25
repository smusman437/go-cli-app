package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Go CLI App!")
	mainMenu(reader)
}

func mainMenu(reader *bufio.Reader) {
	fmt.Println("Main Menu:")
	fmt.Println("1. Greet User")
	fmt.Println("2. List Files")
	fmt.Println("3. Show Date")
	fmt.Println("Enter the number of your choice:")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		greetUserSection(reader)
	case "2":
		listFilesSection()
	case "3":
		showDateSection()
	default:
		fmt.Println("Invalid choice, please try again.")
		mainMenu(reader)
	}
}

func greetUserSection(reader *bufio.Reader) {
	fmt.Println("Enter your name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Hello, %s! How are you feeling today?\n", name)
	fmt.Println("1. Happy")
	fmt.Println("2. Sad")
	fmt.Println("3. Excited")

	mood, _ := reader.ReadString('\n')
	mood = strings.TrimSpace(mood)

	switch mood {
	case "1":
		fmt.Println("Great to hear you're happy!")
	case "2":
		fmt.Println("Sorry to hear you're feeling sad. Hope things get better.")
	case "3":
		fmt.Println("Exciting times! Stay positive!")
	default:
		fmt.Println("Invalid mood, but I hope you're feeling well.")
	}

	fmt.Println("Returning to Main Menu...")
	mainMenu(reader)
}

func listFilesSection() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Files in current directory:")
	for _, file := range files {
		fmt.Println(file.Name())
	}

	fmt.Println("Returning to Main Menu...")
	mainMenu(bufio.NewReader(os.Stdin))
}

func showDateSection() {
	fmt.Println("Current Date:", strings.Split(time.Now().String(), " ")[0])

	fmt.Println("Returning to Main Menu...")
	mainMenu(bufio.NewReader(os.Stdin))
}
