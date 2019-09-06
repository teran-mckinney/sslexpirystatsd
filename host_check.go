package main

import (
	"crypto/tls"
	"time"
)

func hostCheck(host string) (secondsTillExpiry int, err error) {
	conn, err := tls.Dial("tcp", host, nil)
	if err != nil {
		return
	}

	// secondsTillExpiry gets initalized to 0.
	var seconds int
	for _, chain := range conn.ConnectionState().VerifiedChains {
		for _, certificate := range chain {
			seconds = int(certificate.NotAfter.Sub(time.Now()).Seconds())
			if secondsTillExpiry == 0 || secondsTillExpiry > seconds {
				secondsTillExpiry = seconds
			}
		}
	}
	return
}
