package scraper

import (
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	HTTPClient *http.Client
}

func (s *Scraper) Fetch(method string, urlStr string) (*goquery.Document, error) {
	req, err := GenerateRequest(method, urlStr, nil)

	if err != nil {
		return nil, err
	}

	resp, err := s.GetClient().Do(req)

	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromResponse(resp)
}

func (s *Scraper) GetClient() *http.Client {
	if s.HTTPClient == nil {
		s.HTTPClient = &http.Client{}
	}

	return s.HTTPClient
}

func GenerateRequest(method string, urlStr string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "PuckmanLabs API Scraper")

	return req, nil
}
