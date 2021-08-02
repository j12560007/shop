package config

import (
	"fmt"
	"path/filepath"
	"github.com/BurntSushi/toml"
)

type TServerConfig struct {
	Port string
	SqlDb string
	SqlSchema string
	SqlUser string
	SqlPassword string
}

var ServerConf TServerConfig
var ConfPath string = "config"
var ServerConfFileName = "server.toml"

func Init(mainPath string) {
	var severConfPath = filepath.Join(mainPath, ConfPath, ServerConfFileName)
	fmt.Println("path=", severConfPath)

	if _, err := toml.DecodeFile(severConfPath, &ServerConf); err != nil {
		panic(err)
	}

}