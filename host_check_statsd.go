package main

import (
	"log"

	"gopkg.in/alexcesaro/statsd.v2"
)

func hostCheckStatsd(host string) (secondsTillExpiry int, err error) {
	s, err := statsd.New(statsd.Prefix("sslexpiry.hostCheck"))
	if err != nil {
		// Since we are calling for the sake of statsd, this should fail.
		return
	}
	defer s.Close()
	defer s.NewTiming().Send("timetocheck." + host)
	secondsTillExpiry, err = hostCheck(host)
	if err != nil {
		return
	}
	s.Gauge("expiry."+host, secondsTillExpiry)
	return
}

func hostCheckStatsdConfiguration(conf Configuration) (err error) {
	for _, host := range conf.Hosts {
		seconds, hostCheckError := hostCheckStatsd(host)
		if hostCheckError == nil {
			log.Printf("%s: %d seconds remaining.", host, seconds)
		} else {
			err = hostCheckError
			log.Printf("Got error on %s: %s", host, hostCheckError.Error())
		}
	}
	return
}
