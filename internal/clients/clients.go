package clients

import (
	"github.com/jeffting/nasa-api/internal/clients/cache"
	"github.com/jeffting/nasa-api/internal/clients/nasa"
)

type Clients struct {
	Nasa  nasa.NasaCaller
	Cache cache.CacheInterface
}

// client intialization
func InitializeClients() Clients {
	apiKey := "DEMO_KEY" // should store api key as secret when time comes
	return Clients{
		Nasa:  nasa.NewClient("https://api.nasa.gov", apiKey),
		Cache: cache.NewClient(),
	}
}
