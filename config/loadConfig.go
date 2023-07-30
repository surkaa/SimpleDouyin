package config

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strings"
)

// Config 全局配置文件 TODO 解决高并发下访问的问题
var Config map[string]interface{}

func LoadConfig(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("配置文件加载失败")
	}
	ext := strings.ToLower(file[strings.LastIndex(file, ".")+1:])
	Config = make(map[string]interface{})
	switch ext {
	case "json":
		err = json.Unmarshal(data, &Config)
	case "yaml":
		err = yaml.Unmarshal(data, &Config)
	default:
		log.Fatalf("unsupported file format: %s", ext)
	}

	if err != nil {
		log.Fatalf("配置文件加载失败")
	}
}
