package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jeffting/nasa-api/entities"
	"github.com/jeffting/nasa-api/internal/busi"
	"github.com/jeffting/nasa-api/internal/clients"
)

func GetImagesDayHandler(clients clients.Clients) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		earthDate := r.URL.Query().Get("earth_date")
		if earthDate == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("earth_date is required query param"))
			return
		} else {
			// check date is in correct format
			_, err := time.Parse(dateLayout, earthDate)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("please use format YYYY-MM-DD"))
				return
			}
		}

		// extend api by getting other query params like "rover" or "camera" or "images_per_day",
		// but for now we'll use the defaults
		images, err := busi.GetImagesDay(ctx, clients, earthDate, defaultRover, defaultCamera, entities.UnsetImagesPerDay)
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

// func getDefaultEarthDate() string {
// 	defaultDate := time.Now().Add(-24 * defaultDaysAgo * time.Hour).Format(dateLayout)
// 	return defaultDate
// }
