package config

import (
	"log"
	"os"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

type config struct {
	Env  string `mapstructure:"ENV"`
	Port string `mapstructure:"PORT"`
	Dsn  string `mapstructure:"DSN"`
}

var cfg *config

func InitConfig(filenames ...string) {
	vi := viper.NewWithOptions()

	if len(filenames) > 0 {
		if _, err := os.Stat(filenames[0]); err != nil {
			log.Fatalf("config file not found: %v", err)
		}

		vi.SetConfigFile(filenames[0])
	} else {
		if _, err := os.Stat(".env"); err == nil {
			vi.SetConfigFile(".env")
		}
	}

	setDefaultValues(vi)
	vi.AutomaticEnv()

	if vi.ConfigFileUsed() != "" {
		if err := vi.ReadInConfig(); err != nil {
			log.Fatalf("failed to read config: %v", err)
		}
	} else {
		var out map[string]any
		if err := mapstructure.Decode(config{}, &out); err != nil {
			log.Fatalf("failed to decode config: %v", err)
		}

		for key := range out {
			vi.MustBindEnv(key)
		}
	}

	if err := vi.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to parse config into struct: %v", err)
	}
}

func ReadConfig(filenames ...string) *config {
	if cfg == nil {
		InitConfig(filenames...)
	}

	return cfg
}

func setDefaultValues(vi *viper.Viper) {
	vi.SetDefault("ENV", "development")
	vi.SetDefault("PORT", "8080")
}
