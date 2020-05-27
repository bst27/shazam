package shazam

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// FetchCityCharts returns Shazam charts (= a list of tracks) for the given city. Tracks are ordered
// ascending by their chart position. You can use shazam.Cities() to lookup city IDs by name.
func FetchCityCharts(cityID string) ([]Track, error) {
	url := "https://www.shazam.com/shazam/v3/de/DE/web/-/tracks/sma2-chart-" + cityID + "?pageSize=50&startFrom=0"

	client := http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := response{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return response.Tracks, nil
}
