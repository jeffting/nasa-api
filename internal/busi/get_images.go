package busi

import (
	"context"
	"sync"
	"time"

	"github.com/jeffting/nasa-api/entities"
	"github.com/jeffting/nasa-api/internal/clients"
)

const (
	defaultDaysAgo = 10
	dateLayout     = "2006-01-02"
)

// GetImages gets images for the last 10 days
func GetImages(ctx context.Context, clients clients.Clients, rover, camera string, imagesPerDay int) (entities.Images, error) {

	// check cache
	cacheKey := time.Now().Add(-24*defaultDaysAgo*time.Hour).Format(dateLayout) + ".json"
	cachedImages, err := clients.Cache.Get(ctx, cacheKey)
	if err == nil && len(cachedImages) > 0 {
		return cachedImages, nil
	}

	collectedImages := entities.Images{}
	// error channel for seeing if errors in goroutines
	errChan := make(chan error)
	var wg sync.WaitGroup
	// async calls to get the last 10 days
	for i := defaultDaysAgo; i > 0; i-- {
		wg.Add(1)
		go func(daysAgo int, ec chan error) {
			defer wg.Done()
			earthDate := time.Now().Add(-24 * time.Duration(daysAgo) * time.Hour).Format(dateLayout)

			images, err := GetImagesDay(ctx, clients, earthDate, rover, camera, imagesPerDay)
			if err != nil {
				ec <- err
			}

			collectedImages[earthDate] = images[earthDate]
		}(i, errChan)
	}
	wg.Wait()

	close(errChan)

	select {
	case chanErr := <-errChan:
		if chanErr != nil {
			return entities.Images{}, chanErr
		}
	}

	// set cache
	clients.Cache.Set(ctx, cacheKey, collectedImages)

	return collectedImages, nil
}
