package libconf

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"

	pkgerr "github.com/Nikitastarikov/practice-on-golang/pkg/error"
)

type Config struct {
	Psql Psql `mapstructure:"psql"`
}

type Psql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

var (
	onceConfig sync.Once
	config     *Config
)

func Get() *Config {
	onceConfig.Do(func() {
		viper.AddConfigPath("exercise4/configs")
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

	fmt.Printf("Config print:\n")
	fmt.Printf("Psql:\n%v", c.Psql)

	fmt.Println()

	return nil
}
