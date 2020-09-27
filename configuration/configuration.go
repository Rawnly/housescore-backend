package configuration

import (
	"fmt"
	"github.com/schigh/str"
	"os"
	"time"
)

type Configuration struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Name     string `json:"database"`
		Password string `json:"password"`
	} `json:"database"`
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
	JWT struct {
		Secret     string        `json:"secret"`
		Expiration time.Duration `json:"expiration"`
	} `json:"jwt"`
}

func (c *Configuration) DatabaseConnectionString() string {
	d := c.Database
	return fmt.Sprintf("%s:%s@/%s?parseTime=true", d.User, d.Password, d.Name)
}

var Config = Configuration{}

func Init() {
	if os.Getenv("LOG_LEVEL") != "" {
		Config.Log.Level = os.Getenv("LOG_LEVEL")
	}

	if os.Getenv("DB_USER") != "" {
		Config.Database.User = os.Getenv("DB_USER")
	}
	if os.Getenv("DB_PASSWORD") != "" {
		Config.Database.Password = os.Getenv("DB_PASSWORD")
	}
	if os.Getenv("DB_HOST") != "" {
		Config.Database.Host = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_PORT") != "" {
		Config.Database.Port = os.Getenv("DB_PORT")
	}
	if os.Getenv("DB_NAME") != "" {
		Config.Database.Name = os.Getenv("DB_NAME")
	}

	if os.Getenv("JWT_SECRET") != "" {
		Config.JWT.Secret = os.Getenv("JWT_SECRET")
	}

	if os.Getenv("JWT_EXPIRATION") != "" {
		Config.JWT.Secret = os.Getenv("JWT_EXPIRATION")
	}
}

func Default() {
	Config.Log.Level = "debug"

	// Database Configuration
	Config.Database.User = "root"
	Config.Database.Password = "toor"
	Config.Database.Host = "localhost"
	Config.Database.Port = "3306"
	Config.Database.Name = "housescore"

	// JWT
	Config.JWT.Secret = str.SHA256("federico")
	Config.JWT.Expiration = time.Hour * 24 // 1 day
}
