package main

import "net/url"
import "fmt"
import "net"

func main() {
	p := fmt.Println

	s := "postgres://user:pwd@host.com:5432/path?k=v#f=123"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	p(u.Scheme)
	p(u.User)
	p(u.User.Username())
	pw, e := u.User.Password()
	p(pw)
	p(e)

	p(u.Host)

	host, port, b := net.SplitHostPort(u.Host)
	p(host)
	p(port)
	p(b)

	p(u.Path)
	p(u.Fragment)
	p(u.RawQuery)
	p(u.RawPath)

	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"])
	p(m["k"][0])
	p(m["k2"])
}
