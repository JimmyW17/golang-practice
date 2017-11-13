package main

import (
	"fmt"
)

func main() {
	characterClasses := loadClasses()

	for i, class := range characterClasses {
		fmt.Println(i, class)
	}
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
