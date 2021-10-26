package http

import (
	"encoding/xml"
	"io/ioutil"
	"kokuzeityo/parser"
	"log"
	"net/http"
	"time"
)

type GetCorpRequest struct {
	Id             string
	Numbers        []string
	ResType        string
	History        string
	Endpoint       string
	Version        string
	RequestTimeOut int
}

type Apis interface {
	GetCorporationLists(req *GetCorpRequest)
}

type Api struct {
}

func (api Api) GetCorporationLists(req *GetCorpRequest) (*parser.Kcorporations, error) {

	apiPath := "num"

	url := req.Endpoint + "/" + req.Version + "/" + apiPath
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("http request error: %v", err)
		return nil, err
	}

	numbers := ""

	for i, v := range req.Numbers {
		numbers = numbers + v
		if i != len(req.Numbers)-1 {
			numbers = numbers + ","
		}
	}

	params := request.URL.Query()
	params.Add("id", req.Id)
	params.Add("number", numbers)
	params.Add("type", req.ResType)
	params.Add("history", req.History)
	request.URL.RawQuery = params.Encode()

	log.Println(request.URL.String())

	client := &http.Client{
		Timeout: time.Duration(int64(req.RequestTimeOut)) * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Printf("http connection error :%v", err)
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("http response body read error :%v", err)
		return nil, err
	}

	log.Println(string(body))

	var xmlRes parser.Kcorporations

	err = xml.Unmarshal(body, &xmlRes)
	if err != nil {
		log.Printf("xml parse error :%v", err)
		return nil, err
	}

	return &xmlRes, nil

}
