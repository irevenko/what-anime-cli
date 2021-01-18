package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

type Response struct {
	Docs []struct {
		TitleRomanji string  `json:"title_romaji"`
		Similarity   float64 `json:"similarity"`
	} `json:"docs"`
}

const (
	fileSearchURL = "https://trace.moe/api/search"
	linkSearchURL = "https://trace.moe/api/search?url="
)

var (
	imagePath string
	imageLink string
)

func main() {
	flag.StringVar(&imagePath, "i", "", "Anime Image Path")
	flag.StringVar(&imageLink, "l", "", "Anime Image Link")
	flag.Parse()

	if imagePath != "" {
		searchByImageFile(imagePath)
	}
	if imageLink != "" {
		searchByImageLink(imageLink)
	}

}

func searchByImageFile(imagePath string) {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Searching for the anime: "
	s.Start()

	imageFile, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(imageFile)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	encodedImage := base64.StdEncoding.EncodeToString(content)

	reqBody, err := json.Marshal(map[string]string{
		"image": encodedImage,
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(fileSearchURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var animeResp Response
	json.Unmarshal(body, &animeResp)

	s.Stop()

	fmt.Println(animeResp.Docs[0].TitleRomanji)
	fmt.Println(animeResp.Docs[0].Similarity)
	// fmt.Println(string(body))
}

func searchByImageLink(imageLink string) {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Searching for the anime: "
	s.Start()

	reqBody, _ := json.Marshal(map[string]string{})
	resp, err := http.Post(linkSearchURL+imageLink, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var animeResp Response
	json.Unmarshal(body, &animeResp)

	s.Stop()

	fmt.Println(animeResp.Docs[0].TitleRomanji)
	fmt.Println(animeResp.Docs[0].Similarity)
	//fmt.Println(string(body))
}
