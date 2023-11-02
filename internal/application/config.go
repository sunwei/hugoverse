package application

import (
	"github.com/sunwei/hugoverse/internal/domain/config/entity"
	"github.com/sunwei/hugoverse/pkg/hugo/config/allconfig"
	"path"
)

func AllConfigurationInformation(projPath string) (*allconfig.Configs, error) {
	c := entity.Config{Path: path.Join(projPath, "config.toml")}

	return c.All()
}
