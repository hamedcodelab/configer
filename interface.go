package configer

type Register interface {
	SetConfigType(typ string)
	LoadConfig(c any) error
	Unmarshal(rawVal interface{}, opts ...interface{}) error
	SubConfig(SubKeyConfig string) Register
}
