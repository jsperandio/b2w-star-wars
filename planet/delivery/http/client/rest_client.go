package http

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type RESTClient struct {
	apiUrl string
	client resty.Client
}

// NewRestClient will create new an RESTClient for given api url
func NewRESTClient(url string, maxRetry int, secwait int, maxwaitsec int) *RESTClient {
	c := *resty.New()
	c.SetRetryCount(maxRetry).
		SetRetryWaitTime(time.Duration(secwait) * time.Second).
		SetRetryMaxWaitTime(20 * time.Second)

	return &RESTClient{
		apiUrl: url,
		client: c,
	}
}

func (rc *RESTClient) Get(route string, automarshal interface{}) (result *resty.Response, err error) {
	fmt.Println(route)
	return rc.client.R().
		SetResult(automarshal).
		ForceContentType("application/json").
		Get(rc.apiUrl + route)
}
