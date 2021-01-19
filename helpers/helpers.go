package heplers

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

// HandleError is a reusable error checking function
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// PrintAnimeSimilarity is for colorful output
func PrintAnimeSimilarity(similarity string) {
	if similarity > "0.89" {
		fmt.Println(color.GreenString(similarity))
	} else if similarity > "0.80" {
		fmt.Println(color.YellowString(similarity))
	} else {
		fmt.Println(color.RedString(similarity))
	}
}

// PrintIsAdult is for colorful output
func PrintIsAdult(isAdult bool) {
	if isAdult == true {
		fmt.Println(color.GreenString("true"))
	} else {
		fmt.Println(color.RedString("false"))
	}
}
