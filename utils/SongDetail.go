package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online_song/logger"
	"os"
)

// структура из свагера по ТЗ
type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func GetSongInfo(group, song string) (*SongDetail, error) {
	IncompleteAPIURL := os.Getenv("Your_URL")
	if IncompleteAPIURL == "" {
		logger.Logger.Warn("Вы не заполнили env с url внешнего API")
		newSong := SongDetail{ReleaseDate: "не заполнено", Text: "не заполнено", Link: "не заполнено"}
		return &newSong, nil
	}
	url := fmt.Sprintf("%s/info?group=%s&song=%s", IncompleteAPIURL, group, song)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()
	var detail SongDetail
	if err = json.NewDecoder(resp.Body).Decode(&detail); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	return &detail, nil
}
