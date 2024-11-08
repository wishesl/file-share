package global

import (
	"file-share/pkg/gormx"
	"fmt"
	"log"
	"net"

	"strings"

	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

const AppName = "share-file"

var (
	DB  *gorm.DB
	Cfg = Config{"", 8061, 1}
)

func init() {
	db, err := gormx.New(gormx.Config{Type: "sqlite3", DSN: "share"})
	if err != nil {
		log.Panic(err)
	}
	DB = db
}

func init() {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		InitIp()
		return
	}
	_ = cfg.MapTo(&Cfg)
	if Cfg.Ip == "" {
		InitIp()
	}
}

func InitIp() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			ipStr := ip.String()
			fmt.Println("IP:", ipStr)
			if strings.HasPrefix(ipStr, "192.") {
				fmt.Println("Found 192 IP:", ipStr)
				Cfg.Ip = ipStr
				return
			}
		}
	}
}
