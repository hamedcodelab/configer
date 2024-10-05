package configer

import (
	"github.com/spf13/viper"
)

type config struct {
	viper *viper.Viper
}

func New() Register {
	return &config{
		viper: viper.New(),
	}
}

func (c *config) SetConfigType(typ string) {
	c.viper.SetConfigType(typ)
}

func (c *config) SetConfigName(name string) {
	c.viper.SetConfigName(name)
}

func (c *config) SetConfigPath(path string) {
	c.viper.AddConfigPath(path)
}

func (c *config) LoadConfig(str any) error {
	err := c.viper.ReadInConfig()
	if err != nil {
		return err
	}
	return c.Unmarshal(str)
}

func (c *config) Unmarshal(rawVal interface{}, opts ...interface{}) error {
	return c.viper.Unmarshal(rawVal)
}

func (c *config) ValueOf(key string) Register {
	v := c.viper.Sub(key)
	if v == nil {
		return nil
	}
	return &config{
		viper: v,
	}
}
