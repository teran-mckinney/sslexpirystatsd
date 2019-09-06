package main

import (
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
