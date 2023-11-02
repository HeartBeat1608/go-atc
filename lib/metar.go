package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MetarClouds struct {
	CloudCover string `json:"cover"`
	CloudBase  int    `json:"base"`
}

type MetarWeather struct {
	MetarId        int           `json:"metar_id"`
	IcaoId         string        `json:"icaoId"`
	ReportTime     string        `json:"reportTime"`
	Temperature    int           `json:"temp"`
	DewPoint       int           `json:"dewp"`
	WindDirection  int           `json:"wdir"`
	WindSpeed      int           `json:"wspd"`
	WindGust       int           `json:"wgust"`
	Altimeter      int           `json:"altim"`
	Visibility     string        `json:"visib"`
	Snow           string        `json:"snow"`
	RawObservation string        `json:"rawOb"`
	Latitude       float32       `json:"lat"`
	Longitude      float32       `json:"lon"`
	Elevation      int           `json:"elev"`
	AirportName    string        `json:"name"`
	Clouds         []MetarClouds `json:"clouds"`
}

type SimpleMetarWeather struct {
	IcaoId        string `json:"icaoId"`
	Temperature   int    `json:"temp"`
	DewPoint      int    `json:"dewp"`
	WindDirection int    `json:"wdir"`
	WindSpeed     int    `json:"wspd"`
	WindGust      int    `json:"wgust"`
	AirportName   string `json:"name"`
}

type Metar struct {
	Icao          string
	LatestWeather MetarWeather
}

func NewMetar(icao string) (*Metar, error) {
	return &Metar{
		Icao:          icao,
		LatestWeather: MetarWeather{},
	}, nil
}

func (m Metar) Pretty() string {
	bfr, err := json.Marshal(m.LatestWeather)
	if err != nil {
		return err.Error()
	}
	stringified := string(bfr)
	return stringified
}

func (m *Metar) Refresh() error {
	url := fmt.Sprintf("https://aviationweather.gov/api/data/metar?ids=%s&format=json&taf=false", m.Icao)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	var weatherResponse []MetarWeather
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(respBytes, &weatherResponse); err != nil {
		return err
	}

	m.LatestWeather = weatherResponse[0]
	return nil
}

func (m Metar) GetCurrentWeather() MetarWeather {
	return m.LatestWeather
}

func (m Metar) GetStringifiedWeather() string {
	res := fmt.Sprintf("ICAO: %s\nTemperature: %d\nDew Point: %d\nAltimeter: %d\nAirport Name: %s\nWinds: %d Knots at %d, gusting at %d\nVisibility: %s\nElevation: %d\n", m.Icao, m.LatestWeather.Temperature, m.LatestWeather.DewPoint, m.LatestWeather.Altimeter, m.LatestWeather.AirportName, m.LatestWeather.WindSpeed, m.LatestWeather.WindDirection, m.LatestWeather.WindGust, m.LatestWeather.Visibility, m.LatestWeather.Elevation)
	return res
}

func (m Metar) GetClouds() []MetarClouds {
	return m.LatestWeather.Clouds
}

func (m Metar) GetSimpleWeather() (SimpleMetarWeather, error) {
	var err error
	var b []byte
	var res SimpleMetarWeather

	if b, err = json.Marshal(m.LatestWeather); err != nil {
		return SimpleMetarWeather{}, err
	}

	if err = json.Unmarshal(b, &res); err != nil {
		return SimpleMetarWeather{}, err
	}

	return res, nil
}
