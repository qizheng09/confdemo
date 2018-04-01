package main

import (
	"configdemo/config"
	"fmt"
)

func main() {
	var conf *config.Conf
	conf = config.GetConfFromFile()
	fmt.Println(conf)
}