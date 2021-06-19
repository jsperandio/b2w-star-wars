package http

import (
	"net/url"
	"time"
)

// a Planet Schema from SWAPI
type SwapiPlanet struct {
	Name           string    `json:"name"`
	RotationPeriod string    `json:"rotation_period"`
	OrbitalPeriod  string    `json:"orbital_period"`
	Diameter       string    `json:"diameter"`
	Climate        string    `json:"climate"`
	Gravity        string    `json:"gravity"`
	Terrain        string    `json:"terrain"`
	SurfaceWater   string    `json:"surface_water"`
	Population     string    `json:"population"`
	Residents      []string  `json:"residents"`
	Films          []string  `json:"films"`
	Created        time.Time `json:"created"`
	Edited         time.Time `json:"edited"`
	URL            string    `json:"url"`
}

// Http Response from SWAPI Search
type Response struct {
	Count    int           `json:"count"`
	Next     interface{}   `json:"next"`
	Previous interface{}   `json:"previous"`
	Results  []SwapiPlanet `json:"results"`
}

type SwapiClient interface {
	GetPlanetByName(name string) (*SwapiPlanet, error)
}

// Swapi represents a reference of SWAPI
type Swapi struct {
	client *RESTClient
}

// NewSwapi creates a Swapi definition for https://swapi.dev/
func NewSwapi(swac *RESTClient) SwapiClient {
	return &Swapi{
		client: swac,
	}
}

func (s *Swapi) encodeParam(rawurl string) string {
	return url.QueryEscape(rawurl)
}

// Get a planet by a given name
// Ex : https://swapi.dev/api/planets/?search=Tatooine
func (s *Swapi) GetPlanetByName(name string) (*SwapiPlanet, error) {

	r := Response{}
	resp, err := s.client.Get("planets/?search="+s.encodeParam(name), r)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, err
	}

	re := resp.Result().(*Response)

	if re.Count > 0 {
		swp := re.Results[0]
		return &swp, err
	}

	return nil, err
}
