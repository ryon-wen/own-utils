package dao

import (
	"fmt"
	"log"

	"github.com/beego/beego/httplib"
)

var esUrl string

func InitElasticsearch(host string, port int, indexKey string) {
	esUrl = fmt.Sprintf("http://%s:%d/%s", host, port, indexKey)
	log.Println("init elasticsearch success: " + esUrl)
}

func EsSearch(sr interface{}) (string, error) {
	req, err := httplib.Get(fmt.Sprintf("%s/_search", esUrl)).JSONBody(&sr)
	if err != nil {
		return "", err
	}
	str, err := req.String()
	return str, err
}
