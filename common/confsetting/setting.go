package confsetting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	JwtKey   string
	PageSize int

	ServerPort   string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
)

func init() {
	var err error
	Cfg, err = ini.Load("./conf/conf.ini")
	if err != nil {
		log.Fatalf("[confsetting-setting] init load err : %v\n", err)
	}

	appLoad()
	serverLoad()
}

func appLoad() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("[confsetting-setting] app load err : %v\n", err)
	}

	JwtKey = sec.Key("JWT_KEY").MustString("1266wywyh")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func serverLoad() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("[confsetting-setting] server load err : %v\n", err)
	}

	ServerPort = sec.Key("PORT").MustString("8090")
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
