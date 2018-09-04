package server

import (
	"encoding/json"
	"github.com/fatalc/ghca/encrypt"
	"log"
	"net/http"
)

const HELP_MSG = "电信天翼飞Young 3.09 算号工具\n使用方法: GET /?username=<username>&password=<password>"

type GhcaHandler struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Encrypt  string `json:"encrypt,omitempty"`
	Help     string `json:"help,omitempty"`
}

func (gh *GhcaHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	gh.Username = request.URL.Query().Get("username")
	gh.Password = request.URL.Query().Get("password")

	if gh.Username == "" || gh.Password == "" {
		writer.WriteHeader(http.StatusPaymentRequired)
		gh.Help = HELP_MSG
	} else {
		gh.Help = ""
		writer.WriteHeader(http.StatusOK)
		gh.Encrypt = encrypt.GhcaEncode(gh.Username, gh.Password)
	}

	result, e := json.Marshal(gh)
	if e != nil {
		log.Print(e.Error())
	}
	writer.Write(result)
}
