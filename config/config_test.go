package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadFileNotExist(t *testing.T) {

	assert := require.New(t)

	cfg, err := Load("/tmp/non_existent_file")
	assert.Nil(cfg)
	assert.NotNil(err)
}

func TestLoadEmptyFile(t *testing.T) {
	assert := require.New(t)

	cfg, err := Load("")

	assert.Nil(cfg)
	assert.NotNil(err)
}

func TestLoadInvalidToml(t *testing.T) {
	assert := require.New(t)

	// create a temp file and write invalid toml into it
	f, err := ioutil.TempFile("", t.Name())
	assert.Nil(err)
	defer os.Remove(f.Name())
	defer f.Close()

	_, err = f.WriteString("invalidTOML")
	assert.Nil(err)
	_ = f.Sync()

	cfg, err := Load(f.Name())
	assert.NotNil(err)
	assert.Nil(cfg)
}

func TestLoadValidToml(t *testing.T) {

	assert := require.New(t)

	cfg, err := Load("../config.toml")
	assert.Nil(err)
	assert.NotNil(cfg)
}

func TestLoadValidTomlInvalidConfig(t *testing.T) {

	assert := require.New(t)

	// create a temp file and write valid toml but invalid config
	f, err := ioutil.TempFile("", t.Name())
	assert.Nil(err)
	defer os.Remove(f.Name())
	defer f.Close()

	_, err = f.WriteString("a = 3\n")
	assert.Nil(err)
	_ = f.Sync()

	cfg, err := Load(f.Name())
	assert.NotNil(err)
	assert.Nil(cfg)
}
