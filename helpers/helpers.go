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
		_, err := fmt.Fprintln(color.Output, color.GreenString(similarity))
		HandleError(err)
	} else if similarity > "0.80" {
		_, err := fmt.Fprintln(color.Output, color.YellowString(similarity))
		HandleError(err)
	} else {
		_, err := fmt.Fprintln(color.Output, color.RedString(similarity))
		HandleError(err)
	}
}

// PrintSceneTime is for HH:MM:SS output
func PrintSceneTime(at float64) {
	h := math.Floor(at / 3600)
	m := math.Floor((at - h*3600) / 60)
	s := at - (h*3600 + m*60)
	_, err := fmt.Fprintln(color.Output, color.YellowString("%02.f:%02.f:%02.f", h, m, s))
	HandleError(err)
}

// PrintIsAdult is for colorful output
func PrintIsAdult(isAdult bool) {
	if isAdult == true {
		_, err := fmt.Fprintln(color.Output, color.GreenString("true"))
		HandleError(err)
	} else {
		_, err := fmt.Fprintln(color.Output, color.RedString("false"))
		HandleError(err)
	}
}
