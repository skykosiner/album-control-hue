package albums

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Album struct {
	Title string
	Id    string
}

func (a *Album) GetAlbums() []Album {
	var albums []Album
	bytes, err := os.ReadFile(fmt.Sprintf("%s/personal/taylor_albums", os.Getenv("HOME")))

	if err != nil {
		log.Fatal("Error getting taylors albums")
	}

	albumStrArr := strings.Split(string(bytes), "\n")

	for _, album := range albumStrArr {
		strArr := strings.Split(album, "id: ")
		if album != "" {
			albums = append(albums, Album{strArr[0], strArr[1]})
		}
	}

	return albums
}
