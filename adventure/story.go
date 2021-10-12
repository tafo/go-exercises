package adventure

type Chapter struct {
	Title   	string   `json:"title"`
	Paragraphs 	[]string `json:"story"`
	Options 	[]Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
