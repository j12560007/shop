package main

import (
	"os"
	"path/filepath"
	"shop/config"
	"shop/routes"
)

var MainPath = filepath.Dir(os.Args[0])

func main() {
	config.Init(MainPath)

	// result := dao.GetUserInfoById(1)
	// fmt.Printf("result=%+v\n", result)

	router := routes.SetRouteer()
	router.Run(config.ServerConf.Port)
}
