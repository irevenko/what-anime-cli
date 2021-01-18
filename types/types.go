package types

type Response struct {
	Docs []struct {
		TitleRomanji string  `json:"title_romaji"`
		Similarity   float64 `json:"similarity"`
	} `json:"docs"`
}
