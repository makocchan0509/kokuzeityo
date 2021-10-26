package main

import (
	"context"
	"kokuzeityo/http"
	"kokuzeityo/parser"
	"log"
	"net"
	"os"
	"strconv"

	"kokuzeityo"

	"google.golang.org/grpc"
)

var appPort string
var cashMode string
var endpoint string
var reqTimeOutSec int

var defaultPort = "8080"
var defaultMode = "0"
var defaultTimeOutSec = 10

type server struct {
	kokuzeityo.UnimplementedKokuzeityoServer
}

func validate(in *kokuzeityo.GetCorpReq) (bool, string) {

	result := true
	msg := ""

	if in.AppId == "" {
		result = false
		msg = "invalid App Id. You must set AppId."
		log.Println(msg)
		return result, msg
	}

	if len(in.CorporateNumber) == 0 {
		result = false
		msg = "invalid CorporateNumber. You must set least one Corporate number"
		log.Println(msg)
		return result, msg
	}

	if in.Version == "" || in.Version != "4" {
		result = false
		msg = "invalid Version. You must set version 4."
		log.Println(msg)
		return result, msg
	}

	return result, msg

}

func (s *server) GetCorporation(ctx context.Context, in *kokuzeityo.GetCorpReq) (*kokuzeityo.GetCorpRes, error) {
	log.Printf("Received: %v", in)

	result, msg := validate(in)

	if !result {
		return &kokuzeityo.GetCorpRes{
			Result: "1",
			Msg:    msg,
			Count:  0,
		}, nil
	}

	var req http.GetCorpRequest
	req.Id = in.AppId
	req.Numbers = in.CorporateNumber
	req.Version = "4"
	req.ResType = "12"
	req.RequestTimeOut = reqTimeOutSec
	req.Endpoint = endpoint
	req.History = "0"

	var api http.Api

	res, err := api.GetCorporationLists(&req)
	if err != nil {
		return &kokuzeityo.GetCorpRes{
			Result: "1",
			Msg:    "error",
			Count:  0,
		}, nil
	}

	corps := xmlToGrpcRes(res.Kcorporation, res.Kcount.Text)

	return &kokuzeityo.GetCorpRes{
		Result: "0",
		Msg:    "",
		Count:  int32(res.Kcount.Text),
		Corps:  corps,
	}, nil

}

func main() {

	getEnv()
	lis, err := net.Listen("tcp", ":"+appPort)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()
	kokuzeityo.RegisterKokuzeityoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv() {
	appPort = os.Getenv("APP_PORT")
	cashMode = os.Getenv("CASH_MODE")
	endpoint = os.Getenv("ENDPOINT")
	sec := os.Getenv("REQUEST_TIMEOUT")

	if appPort == "" {
		appPort = defaultPort
	}

	if cashMode == "" || cashMode != "0" {
		cashMode = defaultMode
	}

	if endpoint == "" {
		log.Fatalf("You must set endpoint value to environment parameter.")
	}

	if sec == "" {
		reqTimeOutSec = defaultTimeOutSec
	} else {
		t, err := strconv.Atoi(sec)
		if err != nil {
			log.Fatalf("Invalid http request timer.")
		}
		reqTimeOutSec = t
	}
}
func xmlToGrpcRes(corps []*parser.Kcorporation, count int) []*kokuzeityo.Corp {

	var result []*kokuzeityo.Corp
	for i := 0; i < count; i++ {
		var corp kokuzeityo.Corp
		corp.CorporateNumber = corps[i].KcorporateNumber.Text
		corp.UpdateDate = corps[i].KupdateDate.Text
		corp.ChangeDate = corps[i].KchangeDate.Text
		corp.CorpName = corps[i].Kname.Text
		corp.NameImageId = corps[i].KnameImageId.Text
		corp.Kind = corps[i].Kkind.Text
		corp.PrefectureName = corps[i].KprefectureName.Text
		corp.CityName = corps[i].KcityName.Text
		corp.StreetNumber = corps[i].KstreetNumber.Text
		corp.AddressImageId = corps[i].KaddressImageId.Text
		corp.PrefectureCd = corps[i].KprefectureCode.Text
		corp.CityCode = corps[i].KcityCode.Text
		corp.PostCode = corps[i].KpostCode.Text
		corp.AddressOutside = corps[i].KaddressOutside.Text
		corp.AddressOutsideImageId = corps[i].KaddressOutsideImageId.Text
		corp.CloseDate = corps[i].KcloseDate.Text
		corp.CloseCause = corps[i].KcloseCause.Text
		corp.SuccessorCorporateNumber = corps[i].KsuccessorCorporateNumber.Text
		corp.ChangeCause = corps[i].KchangeCause.Text
		corp.AssignmentDate = corps[i].KassignmentDate.Text
		corp.Latest = corps[i].Klatest.Text
		corp.EnName = corps[i].KenName.Text
		corp.EnPrefectureName = corps[i].KenPrefectureName.Text
		corp.EnCityName = corps[i].KenCityName.Text
		corp.EnAddressOutside = corps[i].KenAddressOutside.Text
		corp.Furigana = corps[i].Kfurigana.Text
		corp.Hihyoji = corps[i].Khihyoji.Text

		result = append(result, &corp)
	}
	return result
}
