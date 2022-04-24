package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Mongo      `json:"mongo"`
	TencentCOS `json:"cos"`
	LogConfig  `json:"log"`
}

type Mongo struct {
	Uri string `json:"uri"`
	DB  string `json:"db"`
}

type TencentCOS struct {
	Bucket     string `json:"bucket"`
	Region     string `json:"region"`
	SecretID   string `json:"secretID"`
	SecretKey  string `json:"secretKey"`
	BaseURL    string `json:"baseURL"`
	PathPrefix string `json:"pathPrefix"`
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

// Conf 全局配置变量
var Conf = new(Config)

// InitConfig 初始化配置；从指定文件加载配置文件
func InitConfig(filePath string) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, Conf)
}
