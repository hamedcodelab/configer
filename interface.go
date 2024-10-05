package configer

type Register interface {
	SetConfigType(typ string)
	SetConfigName(name string)
	SetConfigPath(path string)
	LoadConfig(c any) error
	Unmarshal(rawVal interface{}, opts ...interface{}) error
	ValueOf(key string) Register
}
