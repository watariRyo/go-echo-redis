package conf

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Db        Db        `yaml:"db"`
	ApiServer ApiServer `yaml:"apiServer"`
	Redis     Redis     `yaml:"redis"`
	Jwt       Jwt       `yaml:"jwt"`
}

type Db struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type ApiServer struct {
	Port        string `yaml:"port"`
	CorsOrigins string `yaml:"corsOrigins"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Key      string `yaml:"key"`
}

type Jwt struct {
	Key string `yaml:"key"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf/")

	// 環境変数が指定されていればそちらを優先
	viper.AutomaticEnv()
	// データ構造切り替え
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("設定ファイル読み込みエラー： %s \n", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error: %s \n", err)
	}

	return &cfg, nil
}
