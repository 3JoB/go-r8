package config

import (
	"fmt"
	"runtime"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"

	r "github.com/3JoB/ulib/runtime"
)

// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
var (
	k = koanf.New(".")
)

func init() {
	var c string
	// Load YAML config and merge into the previously loaded config (because we can).
	if runtime.GOOS == "windows" {
		dir, err := r.UserConfigDir()
		if err != nil {
			panic(err)
		}
		c = fmt.Sprintf("%v/r8/config.yml", dir)
	} else {
		c = "/etc/r8/config.yml"
	}
	if err := Load(c); err != nil {
		panic(err)
	}
}

func SqlitePath() string {
	if runtime.GOOS == "windows" {
		dir, err := r.UserConfigDir()
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("%v/r8/r8db.sqlite3", dir)
	}
	return "/etc/r8/r8db.sqlite3"
}

func Load(config string) error {
	return k.Load(file.Provider(config), yaml.Parser())
}

func F() *koanf.Koanf {
	return k
}
