package main

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"fmt"
)

func main() {
	viper.SetConfigName("configdemo")
	gopath := os.Getenv("GOPATH")
	for _, p := range filepath.SplitList(gopath) {
		confpath := filepath.Join(p, "conf")
		viper.AddConfigPath(confpath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		return
	}
	fmt.Println(viper.GetString("node"))
	fmt.Println(viper.GetString("database"))
	fmt.Println(viper.GetString("index"))
}
