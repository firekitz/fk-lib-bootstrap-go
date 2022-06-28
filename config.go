package fkBootstrap

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"os"
	"path"
	"reflect"
	"strings"
)

type Payload struct {
	Parse       string `json:"parse,omitempty"`
	Username    string `json:"username,omitempty"`
	IconUrl     string `json:"icon_url,omitempty"`
	IconEmoji   string `json:"icon_emoji,omitempty"`
	Channel     string `json:"channel,omitempty"`
	Text        string `json:"text,omitempty"`
	LinkNames   string `json:"link_names,omitempty"`
	UnfurlLinks bool   `json:"unfurl_links,omitempty"`
	UnfurlMedia bool   `json:"unfurl_media,omitempty"`
	Markdown    bool   `json:"mrkdwn,omitempty"`
}

var SERVICE_NAME = "zigot"

func GetWorkingDir() string {
	dir, _ := os.Getwd()
	return dir
}

func readEnv(cfg interface{}) error {
	if err := envconfig.Process("", cfg); err != nil {
		return fmt.Errorf("fatal error config struct: %v", err)
	}
	return nil
}

func localConfig(configFile string) error {
	p := GetWorkingDir()
	configFile = path.Join(p, configFile)
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("fatal error config file: %v", err)
	}
	return nil
}

func LoadEnv(cfg interface{}) error {
	env := os.Getenv("GO_ENV")
	switch env {
	case "dev", "stage", "prod", "meta":
		if err := readEnv(cfg); err != nil {
			return err
		}
	default:
		file := fileMatch(env)
		if err := localConfig(file); err != nil {
			return err
		}
		if err := viper.Unmarshal(cfg); err != nil {
			return fmt.Errorf("unmarshal error: %v", err)
		}
	}

	defer func() {
		var f []string
		f = append(f, "SERVICE_NAME")
		fields, _ := GetFieldFromStruct(cfg, f)
		SERVICE_NAME = fields["SERVICE_NAME"]
	}()

	return nil
}

func fileMatch(env string) string {
	if env == "" {
		env = "app.env"
	} else {
		env = "app-local.env"
	}
	return env
}

func ChangeWorkingDirectory(path string) error {
	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("error occurred while changing the working directory: %v", err)
	}
	return nil
}

func GetFieldFromStruct(cfg interface{}, field []string) (result map[string]string, err error) {
	result = map[string]string{}
	err = nil
	s := reflect.ValueOf(cfg)
	if s.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("error args != pointer type")
	}
	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return nil, fmt.Errorf("error args != struct type")
	}
	typeOfSpec := s.Type()

	for _, v := range field {
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fType := typeOfSpec.Field(i)
			if v == fType.Name {
				result[v] = f.String()
				break
			}
		}
	}

	return
}

func IsLocal(cfg interface{}) bool {
	var f []string
	f = append(f, "ENV")
	fields, _ := GetFieldFromStruct(cfg, f)
	env := fields["ENV"]
	return strings.Contains(env, "local")
}
