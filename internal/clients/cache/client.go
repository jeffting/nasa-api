package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jeffting/nasa-api/entities"
)

const cacheFolder = "./cache"

type CacheInterface interface {
	Set(context.Context, string, entities.Images) error
	Get(context.Context, string) (entities.Images, error)
}

type Cache struct{}

func NewClient() Cache {
	return Cache{}
}

func (c Cache) Set(ctx context.Context, key string, images entities.Images) error {
	cacheKey := fmt.Sprintf("%s/%s", cacheFolder, key)
	file, err := json.MarshalIndent(images, "", " ")
	if err != nil {
		log.Print("error setting up cache set")
	}

	err = ioutil.WriteFile(cacheKey, file, 0644)
	if err != nil {
		log.Print("error writing to cache file")
	}
	return err
}

func (c Cache) Get(ctx context.Context, key string) (entities.Images, error) {
	cacheKey := fmt.Sprintf("%s/%s", cacheFolder, key)
	file, err := ioutil.ReadFile(cacheKey)
	if err != nil {
		log.Print("cache miss", err)
		return entities.Images{}, err
	}

	images := entities.Images{}

	err = json.Unmarshal([]byte(file), &images)
	if err != nil {
		log.Print("cache error", err)
		return entities.Images{}, err
	}
	fmt.Println("cache hit")
	return images, nil
}
