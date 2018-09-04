package main

import (
	"flag"
	"github.com/fatalc/ghca/encrypt"
	"os"
)

var (
	username = flag.String("username", "", "参与算号的账户")
	password = flag.String("password", "", "参与算号的账户密码")
)

func init() {
	flag.Parse()
}
func main() {
	if *username == "" || *password == "" {
		println("username 与 password 必需被设置")
		flag.Usage()
		os.Exit(1)
	}
	println(encrypt.GhcaEncode(*username, *password))
}
