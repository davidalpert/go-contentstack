package management

type Configuration struct {
	Host      string `yaml:"host" env:"CS_MANAGEMENT_API_HOST" env-default:"https://api.contentstack.io"`
	Key       string `yaml:"key" env:"CS_MANAGEMENT_API_KEY" env-default:""`
	Token     string `yaml:"token" env:"CS_MANAGEMENT_API_TOKEN" env-default:""`
	Debug     bool   `yaml:"debug" env:"CS_MANAGEMENT_API_DEBUG" env-default:"false"`
	UserAgent string `yaml:"user_agent" env:"CS_MANAGEMENT_API_USER_AGENT" env-default:"go-contentstack API client"`
}
