package main

import (
	"flag"

	cli "github.com/irevenko/what-anime-cli/cli"
)

var (
	imagePath string
	imageLink string
)

func main() {
	flag.StringVar(&imagePath, "f", "", "Anime Image File Path")
	flag.StringVar(&imageLink, "l", "", "Anime Image Link")
	flag.Parse()

	if imagePath != "" {
		cli.SearchByImageFile(imagePath)
	}
	if imageLink != "" {
		cli.SearchByImageLink(imageLink)
	}

}
