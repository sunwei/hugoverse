package main

import (
	"fmt"
	"github.com/spf13/afero"
	"github.com/sunwei/hugoverse/pkg/hugo/config/allconfig"
	"os"
	"path/filepath"
)

func main() {
	tempDir := os.TempDir()
	fmt.Println("temp dir:", tempDir)

	configFilename := filepath.Join(tempDir, "hugo.toml")
	config := `
baseURL = "https://example.com"
defaultContentLanguage = 'en'

[module]
[[module.mounts]]
source = 'content/en'
target = 'content/en'
lang = 'en'
[[module.mounts]]
source = 'content/nn'
target = 'content/nn'
lang = 'nn'
[[module.mounts]]
source = 'content/no'
target = 'content/no'
lang = 'no'
[[module.mounts]]
source = 'content/sv'
target = 'content/sv'
lang = 'sv'
[[module.mounts]]
source = 'layouts'
target = 'layouts'

[languages]
[languages.en]
title = "English"
weight = 1
[languages.nn]
title = "Nynorsk"
weight = 2
[languages.no]
title = "Norsk"
weight = 3
[languages.sv]
title = "Svenska"
weight = 4
`
	if err := os.WriteFile(configFilename, []byte(config), 0666); err != nil {
		fmt.Println(err)
	}
	d := allconfig.ConfigSourceDescriptor{
		Fs:       afero.NewOsFs(),
		Filename: configFilename,
	}

	configs, err := allconfig.LoadConfig(d)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(configs.LoadingInfo.ConfigFiles)
	fmt.Println(configs.Languages)
	for _, mc := range configs.Modules {
		for _, mount := range mc.Mounts() {
			fmt.Println(mount.Source, mount.Target)
		}
	}
}
