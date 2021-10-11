package ainaver

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const DefaultRestUrl string = "https://naveropenapi.apigw.ntruss.com"

type Client struct {
	HttpClient *http.Client
	Auth       AuthStruct
	Host       string
	Base       string
}

type AuthStruct struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

func NewClient(accessKey string, secretKey string) *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		Auth: AuthStruct{
			AccessKey: accessKey,
			SecretKey: secretKey,
		},
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("X-NCP-APIGW-API-KEY-ID", c.Auth.AccessKey)
	req.Header.Set("X-NCP-APIGW-API-KEY", c.Auth.SecretKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return body, err
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}
