package nasa

import (
	"github.com/jeffting/nasa-api/entities"
)

const dateLayout = "2006-01-02"

func fromImagesToDomain(images images, date string) entities.Images {
	domainImages := entities.Images{date: []string{}}

	for _, photo := range images.Photos {
		if v, ok := domainImages[photo.EarthDate]; ok {
			domainImages[photo.EarthDate] = append(v, photo.ImgSrc)
			continue
		}
		photoList := []string{photo.ImgSrc}
		domainImages[photo.EarthDate] = photoList
	}

	return domainImages
}
