package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config is the type that maps the items
type Config map[string]interface{}

type decoder interface {
	Decode(interface{}) error
}

func getDecoder(r io.Reader, ext string) decoder {
	switch ext {
	case "json":
		return json.NewDecoder(r)
	case "yml":
		return yaml.NewDecoder(r)
	}

	return nil
}

// Load reads the file and returns the config instance
func Load(file string, defaults map[string]interface{}) (Config, error) {
	cfg := make(Config, 0)
	for k, v := range defaults {
		cfg[k] = v
	}

	var ext string
	if dot := strings.LastIndex(file, "."); dot != -1 {
		ext = file[dot+1:]
	}

	if len(ext) == 0 {
		return cfg, errors.New("Invalid filename, only .json and .yml are supported")
	}

	f, err := os.Open(file)
	if err != nil {
		return cfg, err
	}

	r := bufio.NewReader(f)

	if decoder := getDecoder(r, ext); decoder != nil {
		err = decoder.Decode(&cfg)
	}

	return cfg, err
}
