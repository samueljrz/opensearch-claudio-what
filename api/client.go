package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const 

func NewClientFromHTTP(httpClient *http.Client) *Client {
	client := &Client{http: httpClient}
	return client
}
