package sncf

import (
	"encoding/json"
	"strings"

	"github.com/parnurzeal/gorequest"
)

type City string

type CityListResponse struct {
	Cities []City `json:"CITIES"`
}

func GetCityList() ([]City, error) {
	url := "http://www.sncf.com/theme/js/cityList.js"
	request := gorequest.New()
	var result CityListResponse

	_, body, errs := request.Get(url).End()
	if len(errs) > 0 {
		// FIXME: concat errors
		return nil, errs[0]
	}

	body = strings.TrimSpace(body)
	body = body[strings.Index(body, "{"):]
	body = body[:len(body)-1]

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, err
	}

	return result.Cities, nil
}
