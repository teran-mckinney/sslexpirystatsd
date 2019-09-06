package main

import (
	"log"
	"testing"
)

func TestHostCheck(t *testing.T) {
	seconds, err := hostCheck("go-beyond.org:443")
	if err != nil {
		t.Errorf("Got error when we should not have: %s", err.Error())
	} else {
		log.Printf("Seconds left: %d", seconds)
	}
}
