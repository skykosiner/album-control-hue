package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/skykosiner/taylor-swift-albums/pkg/albums"
	"github.com/skykosiner/taylor-swift-albums/pkg/lights"
)

func main() {
	if len(os.Args[1:]) <= 0 {
		log.Fatal("Zero args passed in")
	}

	var a *albums.Album = &albums.Album{}
	albumsArr := a.GetAlbums()

	for _, value := range albumsArr {
		if os.Args[1:][0] == strings.TrimSuffix(value.Title, " ") {
			lights.LightMeDaddy(os.Args[1:][0])
			cmd := fmt.Sprintf(`qdbus org.mpris.MediaPlayer2.spotify \
            /org/mpris/MediaPlayer2 \
            org.mpris.MediaPlayer2.Player.OpenUri \
            spotify:album:%s
            `, value.Id)

			play := exec.Command("bash", "-c", cmd)

			_, err := play.Output()

			if err != nil {
				log.Fatal("Error playing album")
			}
		}
	}
}
