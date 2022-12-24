package nasa

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/jeffting/nasa-api/entities"
)

type NasaCaller interface {
	GetImages(context.Context, string, string, string) (entities.Images, error)
}

type Nasa struct {
	domain     string
	apiKey     string
	httpClient http.Client
}

func NewClient(domain string, apiKey string) Nasa {
	return Nasa{domain, apiKey, *http.DefaultClient}
}

func (c Nasa) GetImages(ctx context.Context, earthDate, rover, camera string) (entities.Images, error) {
	params := url.Values{
		"earth_date": {earthDate},
		"camera":     {camera},
		"api_key":    {c.apiKey},
	}
	endpoint := c.domain + "/mars-photos/api/v1/rovers/" + rover + "/photos?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Print("error creating request", err)
		return entities.Images{}, err
	}

	var images images

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Print("api request error", err)
		return entities.Images{}, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&images); err != nil {
		log.Print("decode error", err)
		return entities.Images{}, err
	}

	return fromImagesToDomain(images, earthDate), err
}
