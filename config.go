package configer

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
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

func (c *config) LoadConfig(f any) error {
	if f == nil {
		return fmt.Errorf("f is nil,must be io.reader")
	}

	reader, ok := f.(io.Reader)
	if !ok {
		return fmt.Errorf("f is not io.reader")
	}

	return c.viper.ReadConfig(reader)
}

func (c *config) Unmarshal(rawVal interface{}, opts ...interface{}) error {
	return c.viper.Unmarshal(rawVal)
}

func (c *config) SubConfig(SubKeyConfig string) Register {
	v := c.viper.Sub(SubKeyConfig)
	if v == nil {
		return nil
	}
	return &config{
		viper: v,
	}
}
