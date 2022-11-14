package loader

type Config struct {
	Path   string
	Region string
	Label  string
}

func NewConfig() *Config {
	return &Config{
		Path:   "",
		Region: "",
		Label:  "",
	}
}
