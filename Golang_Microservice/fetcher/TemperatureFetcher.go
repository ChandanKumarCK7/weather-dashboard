package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type MainData struct {
	Temp     float32 `json:"temp"`
	Humidity int     `json:"humidity"`
}

type CityData struct {
	Main           MainData `json:"main"`
	FetchedTime    int64    `json:"fetched_time"`
	FetchedTimeStr string   `json:"fetched_time_str"`
}

type WeatherData struct {
	Name     string   `json:"name"`
	CityData CityData `json:"City_Data"`
}

type TemperatureFetcher struct {
	Cities []string `json:"cities"` // to be set by outside services"
	ApiKey string   `json:"api_key"`
}

func (tf *TemperatureFetcher) ConstructData(cities []string, ApiKey string) {
	fmt.Println("constructing data with info ", cities, " ApiKey ", ApiKey)
	tf.Cities = cities
	tf.ApiKey = ApiKey
}
func (tf *TemperatureFetcher) FetchTemperature() string {
	if tf.Cities == nil {
		cities := []string{"hyderabad", "bangalore"}
		tf.SetCities(cities)
	}

	fmt.Println(tf.Cities)
	fmt.Printf("%T ", tf.Cities)

	fetchedWeatherData := tf.weatherFetcher(tf.Cities)
	fmt.Println("fetchedWeatherData", fetchedWeatherData)
	jsonData := toJson(fetchedWeatherData)
	fmt.Println("fetchedWeatherData in JSON format ", jsonData)

	return jsonData
}

func (tf *TemperatureFetcher) SetCities(cities []string) {
	tf.Cities = cities
}
func ToString(wd []WeatherData) string {
	result := "[\n"
	for _, data := range wd {
		result += fmt.Sprintf("   \"%s\" :{ \"main\": {\"temp\": %.2f, \"humidity\": %d}, \"fetched_time\": %d, \"fetched_time_str\": \"%s\"},\n",
			data.Name,
			data.CityData.Main.Temp,
			data.CityData.Main.Humidity,
			data.CityData.FetchedTime,
			data.CityData.FetchedTimeStr,
		)
	}
	result += "]\n"
	return result
}

func toJson(wd []WeatherData) string {
	fmt.Print("weather data at JSON marshaller ", wd)
	jsonData, err := json.Marshal(wd)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return err.Error()
	}
	return string(jsonData)
}

func (tf *TemperatureFetcher) weatherFetcher(cities []string) []WeatherData {
	var fetchedData []WeatherData

	for _, city := range cities {
		request := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + tf.ApiKey
		fmt.Println("request to be sent is ", request)

		response, err := http.Get(request)
		if err != nil {
			fmt.Println("error making HTTP request:", err)
			continue
		}
		defer response.Body.Close()

		resp, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("error reading response body:", err)
			continue
		}

		currWeatherData := WeatherData{}
		var CityData CityData
		err = json.Unmarshal(resp, &CityData)

		if err != nil {
			fmt.Println("error decoding JSON:", err)
			continue
		}

		fmt.Println("response ", CityData)
		currentTime := time.Now()
		CityData.FetchedTime = currentTime.Unix()
		CityData.FetchedTimeStr = currentTime.Format("2006-01-02 15:04:05")
		currWeatherData.CityData = CityData
		currWeatherData.Name = city

		fetchedData = append(fetchedData, currWeatherData)
	}

	return fetchedData
}
