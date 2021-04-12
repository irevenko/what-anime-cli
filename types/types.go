package types

// Response - types for the trace.moe API
type Response struct {
	Docs []struct {
		TitleRomanji string  `json:"title_romaji"`
		TitleEnglish string  `json:"title_english"`
		TitleNative  string  `json:"title_native"`
		Similarity   float64 `json:"similarity"`
		Episode      int     `json:"episode"`
		At           float64 `json:"at"`
		Season       string  `json:"season"`
		IsAdult      bool    `json:"is_adult"`
	} `json:"docs"`
}
