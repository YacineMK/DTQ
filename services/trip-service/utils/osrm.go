// utils/osrm.go
package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/YacineMK/DTQ/services/trip-service/pkg/types"
	"github.com/YacineMK/DTQ/shared/env"
)

var (
	baseURL = env.GetEnv("OSRM_API_ENDPOINT", "http://localhost:5000")
	profile = "driving"
)

func GetRoute(coords string) (*types.OsrmApiResponse, error) {
	endpoint := fmt.Sprintf("%s/route/v1/%s/%s?geometries=geojson", baseURL, profile, url.PathEscape(coords))

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data types.OsrmApiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
