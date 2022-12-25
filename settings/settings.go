package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 实例化对象：对应yaml 文件
var Conf = new(AppConfig)

// AppConfig struct
type AppConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	Pagesize     int    `mapstructure:"pagesize"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

// MySQL config struct
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// Redis config struct
type RedisConfig struct {
	Addr     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// LogConfig
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// 定义一个初始化函数
// 1. 读取config.yaml
// 2. 文件变更，自动识别（main.go web 热加载自动）
// viper
func Init() error {
	// 目的：就是读取－> Conf
	viper.SetConfigFile("./conf/config.yaml")
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		viper.Unmarshal(&Conf)
	})
	viper.WatchConfig()

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed,err:%v", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("Unmarshal failed,err:%v", err))
	}
	return err
}
