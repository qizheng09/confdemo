package main

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	log "github.com/inconshreveable/log15"
)

func main() {
	viper.SetConfigName("configdemo")
	gopath := os.Getenv("GOPATH")
	for _, p := range filepath.SplitList(gopath) {
		confpath := filepath.Join(p, "conf")
		viper.AddConfigPath(confpath)
	}

	// normal mode
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Read config file error!", "err", err.Error())
		return
	}
	log.Info("Config node:", "node", viper.GetString("node"))
	log.Warn("Config database:", "database", viper.GetString("database"))
	log.Crit("Config index:", "index", viper.GetString("index"))
	log.Info("Configs:", "node", viper.GetString("node"), "database", viper.GetString("database"), "index", viper.GetString("index"))

	// contexts mode
	ctxLog := log.New("contexts", "log contexts")
	ctxLog.Info("Config node:", "node", viper.GetString("node"))
	ctxLog.Warn("Config database:", "database", viper.GetString("database"))
	ctxLog.Crit("Config index:", "index", viper.GetString("index"))

	// specify output
	specifyLog := log.New()
	specifyLog.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))
	specifyLog.Info("Config node:", "node", viper.GetString("node"))
	specifyLog.Warn("Config database:", "database", viper.GetString("database"))
	specifyLog.Crit("Config index:", "index", viper.GetString("index"))
}
