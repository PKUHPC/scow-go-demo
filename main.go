package main

import (
	"context"
	"log"

	"github.com/PKUHPC/scow-grpc-api-client-demo/gen/go/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("192.168.88.100:7571", grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := server.NewAccountServiceClient(conn)

	resp, err := client.GetAccounts(context.Background(), &server.GetAccountsRequest{})

	if err != nil {
		panic(err)
	}

	log.Printf("Account list: %v", resp)
}
