package forecast

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"cli_weather/points"
)

type Properties struct {
	Periods WeatherArr `json:"properties"`
}

type WeatherArr struct {
	Forecasts []Weather `json:"periods"`
}

type Weather struct {
	Name string `json:"name"`
	Temp int `json:"temperature"`
	Forecast string `json:"detailedForecast"`
}

func GetForecast(state string, city string) []Weather {

	pnts := points.GetPoints(state, city)

	resp, err := http.Get("https://api.weather.gov/gridpoints/"+pnts+"/forecast")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var w Properties

	json.Unmarshal(body, &w)

	return w.Periods.Forecasts
}