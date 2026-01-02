package util

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ilmsg/fluffy-octo-tribble/model"
)

func GetDataResponse(w http.ResponseWriter, status int, message string, data any) {
	dataRes := model.DataResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(dataRes)
}

func PageLimit(r *http.Request) (page int, limit int) {
	page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	return
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Logic to execute BEFORE the next handler
		// log.Printf("Started request: %s %s", r.Method, r.RequestURI)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Logic to execute AFTER the next handler completes
		log.Printf("Finished request: %s %s in %v", r.Method, r.RequestURI, time.Since(start))
	})
}
