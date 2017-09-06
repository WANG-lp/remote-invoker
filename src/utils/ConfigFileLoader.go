package utils

import (
	"github.com/go-ini/ini"
	"log"
)

func ConfigFileLoader(filename string, cmdName string) *ini.Section  {
	cfg, err := ini.InsensitiveLoad(filename)

	if err != nil {
		log.Println("config file load error...")
	}

	section, err := cfg.GetSection(cmdName)

	return section
}
