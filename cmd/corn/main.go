package main

import (
	"fmt"
	"github.com/spf13/viper"
	"jasen.project/pkg"
)

func main() {
	pkg.InitConfig()
	port := viper.GetString("server.port")
	fmt.Println("read port:" + port)
}
