package points

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"cli_weather/latlong"
)

type Coordinates struct {
	Properties Coordinate `json:"properties"`
}

type Coordinate struct {
	GridId string `json:"gridId"`
	GridX int `json:"gridX"`
	GridY int `json:"gridY"`
}

func GetPoints(state string, city string) string {
	
	latlng := latlong.GetLatLong(state, city)

	wPoints := "https://api.weather.gov/points/"+latlng

	wpr, err := http.Get(wPoints)
	if err != nil {
		fmt.Println(err)
	}

	defer wpr.Body.Close()

	wprBody, err := ioutil.ReadAll(wpr.Body)
	if err != nil {
		fmt.Println(err)
	}

	var coord Coordinates

	json.Unmarshal(wprBody, &coord)

	gridId := coord.Properties.GridId
	gridX := coord.Properties.GridX
	gridY := coord.Properties.GridY

	grx := strconv.Itoa(gridX)
	gry := strconv.Itoa(gridY)

	grid := gridId+"/"+grx+","+gry

	return grid
}