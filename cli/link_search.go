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
	linkSearchURL = "https://api.trace.moe/search?anilistInfo&url="
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

	fmt.Println("🌸 Title Native:", animeResp.Result[0].Anilist.Title.Native)
	fmt.Println("🗻 Title Romaji:", animeResp.Result[0].Anilist.Title.Romaji)
	fmt.Println("🗽 Title English:", animeResp.Result[0].Anilist.Title.English)
	fmt.Print("📊 Similarity: ")
	helpers.PrintAnimeSimilarity(strconv.FormatFloat(animeResp.Result[0].Similarity, 'f', 6, 64))
	_, err = fmt.Fprintln(color.Output, "📺 Episode Number: "+color.MagentaString(strconv.Itoa(animeResp.Result[0].Episode)))
	helpers.HandleError(err)
	fmt.Print("⌚ Scene From: ")
	helpers.PrintSceneTime(animeResp.Result[0].From)
	fmt.Print("⌚ Scene To: ")
	helpers.PrintSceneTime(animeResp.Result[0].To)
	fmt.Print("🍓 Is Adult: ")
	helpers.PrintIsAdult(animeResp.Result[0].Anilist.IsAdult)
	//fmt.Println(string(body))
}
