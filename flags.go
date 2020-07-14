package main

import "flag"

type Config struct {
	fqdn   string
	apikey string
}

func (c *Config) Setup() {
	flag.StringVar(&c.fqdn, "h", "nagiosxi.site.local",
		"Name of the site, fqdn preferred")
	flag.StringVar(&c.apikey, "a", "xxxxx", "API key for Nagios XI")
}
