package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// 创建一个新的Viper实例
	v := viper.New()

	// 设置需要被发现的文件名称 不带后缀
	v.SetConfigName("test")

	// 设置配置文件所在的路径（可以是相对路径或绝对路径）
	v.AddConfigPath("./")

	// 设置配置文件的类型（可选，支持多种类型，如JSON、YAML等）
	v.SetConfigType("yaml")

	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to read config file: %s\n", err)
		return
	}

	// 从配置文件中读取具体的配置项
	value := v.GetString("key")
	fmt.Println("Value:", value)
}
