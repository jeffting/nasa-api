package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jeffting/nasa-api/internal/busi"
	"github.com/jeffting/nasa-api/internal/clients"
)

const (
	defaultRover        = "curiosity"
	defaultCamera       = "NAVCAM"
	defaultImagesPerDay = 3
	dateLayout          = "2006-01-02"
)

func GetImagesHandler(clients clients.Clients) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// extend api by getting other query params like "rover" or "camera" or "images_per_day",
		// but for now we'll use the defaults

		images, err := busi.GetImages(ctx, clients, defaultRover, defaultCamera, defaultImagesPerDay)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(images)
	}
	return http.HandlerFunc(fn)
}
