package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	characterClasses := loadClasses()

	for i, class := range characterClasses {
		fmt.Println(i, class)
	}

	charName := getName()
	fmt.Print("Welcome, " + charName)
}

type characterClasses []string

func loadClasses() characterClasses {
	charClasses := characterClasses{}
	classes := []string{"Warrior", "Rogue", "Mage", "Healer"}
	for _, class := range classes {
		charClasses = append(charClasses, class)
	}
	return charClasses
}

func getName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	text, _ := reader.ReadString('\n')
	return text
}
