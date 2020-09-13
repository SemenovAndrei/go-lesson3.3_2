package qr

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	apiMethod = "create-qr-code"
)

type Service struct {
	url     string
	version string
	con     context.Context
	client  *http.Client
}

func NewService(
	url string,
	version string,
	con context.Context,
	client *http.Client,
) *Service {
	return &Service{
		url: url, version: version, con: con, client: client,
	}
}

func (s *Service) Encode(line string) (data []byte, err error) {
	values := make(url.Values)
	reqURL := fmt.Sprintf("%s/%s/%s", s.url, s.version, apiMethod)
	values.Set("data", line)
	values.Set("size", "500x500")

	req, err := http.NewRequestWithContext(
		s.con,
		http.MethodGet,
		fmt.Sprintf("%s?%s", reqURL, values.Encode()),
		nil,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return body, nil
}