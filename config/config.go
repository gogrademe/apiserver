package config

import (
	"code.google.com/p/gcfg"
)

type Conf struct {
	Database struct {
		User   string
		DBName string
	}
}

func Load(cfgFile string) (*Conf, error) {

	cfg := &Conf{}

	err := gcfg.ReadFileInto(cfg, cfgFile)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
