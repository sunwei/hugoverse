// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package htesting

import (
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/sunwei/hugoverse/pkg/hugo/config"
	"github.com/sunwei/hugoverse/pkg/hugo/config/testconfig"
	"github.com/sunwei/hugoverse/pkg/hugo/helpers"
	"github.com/sunwei/hugoverse/pkg/hugo/hugofs"
	"github.com/sunwei/hugoverse/pkg/hugo/resources"
)

func NewTestResourceSpec() (*resources.Spec, error) {
	cfg := config.New()

	imagingCfg := map[string]any{
		"resampleFilter": "linear",
		"quality":        68,
		"anchor":         "left",
	}

	cfg.Set("imaging", imagingCfg)
	afs := afero.NewMemMapFs()

	conf := testconfig.GetTestConfig(afs, cfg)
	fs := hugofs.NewFrom(hugofs.NewBaseFileDecorator(afs), conf.BaseConfig())
	s, err := helpers.NewPathSpec(fs, conf, nil)
	if err != nil {
		return nil, err
	}

	spec, err := resources.NewSpec(s, nil, nil, nil, nil, nil, nil)
	return spec, err
}

func NewResourceTransformer(filename, content string) (resources.ResourceTransformer, error) {
	spec, err := NewTestResourceSpec()
	if err != nil {
		return nil, err
	}
	return NewResourceTransformerForSpec(spec, filename, content)
}

func NewResourceTransformerForSpec(spec *resources.Spec, filename, content string) (resources.ResourceTransformer, error) {
	filename = filepath.FromSlash(filename)

	fs := spec.Fs.Source
	if err := afero.WriteFile(fs, filename, []byte(content), 0777); err != nil {
		return nil, err
	}

	r, err := spec.New(resources.ResourceSourceDescriptor{Fs: fs, SourceFilename: filename})
	if err != nil {
		return nil, err
	}

	return r.(resources.ResourceTransformer), nil
}
