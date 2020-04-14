package config

import (
	T "mirror/tool"

	"gopkg.in/yaml.v2"
)

type Replaced struct {
	Old string `yaml:"old"`
	New string `yaml:"new"`
}

type Yaml struct {
	Host struct {
		Self  string `yaml:"self"`
		Proxy string `yaml:"proxy"`
	}
	ReplacedURLs []Replaced `yaml:"replaced_urls"`
	EnableSSL    bool       `yaml:"enable_ssl"`
	HandleCookie bool       `yaml:"handle_cookie"`
}

var data = `
enable_ssl: True
handle_cookie: True
host:
  self: lovexl.now.sh
  proxy: s2-us2.startpage.com
replaced_urls:
  - old: www.startpage.com
    new: s2-us2.startpage.com
  - old: s2-us2.startpage.com
    new: lovexl.now.sh
`
var Config *Yaml
var Protocal string

func LoadConfig() {
	Config = new(Yaml)
	err := yaml.Unmarshal([]byte(data), Config)
	T.CheckErr(err)
	if Config.EnableSSL {
		Protocal = "https://"
	} else {
		Protocal = "http://"
	}
}
