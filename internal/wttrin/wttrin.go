package wttrin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const endpoint = "https://wttr.in/"

//GetWeatherImage gets the weather at the location
func GetWeatherImage(location string, inFarenheight bool) (string, error) {
	uri, _ := url.Parse(endpoint)
	uri.Path = location + "_0tqp.png"
	q := uri.Query()
	if inFarenheight {
		q.Set("u", "")
	} else {
		q.Set("m", "")
	}
	uri.RawQuery = q.Encode()

	resp, err := http.Get(uri.String())
	if err != nil {
		return "", errors.Wrapf(err, "uri: %s", uri.String())
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "uri: %s", uri.String())
	}

	err = ioutil.WriteFile(location+".png", data, 0644)
	if err != nil {
		return "", errors.Wrapf(err, "uri: %s", uri.String())
	}

	return location + ".png", nil

}

// GetWeather gets the weather text at a location
func GetWeather(location string, inFarenheight bool) (string, error) {
	uri, _ := url.Parse(endpoint)
	uri.Path = location
	q := uri.Query()
	if inFarenheight {
		q.Set("u", "")
	} else {
		q.Set("m", "")
	}
	q.Set("format", "j1")
	uri.RawQuery = q.Encode()

	resp, err := http.Get(uri.String())
	if err != nil {
		return "", errors.Wrapf(err, "uri: %s", uri.String())
	}

	defer resp.Body.Close()
	var weatherResp response
	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {
		return "", errors.Wrap(err, "json decoding")
	}

	return string(weatherResp.CurrentConditions[0].WeatherDescriptions[0].Value), nil
}
