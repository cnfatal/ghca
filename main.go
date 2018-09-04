package main

import (
	"flag"
	"github.com/fatalc/ghca/encrypt"
	"github.com/fatalc/ghca/server"
	"log"
	"net/http"
	"os"
)

var (
	username = flag.String("username", "", "参与算号的账户")
	password = flag.String("password", "", "参与算号的账户密码")
	isServer = flag.Bool("server", false, "是否以web服务器方式启动")
)

func init() {
	flag.Parse()
}
func main() {
	if *isServer {
		http.Handle("/", &server.GhcaHandler{})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
	if *isServer == false && *username == "" || *password == "" {
		println("当不以服务器形式启动时, {username} 与 {password} 必需被设置")
		flag.Usage()
		os.Exit(1)
	} else {
		println(encrypt.GhcaEncode(*username, *password))
	}
}
