# sslexpirystatsd

## Install

 * `go get go-beyond.org/code/sslexpirystatsd`

## Usage

### Check and give seconds remaining.

 * `sslexpirystatsd check duckduckgo.com:443`

### Check and give seconds remaining. Send to statsd. Fail if we can't connect to statsd.

 * `sslexpirystatsd checkstatsd duckduckgo.com:443`

## License

[Public domain / Unlicense](/LICENSE.txt)
