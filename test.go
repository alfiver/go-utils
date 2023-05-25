package main

import (
	"fmt"
	"time"

	"github.com/alfiver/go-utils/cache"
	"github.com/alfiver/go-utils/db"
	htp "github.com/alfiver/go-utils/http"
	"github.com/alfiver/go-utils/ilog"
	"github.com/alfiver/go-utils/utils"
)

func main() {
	// just for compile
	utils.RandomString(8)
	ilog.Init("logs")
	db.Init("db/sqlite.db")
	fmt.Println(utils.RandomString(16))
	fmt.Println(utils.NewSecret())
	code, _ := utils.Google2FACode("VXCnU8oaCTGIp2dn")
	fmt.Println(code)
	htp.Get("https://www.baidu.com", time.Duration(1200), nil)
	cache.New(5*time.Minute, 10*time.Minute)
}
