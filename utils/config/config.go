package config

import (
	"go.uber.org/zap"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pelletier/go-toml/v2"
)

type app = struct {
	Name        string        `toml:"name"`
	Port        string        `toml:"port"`
	PrintRoutes bool          `toml:"print-routes"`
	Prefork     bool          `toml:"prefork"`
	Production  bool          `toml:"production"`
	IdleTimeout time.Duration `toml:"idle-timeout"`
}

type db = struct {
	Driver string `toml:"driver"`
	MySQL  struct {
		DSN string `toml:"dsn"`
	}
}

type middleware = struct {
}

type Config struct {
	App        app
	DB         db
	Middleware middleware
}

func ParseConfig(name string, debug ...bool) (*Config, error) {
	var contents *Config
	var file []byte
	var err error

	if len(debug) > 0 {
		file, err = os.ReadFile(name)
	} else {
		file, err = os.ReadFile("./config/" + name + ".toml")
	}

	if err != nil {
		return &Config{}, err
	}

	err = toml.Unmarshal(file, &contents)

	return contents, err
}

func NewConfig() *Config {
	config, err := ParseConfig("example")
	if err != nil && !fiber.IsChild() {
		zap.Error(err)
		panic(err)
	}

	return config
}

// ParseAddr From https://github.com/gofiber/fiber/blob/master/helpers.go#L305.
func ParseAddr(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i != -1 {
		return raw[:i], raw[i+1:]
	}
	return raw, ""
}
