package main

import (
	"fmt"
	"net"
	"net/url"
)

func urlParsing() {
	print := fmt.Println
	urlString := "postgres://user:pass@host.com:5432/path?k1=v1&k2=v2&k3=a&k3=b#fragment"

	_url, _error := url.Parse(urlString)
	if _error != nil {
		panic(_error)
	}

	print(_url.Scheme, _url.User, _url.User.Username()) // postgres user:pass user
	password, isSet := _url.User.Password()
	print(password, isSet) // pass true

	print(_url.Host) // host.com:5432
	host, port, _ := net.SplitHostPort(_url.Host)
	print(host, port) // host.com 5432

	print(_url.Fragment) // fragment
	print(_url.Path)     // /path

	print(_url.RawQuery) // k1=v1&k2=v2&k3=a&k3=b

	m, _ := url.ParseQuery(_url.RawQuery)
	print(m)          // map[k1:[v1] k2:[v2] k3:[a b]]
	print(m["k1"])    // [v1]
	print(m["k1"][0]) // v1
	print(m["k2"])    // [v2]
	print(m["k3"])    // [a b]
	print(m["k3"][0]) // a
	print(m["k3"][1]) // b
}
