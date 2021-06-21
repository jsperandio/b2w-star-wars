package http

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type RESTClient struct {
	apiUrl string
	client resty.Client
}

// NewRestClient will create new an RESTClient for given api url
func NewRESTClient(url string, maxRetry int, secWait int, maxWaitSec int) *RESTClient {
	c := *resty.New()
	c.SetRetryCount(maxRetry).
		SetRetryWaitTime(time.Duration(secWait) * time.Second).
		SetRetryMaxWaitTime(time.Duration(maxWaitSec) * time.Second)

	return &RESTClient{
		apiUrl: url,
		client: c,
	}
}

func (rc *RESTClient) Get(route string, automarshal interface{}) (result *resty.Response, err error) {
	return rc.client.R().
		SetResult(automarshal).
		ForceContentType("application/json").
		Get(rc.apiUrl + route)
}
