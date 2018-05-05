// Copyright 2018 The Hugo Authors. All rights reserved.
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

package privacy

import (
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

// TODO(bep) GDPR be a little paranoid about this, so test JSON and YAML too.
func TestDecodeConfigFromTOML(t *testing.T) {
	assert := require.New(t)

	tomlConfig := `

someOtherValue = "foo"

[privacy]
[privacy.youtube]
noCookie = true
[privacy.disqus]
skipAgree = false
`
	cfg, err := config.FromTOMLString(tomlConfig)
	assert.NoError(err)

	pc, err := DecodeConfig(cfg)
	assert.NoError(err)
	assert.NotNil(pc)
	assert.True(pc.YouTube.NoCookie)
}

func TestDecodeConfigFromTOMLCaseInsensitive(t *testing.T) {
	assert := require.New(t)

	tomlConfig := `

someOtherValue = "foo"

[Privacy]
[Privacy.YouTube]
NoCOOKIE = true
[Privacy.Disqus]
SkipAgree = false
`
	cfg, err := config.FromTOMLString(tomlConfig)
	assert.NoError(err)

	pc, err := DecodeConfig(cfg)
	assert.NoError(err)
	assert.NotNil(pc)
	assert.True(pc.YouTube.NoCookie)
	assert.False(pc.Disqus.SkipAgree)
}

func TestDecodeConfigDefault(t *testing.T) {
	assert := require.New(t)

	pc, err := DecodeConfig(viper.New())
	assert.NoError(err)
	assert.NotNil(pc)
	assert.False(pc.YouTube.NoCookie)
	assert.False(pc.Disqus.SkipAgree)
}
