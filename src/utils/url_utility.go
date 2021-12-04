package utils

import (
	"errors"
	"net/http"
)

func UrlValidator(url string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode > 299 || resp.StatusCode < 200 {
		return errors.New(url + "was not valid")
	}
	return nil
}
