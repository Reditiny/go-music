package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

// 系统配置，对应yml
// viper 内置了 mapstructure 用于配置文件到结构体的映射
// yml文件用"-"区分单词, 转为驼峰方便

// Conf 全局配置变量
var Conf = new(config)

type config struct {
	Mysql *MysqlConfig `mapstructure:"mysql" json:"mysql"`
	Redis *RedisConfig `mapstructure:"redis" json:"redis"`
}

// InitConfig 设置读取配置信息
func InitConfig() {
	// 获取当前工作目录
	workDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("读取应用目录失败:%s \n", err))
	}
	// 初始化 viper 单例模式 指定配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/")
	// 读取配置信息
	err = viper.ReadInConfig()
	// 保存配置信息到配置类
	saveConf(Conf)
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败:%s \n", err))
	}
	// 热更新配置
	viper.WatchConfig()
	// 注册热更新回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		saveConf(Conf)
	})

}

func saveConf(Conf *config) {
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("初始化配置文件失败:%s \n", err))
	}
	fmt.Println("保存配置成功")
}

type MysqlConfig struct {
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	Database    string `mapstructure:"database" json:"database"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Query       string `mapstructure:"query" json:"query"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode"`
	TablePrefix string `mapstructure:"table-prefix" json:"tablePrefix"`
	Charset     string `mapstructure:"charset" json:"charset"`
	Collation   string `mapstructure:"collation" json:"collation"`
}

type RedisConfig struct {
	Password string `mapstructure:"password" json:"password"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
}
