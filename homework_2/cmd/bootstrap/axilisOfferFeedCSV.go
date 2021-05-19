package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/homework_2/internal/infrastructure/http"
)

func NewAxilisOfferFeedCSV() *http.AxilisOfferFeedCSV {
	return http.NewAxilisOfferFeedCSV(stdhttp.Client{})
}
