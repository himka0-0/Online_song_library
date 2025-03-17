package models

type ResponseSongs struct {
	Data   []Songs `json:"data"`
	Length int     `json:"length"`
}
