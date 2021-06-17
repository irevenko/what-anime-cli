package types

// Anilist - types for the anilist info
type Anilist struct {
	Id    int `json:"id"`
	IdMal int `json:"idMal"`
	Title struct {
		Native  string `json:"native"`
		Romaji  string `json:"romaji"`
		English string `json:"english"`
	} `json:"title"`
	Synonyms []string `json:"synonyms"`
	IsAdult  bool     `json:"isAdult"`
}

// Response - types for the trace.moe API
type Response struct {
	FrameCount int    `json:"frameCount"`
	Error      string `json:"error"`
	Result     []struct {
		Anilist    Anilist `json:"anilist"`
		Filename   string  `json:"filename"`
		Episode    int     `json:"episode"`
		From       float64 `json:"from"`
		To         float64 `json:"to"`
		Similarity float64 `json:"similarity"`
		Video      string  `json:"video"`
		Image      string  `json:"image"`
	} `json:"result"`
}
