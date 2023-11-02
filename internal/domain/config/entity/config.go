package entity

import (
	"github.com/spf13/afero"
	"github.com/sunwei/hugoverse/pkg/hugo/config/allconfig"
)

type Config struct {
	Path string
}

func (c *Config) All() (*allconfig.Configs, error) {
	d := allconfig.ConfigSourceDescriptor{
		Fs:       afero.NewOsFs(),
		Filename: c.Path,
	}

	return allconfig.LoadConfig(d)
}
