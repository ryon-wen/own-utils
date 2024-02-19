package own

import (
	"fmt"
	"log"
	"strings"

	"github.com/beego/beego/httplib"
)

type ElasticSearch struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	IndexKey string `json:"index_key"`
}

var esUrl string

func InitElasticsearch(e *ElasticSearch) {
	esUrl = fmt.Sprintf("http://%s:%d/%s", e.Host, e.Port, e.IndexKey)
	log.Println("init elasticsearch success: " + esUrl)
}

func EsAdd(id int64, req interface{}) error {
	resp, err := httplib.Post(fmt.Sprintf("%s/_doc/%d", esUrl, id)).JSONBody(&req)
	if err != nil {
		return err
	}
	str, err := resp.String()
	if err != nil {
		return err
	}
	if strings.Contains(str, "error") {
		return fmt.Errorf(str)
	}
	return nil
}

func EsSearch(req interface{}) (string, error) {
	resp, err := httplib.Get(fmt.Sprintf("%s/_search", esUrl)).JSONBody(&req)
	if err != nil {
		return "", err
	}
	str, err := resp.String()
	return str, err
}

type Ik struct {
	Settings struct {
		Index struct {
			Analysis struct {
				Analyzer struct {
					Default struct {
						Type string `json:"type"`
					} `json:"default"`
				} `json:"analyzer"`
			} `json:"analysis"`
		} `json:"index"`
	} `json:"settings"`
	Mappings struct {
		Properties struct {
			Description struct {
				Type     string `json:"type"`
				Analyzer string `json:"analyzer"`
			} `json:"description"`
		} `json:"properties"`
	} `json:"mappings"`
}

const (
	IkMaxWord = "ik_max_word"
	IkSmart   = "ik_smart"
)

func EsSetIk(ikType string) error {
	if ikType != IkMaxWord && ikType != IkSmart {
		return fmt.Errorf("ikType need current value, suggest using `dao.IkMaxWord` or `dao.IkSmart`")
	}
	var ik Ik
	ik.Settings.Index.Analysis.Analyzer.Default.Type = ikType
	ik.Mappings.Properties.Description.Analyzer = "text"
	ik.Mappings.Properties.Description.Type = ikType
	resp, err := httplib.Put(esUrl).JSONBody(&ik)
	if err != nil {
		return err
	}
	str, err := resp.String()
	if err != nil {
		return err
	}
	if strings.Contains(str, "error") {
		return fmt.Errorf(str)
	}
	return nil
}
