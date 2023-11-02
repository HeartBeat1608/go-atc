package lib_test

// write tests for the metar module to test core functionality

import (
	"go_atc/lib"
	"testing"
)

func TestNewMetar(t *testing.T) {
	icao := "EGLL"
	metar, err := lib.NewMetar(icao)
	if err != nil {
		t.Errorf("Error creating new Metar: %s", err)
	}

	if metar == nil {
		t.Errorf("Expected metar to be non-nil")
	}

	if metar != nil && metar.Icao != icao {
		t.Errorf("Expected metar to be %s, got %s", icao, metar.Icao)
	}
}

func TestMetarRefresh(t *testing.T) {
	icao := "EGLL"
	metar, err := lib.NewMetar(icao)
	if err != nil {
		t.Errorf("Error creating new Metar: %s", err)
	}

	if err := metar.Refresh(); err != nil {
		t.Errorf("Error refreshing metar: %s", err)
	}

	if metar.Icao != icao {
		t.Errorf("Expected metar to be %s, got %s", icao, metar.Icao)
	}
}

func TestLoadCurrentWeather(t *testing.T) {
	icao := "EGLL"
	metar, err := lib.NewMetar(icao)
	if err != nil {
		t.Errorf("Error creating new Metar: %s", err)
	}

	if err := metar.Refresh(); err != nil {
		t.Errorf("Error refreshing metar: %s", err)
	}

	weather := metar.GetCurrentWeather()
	if weather.IcaoId == "" {
		t.Errorf("Expected weather to be non-empty")
	}

	clouds := metar.GetClouds()
	if len(clouds) == 0 {
		t.Errorf("Expected clouds to be non-empty")
	}
}

func TestSimpleWeather(t *testing.T) {
	icao := "EGLL"
	metar, err := lib.NewMetar(icao)
	if err != nil {
		t.Errorf("Error creating new Metar: %s", err)
	}

	if err := metar.Refresh(); err != nil {
		t.Errorf("Error refreshing metar: %s", err)
	}

	weather, err := metar.GetSimpleWeather()
	if err != nil {
		t.Errorf("Error getting simple weather: %s", err)
	}

	if weather.IcaoId == "" {
		t.Errorf("Expected weather to be non-empty")
	}
}
