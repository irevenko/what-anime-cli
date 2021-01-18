package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/briandowns/spinner"
	helpers "github.com/irevenko/what-anime-cli/helpers"
)

type Response struct {
	Docs []struct {
		TitleRomanji string  `json:"title_romaji"`
		Similarity   float64 `json:"similarity"`
	} `json:"docs"`
}

func SearchByImageLink(imageLink string) {
	linkSearchURL := "https://trace.moe/api/search?url="
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Searching for the anime: "
	s.Start()

	reqBody, err := json.Marshal(map[string]string{})
	helpers.HandleError(err)
	resp, err := http.Post(linkSearchURL+imageLink, "application/json", bytes.NewBuffer(reqBody))
	helpers.HandleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	helpers.HandleError(err)

	var animeResp Response
	json.Unmarshal(body, &animeResp)

	s.Stop()

	fmt.Println(animeResp.Docs[0].TitleRomanji)
	fmt.Println(animeResp.Docs[0].Similarity)
	//fmt.Println(string(body))
}
