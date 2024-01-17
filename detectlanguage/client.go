package detectlanguage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("TIMEOUT CAN'T BE ZERO")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		}}, nil
}

func (c Client) GetLanguages() ([]languageResponce, error) {
	resp, err := c.client.Get("https://ws.detectlanguage.com/0.2/languages")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r []languageResponce
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c Client) GetLanguage(name string) (languageResponce, error) {
	resp, err := c.client.Get("https://ws.detectlanguage.com/0.2/languages")
	if err != nil {
		return languageResponce{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return languageResponce{}, err
	}

	var r []languageResponce
	if err := json.Unmarshal(body, &r); err != nil {
		return languageResponce{}, err
	}

	for _, lang := range r {
		if name == lang.Name {
			return lang, nil
		}
	}

	return languageResponce{}, errors.New("language not found")
}

func (c Client) AccountStatus(apiKey string) (accountStatusResponce, error) {
	apiURL := "https://ws.detectlanguage.com/0.2/user/status"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return accountStatusResponce{}, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return accountStatusResponce{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return accountStatusResponce{}, err
	}

	var r accountStatusResponce
	if err := json.Unmarshal(body, &r); err != nil {
		return accountStatusResponce{}, err
	}

	return r, err
}

func (c Client) SingleDetect(apiKey, word string) (detectResponce, error) {
	apiURL := "https://ws.detectlanguage.com/0.2/detect"

	var wordForReader string
	wordForReader += fmt.Sprintf("q=%s", word)

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(wordForReader))
	if err != nil {
		return detectResponce{}, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return detectResponce{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return detectResponce{}, err
	}

	var r detectResponce
	if err := json.Unmarshal(body, &r); err != nil {
		return detectResponce{}, err
	}

	return r, nil
}
