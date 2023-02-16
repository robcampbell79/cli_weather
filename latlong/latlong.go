package latlong

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Places struct {
	Place []LongLat `json:"places"`
}

type LongLat struct {
	Lat string `json:"latitude"`
	Lng string `json:"longitude"`
}

func GetLatLong(state string, city string) (string) {
	zippo := "http://api.zippopotam.us/us/"+state+"/"+city

	zr, err := http.Get(zippo)
	if err != nil {
		fmt.Println(err)
	}

	defer zr.Body.Close()

	zBody, err := ioutil.ReadAll(zr.Body)
	if err != nil {
		fmt.Println(err)
	}

	var pl Places

	json.Unmarshal(zBody, &pl)

	var lat string
	var lng string

	lat = pl.Place[0].Lat
	lng = pl.Place[0].Lng

	latlng := lat+","+lng

	return latlng
}