package http

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/allegro/bigcache/v3"
)

const SwapiURL = "https://swapi.dev/api/"

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

// Swapi Interface
type SwapiClient interface {
	GetPlanetByName(name string) (*SwapiPlanet, error)
}

// Swapi represents a reference of SWAPI
type Swapi struct {
	client *RESTClient
	cache  *bigcache.BigCache
}

// NewSwapi creates a Swapi definition for https://swapi.dev/
func NewSwapi(swac *RESTClient) SwapiClient {
	imcache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(1 * time.Minute))

	swac.ApiUrl = SwapiURL
	return &Swapi{
		client: swac,
		cache:  imcache,
	}
}

// encodeParam adjust space and special characters for URL
func (s *Swapi) encodeParam(rawurl string) string {
	return url.QueryEscape(rawurl)
}

func (s *Swapi) getFromCache(name string) (*SwapiPlanet, error) {
	imjson, err := s.cache.Get(name)

	if err == nil {

		var swplnt SwapiPlanet
		err := json.Unmarshal([]byte(imjson), &swplnt)
		return &swplnt, err
	}

	return nil, err
}

func (s *Swapi) SetCache(key string, value SwapiPlanet) {

	json, err := json.Marshal(value)
	if err == nil {
		s.cache.Set(key, []byte(json))
	}
}

// Get a planet by a given name
// Ex : https://swapi.dev/api/planets/?search=Tatooine
func (s *Swapi) GetPlanetByName(name string) (*SwapiPlanet, error) {

	swp, _ := s.getFromCache(name)
	if swp != nil {
		return swp, nil
	}

	r := Response{}
	resp, err := s.client.Get("planets/?search="+s.encodeParam(name), r)

	if err != nil || resp.IsError() {
		return nil, err
	}

	re := resp.Result().(*Response)

	if re.Count > 0 {
		swp := re.Results[0]
		s.SetCache(name, swp)

		return &swp, err
	}

	return nil, err
}
