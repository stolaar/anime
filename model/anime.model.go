package model

type Episode struct {
	Title       string
	Src         string
	Link        string
	Banner      string
	Description string
	AnimeId     string
}

type Anime struct {
	Id    string
	Title string
	Image string
	Link  string
}
