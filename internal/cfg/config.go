package cfg

import (
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path/filepath"
	"strings"
)

var File string
var Dir string

type ManagementApiConfig struct {
	Host  string `yaml:"host" env:"CS_MANAGEMENT_API_HOST" env-default:"https://api.contentstack.io"`
	Key   string `yaml:"key" env:"CS_MANAGEMENT_API_KEY" env-default:""`
	Token string `yaml:"token" env:"CS_MANAGEMENT_API_TOKEN" env-default:""`
}

type DeliveryApiConfig struct {
	Host string `yaml:"delivery_api_host" env:"CS_DELIVERY_API_HOST" env-default:"https://cdn.contentstack.io"`
}

type ContentStackConfig struct {
	DeliveryApi   DeliveryApiConfig   `yaml:"delivery_api"`
	RegionCode    string              `yaml:"region_code" env:"CS_REGION" env-default:"NA"`
	ManagementApi ManagementApiConfig `yaml:"management_api"`
}

type Config struct {
	ContentStack ContentStackConfig `yaml:"content_stack"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) String() string {
	s, _, _ := printers.NewPrinterOptions().WithDefaultOutput("yaml").FormatOutput(c)
	return s
}

func (c *Config) Validate() error {
	if c == nil {
		return fmt.Errorf("config is nil")
	}
	errors := make([]string, 0)

	if c.ContentStack.RegionCode == "" {
		errors = append(errors, "ContentStack RegionCode is required")
	}

	if c.ContentStack.DeliveryApi.Host == "" {
		errors = append(errors, "ContentStack DeliveryApiHost is required")
	}

	if c.ContentStack.ManagementApi.Host == "" {
		errors = append(errors, "ContentStack ManagementApiHost is required")
	}

	if c.ContentStack.ManagementApi.Key == "" {
		errors = append(errors, "ContentStack ManagementApiKey is required")
	}

	if c.ContentStack.ManagementApi.Token == "" {
		errors = append(errors, "ContentStack ManagementApiToken is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ", "))
	}

	return nil
}

func ReadMergedInto(c *Config) error {
	if err := ensureFileExists(File, c.String()); err != nil {
		return err
	}
	return cleanenv.ReadConfig(File, c)
}

func (c *Config) Write() error {
	return c.WriteToFile(File)
}

func (c *Config) WriteToFile(file string) error {
	configDir, _ := filepath.Split(file)

	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return fmt.Errorf("make %s: %v", configDir, err)
	}

	y, _, err := printers.NewPrinterOptions().WithDefaultOutput("yaml").FormatOutput(*c)
	if err != nil {
		return fmt.Errorf("formatting %#v: %v", c, err)
	}

	if err = os.WriteFile(file, []byte(y), os.ModePerm); err != nil {
		return fmt.Errorf("write %s: %v", file, err)
	}

	return nil
}

func ensureFileExists(path string, defaultContent string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return fmt.Errorf("making %s; %v", filepath.Dir(path), err)
	}

	return os.WriteFile(path, []byte(defaultContent), os.ModePerm)
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	Dir = filepath.Join(home, ".cs")
	File = filepath.Join(Dir, "config.yaml")
}
