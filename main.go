package main

import (
	"flag"
	"fmt"
)

func main() {
	c := Config{}
	c.Setup()
	flag.Parse()
	x, err := getNagiosHosts(c.fqdn, c.apikey)
	if err != nil {
		panic(err.Error())
	}
	for _, ou := range x.Hoststatus {
		fmt.Println(ou.Name, ou.Address)
	}

}
