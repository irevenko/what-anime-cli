package cli

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
	helpers "github.com/irevenko/what-anime-cli/helpers"
	types "github.com/irevenko/what-anime-cli/types"
)

// SearchByImageFile is for finding the anime scene by existing image file
func SearchByImageFile(imagePath string) {
	fileSearchURL := "https://trace.moe/api/search"
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Searching for the anime: "
	s.Start()

	imageFile, err := os.Open(imagePath)
	helpers.HandleError(err)

	reader := bufio.NewReader(imageFile)
	content, err := ioutil.ReadAll(reader)
	helpers.HandleError(err)

	encodedImage := base64.StdEncoding.EncodeToString(content)

	reqBody, err := json.Marshal(map[string]string{"image": encodedImage})
	helpers.HandleError(err)

	resp, err := http.Post(fileSearchURL, "application/json", bytes.NewBuffer(reqBody))
	helpers.HandleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	helpers.HandleError(err)

	var animeResp types.Response
	json.Unmarshal(body, &animeResp)

	s.Stop()

	fmt.Println(animeResp.Docs[0].TitleRomanji)
	fmt.Println(animeResp.Docs[0].Similarity)
	// fmt.Println(string(body))
}
