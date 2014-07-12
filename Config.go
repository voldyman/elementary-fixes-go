package main

import (
	"code.google.com/p/gcfg"
)

const (
	CONFIGFILE = "bot.cfg"
)

type Config struct {
	Bot struct {
		Frequency string
	}
	Twitter struct {
		ConsumerKey       string
		ConsumerSecret    string
		AccessTokenKey    string
		AccessTokenSecret string
	}
}

func getConfig() (cfg Config, err error) {
	err = gcfg.ReadFileInto(&cfg, CONFIGFILE)
	return
}
