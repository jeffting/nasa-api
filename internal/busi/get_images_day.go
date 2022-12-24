package busi

import (
	"context"

	"github.com/jeffting/nasa-api/entities"
	"github.com/jeffting/nasa-api/internal/clients"
)

// GetImagesDay gets images for a specific day
func GetImagesDay(ctx context.Context, clients clients.Clients, earthDate, rover, camera string, imagesPerDay int) (entities.Images, error) {
	images, err := clients.Nasa.GetImages(ctx, earthDate, rover, camera)
	if err != nil {
		return entities.Images{}, err
	}
	if imagesPerDay == entities.UnsetImagesPerDay {
		// if images per day not specified, return all images
		return images, nil
	}

	reducedImages := reduceImageList(imagesPerDay, images)

	return reducedImages, nil
}

func reduceImageList(imagesPerDay int, images entities.Images) entities.Images {
	reducedImages := entities.Images{}
	for key, photos := range images {
		reducedList := []string{}
		for i, photo := range photos {
			if i >= imagesPerDay {
				break
			}
			reducedList = append(reducedList, photo)
		}
		reducedImages[key] = reducedList
	}
	return reducedImages
}
