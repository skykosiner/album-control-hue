package lights

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

type hueAPiColor struct {
	On bool      `json:"on"`
	Xy []float64 `json:"xy"`
}

func setColor(light int, url string, xy []float64) {
	jsonReq, err := json.Marshal(hueAPiColor{On: true, Xy: xy})

	if err != nil {
		log.Fatal("There was an error toggling the light", err)
	}

	url = fmt.Sprintf("%slights/%d/state", url, light)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonReq))

	if err != nil {
		log.Fatal("There was an issue turnning off the lights", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		log.Fatal(string(body))
	}

	defer resp.Body.Close()
}

func LightMeDaddy(album string) {
	url := "http://10.0.0.2/api/vGeourmApBqx37QJaJUQ4AxboqUjli1Fj3LtTQdY/"
	lights := []int{16, 1, 7, 14, 5}

	hue := map[string][]float64{
		"taylor-swift": {72, 171, 11},
		"speak-now":    {187, 18, 179},
		"1989":         {187, 187, 187},
		"reputation":   {255, 0, 0},
		"fearless":     {213, 195, 58},
		"lover":        {220, 20, 255},
		"folklore":     {255, 215, 217},
		"evermore":     {255, 215, 217},
		"red":          {255, 0, 0},
		"midnights":    {134, 134, 255},
	}

	for _, light := range lights {
		setColor(light, url, getXY(hue[album]))
	}
}

func getXY(info []float64) []float64 {
	if info[0] > float64(0.04045) {
		info[0] = math.Pow((info[0]+0.055)/(1.0+0.055), 2.4)
	} else {
		info[0] = (info[0] / 12.92)
	}

	if info[1] > 0.04045 {
		info[1] = math.Pow((info[1]+0.055)/(1.0+0.055), 2.4)
	} else {
		info[1] = (info[1] / 12.92)
	}

	if info[2] > 0.04045 {
		info[2] = math.Pow((info[2]+0.055)/(1.0+0.055), 2.4)
	} else {
		info[2] = (info[2] / 12.92)
	}

	X := info[0]*0.664511 + info[1]*0.154324 + info[2]*0.162028
	Y := info[0]*0.283881 + info[1]*0.668433 + info[2]*0.047685
	Z := info[0]*0.000088 + info[1]*0.072310 + info[2]*0.986039
	x := X / (X + Y + Z)
	y := Y / (X + Y + Z)

	return []float64{x, y}
}
