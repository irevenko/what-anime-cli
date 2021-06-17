package heplers

import (
	"fmt"
	"log"
	"math"

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

// PrintSceneTime is for HH:MM:SS output
func PrintSceneTime(at float64) {
	h := math.Floor(at / 3600)
	m := math.Floor((at - h*3600) / 60)
	s := at - (h*3600 + m*60)
	fmt.Println(color.YellowString("%02.f:%02.f:%02.f", h, m, s))
}

// PrintIsAdult is for colorful output
func PrintIsAdult(isAdult bool) {
	if isAdult == true {
		fmt.Println(color.GreenString("true"))
	} else {
		fmt.Println(color.RedString("false"))
	}
}
