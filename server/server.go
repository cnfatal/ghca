package server

import (
	"github.com/fatalc/ghca/encrypt"
	"net/http"
)

const HelpMsg = `电信天翼飞Young 3.09 算号工具
使用方法: GET /?username=<username>&password=<password>
`

type GhcaHandler struct {
}

func (*GhcaHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	username := request.URL.Query().Get("username")
	password := request.URL.Query().Get("password")
	if username == "" || password == "" {
		writer.Write([]byte(HelpMsg))
	} else {
		writer.Write([]byte(encrypt.GhcaEncode(username, password)))
	}
}
