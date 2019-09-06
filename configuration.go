package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Hosts []string `json:"hosts"`
}

func configuration(jsonFile string) (configuration Configuration, err error) {
	fp, err := os.Open(jsonFile)
	if err != nil {
		return
	}
	decoder := json.NewDecoder(fp)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&configuration)
	return
}
