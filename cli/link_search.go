package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	helpers "github.com/irevenko/what-anime-cli/helpers"
	types "github.com/irevenko/what-anime-cli/types"
	"github.com/muesli/termenv"
)

const (
	linkSearchURL = "https://trace.moe/api/search?url="
)

// SearchByImageLink is for finding the anime scene by url which end with image file extension like: .jpg or .png
func SearchByImageLink(imageLink string) {
	_, err := url.ParseRequestURI(imageLink)
	if err != nil {
		log.Fatal("Invalid url")
	}

	termenv.HideCursor()
	defer termenv.ShowCursor()

	s := spinner.New(spinner.CharSets[33], 100*time.Millisecond)
	s.Prefix = "🔎 Searching for the anime: "
	s.FinalMSG = color.GreenString("✔️  Found!\n")

	go catchInterrupt(s)

	s.Start()

	reqBody, err := json.Marshal(map[string]string{})
	helpers.HandleError(err)

	resp, err := http.Post(linkSearchURL+imageLink, "application/json", bytes.NewBuffer(reqBody))
	helpers.HandleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	helpers.HandleError(err)

	var animeResp types.Response
	json.Unmarshal(body, &animeResp)

	s.Stop()

	fmt.Println("🌸 Title Native: " + animeResp.Docs[0].TitleNative)
	fmt.Println("🗻 Title Romaji: " + animeResp.Docs[0].TitleRomanji)
	fmt.Println("🗽 Title English: " + animeResp.Docs[0].TitleEnglish)
	fmt.Print("📊 Similarity: ")
	helpers.PrintAnimeSimilarity(strconv.FormatFloat(animeResp.Docs[0].Similarity, 'f', 6, 64))
	fmt.Println("📺 Episode Number: " + color.MagentaString(strconv.Itoa(animeResp.Docs[0].Episode)))
	fmt.Print("⌚ Scene At: " )
	helpers.PrintSceneAt(animeResp.Docs[0].At)
	fmt.Println("📅 Year & Season: " + color.CyanString(animeResp.Docs[0].Season))
	fmt.Print("🍓 Is Adult: ")
	helpers.PrintIsAdult(animeResp.Docs[0].IsAdult)
	//fmt.Println(string(body))
}
