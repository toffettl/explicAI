package clients

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type BaseHttp struct {
	Client *resty.Request
}

func NewHttpClient(URL string, timeout int64) *BaseHttp {
	httpClient := resty.New().
	SetBaseURL(URL).
	SetTimeout(time.Duration(timeout) * time.Millisecond).R()

	return  &BaseHttp{
		Client: httpClient,
	}
}
