package cmd

import (
	"encoding/json"
	"fmt"
)

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type Output struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

func kelvinToCelsius(kStr float64) string {

	c := kStr - 273.15
	return fmt.Sprintf("%.2f", c)
}

func ParseOutput(data string) (string, error) {
	var output Output
	err := json.Unmarshal([]byte(data), &output)
	if err != nil {
		return "", err
	}

	formatted := fmt.Sprintf("Coordinates \U0001F5FA :\n  Lon: %v\n  Lat: %v\n", output.Coord.Lon, output.Coord.Lat)

	formatted += "Weather:\n"
	for _, w := range output.Weather {
		//formatted += fmt.Sprintf("  Id: %v\n", w.Id)
		formatted += fmt.Sprintf("  Main: %v\n", w.Main)
		formatted += fmt.Sprintf("  Description: %v\n", w.Description)
		//formatted += fmt.Sprintf("  Icon: %v\n", w.Icon)
	}

	formatted += fmt.Sprintf("Base: %v\n", output.Base)
	//celsius := kelvinToCelsius(temp)
	formatted += fmt.Sprintf("The weather \U0001F321 :\n  Temp: %v\n  FeelsLike: %v\n  TempMin: %v\n  TempMax: %v\n  Pressure: %v\n  Humidity: %v\n",
		kelvinToCelsius(output.Main.Temp), kelvinToCelsius(output.Main.FeelsLike), kelvinToCelsius(output.Main.TempMin), kelvinToCelsius(output.Main.TempMax), output.Main.Pressure, output.Main.Humidity)

	formatted += fmt.Sprintf("Visibility: %v meters \n", output.Visibility)

	formatted += fmt.Sprintf("Wind \U0001F32C :\n  Speed: %v\n  Deg: %v\n  Gust: %v\n", output.Wind.Speed, output.Wind.Deg, output.Wind.Gust)

	//formatted += fmt.Sprintf("Clouds:\n  All: %v\n", output.Clouds.All)

	//formatted += fmt.Sprintf("Dt: %v\n", output.Dt)

	//formatted += fmt.Sprintf("Sys:\n  Type:%v\n  Id: %v\n  Country: %v\n  Sunrise: %v\n  Sunset: %v\n",
	//	output.Sys.Type, output.Sys.Id, output.Sys.Country, output.Sys.Sunrise, output.Sys.Sunset)
	formatted += fmt.Sprintf("Country: %v\n", output.Sys.Country)
	//formatted += fmt.Sprintf("Timezone: %v\n", output.Timezone)
	//formatted += fmt.Sprintf("Id: %v\n", output.Id)
	formatted += fmt.Sprintf("Name: %v\n", output.Name)
	//formatted += fmt.Sprintf("Cod: %v\n", output.Cod)

	return formatted, nil
}
