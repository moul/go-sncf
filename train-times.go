package sncf

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type TrainTimesDeparture struct {
	OrigDest string `json:"origdest"`
	Num      string `json:"num"`
	Type     string `json:"type"`
	Picto    string `json:"picto"`
	Voie     string `json:"voie"`
	VoieAttr string `json:"voie_attr"`
	Heure    string `json:"heure"`
	Etat     string `json:"etat"`
	Retard   string `json:"retard"`
	Infos    string `json:"infos"`
}

type TrainTimesDepartures struct {
	Trains  []TrainTimesDeparture `json:"trains"`
	Updated string                `json:"updated"`
}

func GetTrainTimesDeparture(cityCode string) (*TrainTimesDepartures, error) {
	url := fmt.Sprintf("http://www.gares-sncf.com/fr/train-times/departure/%s/gl", cityCode)
	request := gorequest.New()
	var result TrainTimesDepartures

	_, body, errs := request.Get(url).End()
	if len(errs) > 0 {
		// FIXME: concat errors
		return nil, errs[0]
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, err
	}

	return &result, nil
}
