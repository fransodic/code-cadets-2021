package bootstrap

import (
	stdhttp "net/http"
	"time"
)

func NewHttpClient() *stdhttp.Client {
	return &stdhttp.Client{
		Timeout: time.Second * 10,
	}
}
