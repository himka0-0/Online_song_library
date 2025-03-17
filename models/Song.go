package models

type Songs struct {
	ID          uint   `gorm:"primarykey"`
	Muzgroup    string `gorm:"not null"`
	Song        string `gorm:"not null"`
	ReleaseDate string `gorm:"size 100"`
	Text        string `gorm:"text"`
	Link        string `gorm:"text"`
}
