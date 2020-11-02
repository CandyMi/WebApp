package config

import "gopkg.in/ini.v1"

type Config struct {
	env *ini.File
}

var conf *Config = nil

func Default() *Config {
	if conf == nil {
		conf = New("env.ini")
	}
	return conf
}

func New(file string) *Config {
	env, err := ini.LoadSources(ini.LoadOptions{SkipUnrecognizableLines: true}, file)
	if err != nil {
		panic(err)
	}
	return &Config{env: env}
}

func (this *Config) IsMode(mode string) bool {
	return mode == this.env.Section(``).Key(`app_mode`).String()
}

func (this *Config) GetMode(mode string) string {
	return this.env.Section(``).Key(`app_mode`).String()
}

func (this *Config) GetKey(class string, key string) string {
	return this.env.Section(class).Key(key).String()
}
