package libconf

import (
	pkgerr "example.com/m/v2/pkg/error"
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	FileCSVPaths []string `mapstructure:"file_csv_paths"`
	TimeToThink  int      `mapstructure:"time_to_think"`
}

var (
	onceConfig sync.Once
	config     *Config
)

func Get() *Config {
	onceConfig.Do(func() {
		viper.AddConfigPath("exercise1/configs")
		viper.SetConfigName("config")

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err.Error())
			return
		}

		config = newConfig()
	})

	return config
}

func newConfig() *Config {
	cfg := new(Config)

	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("new config error: %v", err.Error())
		return nil
	}

	return cfg
}

func (c *Config) Print() error {
	if c == nil {
		return pkgerr.ErrConfigIsNil
	}

	fmt.Printf("files example:\n")

	for _, f := range c.FileCSVPaths {
		fmt.Printf("%v\n", f)
	}
	fmt.Println()

	return nil
}
