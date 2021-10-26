package http

import "testing"

func TestGetCorporationLists_Success(t *testing.T) {

	var req GetCorpRequest
	req.Id = "KfqqvEJrbipy5"
	req.Numbers = []string{"9010601021385"}
	req.Version = "4"
	req.ResType = "12"
	req.RequestTimeOut = 10
	req.Endpoint = "https://api.houjin-bangou.nta.go.jp"
	req.History = "0"

	var api Api

	_, err := api.GetCorporationLists(&req)
	if err != nil {
		t.Errorf("Failed test err: %v", err)
	} else {
		t.Log("Success test")
	}
}

func TestGetCorporationLists_Fail(t *testing.T) {

	var req GetCorpRequest
	req.Id = "KfqqvEJrbipy5"
	req.Numbers = []string{"9010601021385"}
	req.Version = "4"
	req.ResType = "12"
	req.RequestTimeOut = 10
	req.Endpoint = "http://localhost"
	req.History = "0"

	var api Api

	_, err := api.GetCorporationLists(&req)
	if err != nil {
		t.Logf("Success test err: %v", err)
	} else {
		t.Error("Failed test")
	}
}
