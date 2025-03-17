package models

type DeleteSong struct {
	ID    uint   `json:"id"`
	Group string `json:"group"`
	Song  string `json:"song"`
}
