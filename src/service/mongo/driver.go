package mongo

import "github.com/astaxie/beego"

var maxPool int

func init() {
	var err error
	maxPool, err = beego.AppConfig.Int("mongodbMaxPool")
	if err != nil {
		// todo: panic!!!
		// panic(err)
		println(err)
	}
	// init method to start db
	checkAndInitServiceConnection()
}

func checkAndInitServiceConnection() {
	if service.baseSession == nil {
		service.URL = beego.AppConfig.String("mongodburl")
		err := service.New()
		if err != nil {
			// todo: panic!!!
			// panic(err)
			println(err)
		}
	}
}
