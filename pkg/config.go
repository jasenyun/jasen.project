package pkg

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Server struct {
		Name string
	}
}

func InitConfig() {
	workDir, _ := os.Getwd()
	fmt.Println("当前位置:" + workDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置错误")
	}
}
