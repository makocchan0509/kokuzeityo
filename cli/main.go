package main

import (
	"context"
	"kokuzeityo"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {

	address := "localhost:8080"
	appId := "KfqqvEJrbipy5"
	version := "4"
	corpNumbers := []string{"9010601021385", "6010501033839"}

	res := callGrpc(address, appId, version, corpNumbers)
	log.Printf("Response: %v ", res)
}

func callGrpc(address string, appId string, version string, corpNumbers []string) *kokuzeityo.GetCorpRes {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
	return res
}
