package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/homework_2/internal/infrastructure/http"
)

func NewAxilisOfferFeedJSON(httpClient stdhttp.Client) *http.AxilisOfferFeedJSON {
	return http.NewAxilisOfferFeedJSON(httpClient)
}
