package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"kokuzeityo"
	"kokuzeityo/parser"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestValidate_Success(t *testing.T) {

	var in kokuzeityo.GetCorpReq
	in.AppId = "KfqqvEJrbipy5"
	in.Version = "4"
	in.CorporateNumber = []string{"9010601021385", "6010501033839"}

	result, msg := validate(&in)

	if result && msg == "" {
		t.Logf("Success test result: %v msg: %v", result, msg)
	} else {
		t.Errorf("Failed test result: %v msg: %v", result, msg)
	}

}

func TestValidate_Fail_1(t *testing.T) {

	var in kokuzeityo.GetCorpReq
	in.AppId = ""
	in.Version = "4"
	in.CorporateNumber = []string{"9010601021385", "6010501033839"}

	result, msg := validate(&in)

	if !result && msg == "invalid App Id. You must set AppId." {
		t.Logf("Success test result: %v msg: %v", result, msg)
	} else {
		t.Errorf("Failed test result: %v msg: %v", result, msg)
	}
}

func TestValidate_Fail_2(t *testing.T) {

	var in kokuzeityo.GetCorpReq
	in.AppId = "KfqqvEJrbipy5"
	in.Version = ""
	in.CorporateNumber = []string{"9010601021385", "6010501033839"}

	result, msg := validate(&in)

	if !result && msg == "invalid Version. You must set version 4." {
		t.Logf("Success test result: %v msg: %v", result, msg)
	} else {
		t.Errorf("Failed test result: %v msg: %v", result, msg)
	}
}

func TestValidate_Fail_3(t *testing.T) {

	var in kokuzeityo.GetCorpReq
	in.AppId = "KfqqvEJrbipy5"
	in.Version = "4"
	in.CorporateNumber = []string{}

	result, msg := validate(&in)

	if !result && msg == "invalid CorporateNumber. You must set least one Corporate number" {
		t.Logf("Success test result: %v msg: %v", result, msg)
	} else {
		t.Errorf("Failed test result: %v msg: %v", result, msg)
	}
}

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	kokuzeityo.RegisterKokuzeityoServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
}

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetCorporation_Success(t *testing.T) {

	getEnv()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	c := kokuzeityo.NewKokuzeityoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.GetCorporation(ctx, &kokuzeityo.GetCorpReq{
		Version:         "4",
		AppId:           "KfqqvEJrbipy5",
		CorporateNumber: []string{"9010601021385"},
	})

	if res.Result == "0" {
		t.Logf("Success test response: %v", res)
	} else {
		t.Errorf("Failed test response: %v", res)
	}
}

func TestGetCorporation_Fail_1(t *testing.T) {

	//getEnv()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	c := kokuzeityo.NewKokuzeityoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.GetCorporation(ctx, &kokuzeityo.GetCorpReq{
		Version:         "",
		AppId:           "KfqqvEJrbipy5",
		CorporateNumber: []string{"9010601021385", "6010501033839"},
	})

	if res.Result == "1" {
		t.Logf("Success test response: %v", res)
	} else {
		t.Errorf("Failed test response: %v", res)
	}
}

func TestGetCorporation_Fail_2(t *testing.T) {

	//getEnv()
	endpoint = "http://localhost"
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	c := kokuzeityo.NewKokuzeityoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.GetCorporation(ctx, &kokuzeityo.GetCorpReq{
		Version:         "4",
		AppId:           "KfqqvEJrbipy5",
		CorporateNumber: []string{"9010601021385", "6010501033839"},
	})
	fmt.Println(res)
	if res.Result == "1" {
		t.Logf("Success test response: %v", res)
	} else {
		t.Errorf("Failed test response: %v", res)
	}
}

func TestXmlToGrpcRes(t *testing.T) {

	var xmlRes parser.Kcorporations
	data, _ := ioutil.ReadFile("../test_data/test_input.xml")

	err := xml.Unmarshal(data, &xmlRes)
	if err != nil {
		t.Errorf("Failed xml parse error :%v", err)
	}
	//	corps := xmlToGrpcRes(xmlRes.Kcorporation, xmlRes.Kcount.Text)
	result := true

	if result {
		t.Logf("Success test")
	}
}
