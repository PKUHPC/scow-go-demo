/**
 * Copyright (c) 2022 Peking University and Peking University Institute for Computing and Digital Economy
 * SCOW is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/PKUHPC/scow-go-demo/gen/go/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

func CallApi() {
	// Assume mis-server is listening at 192.168.88.100:7571
	conn, err := grpc.Dial("192.168.88.100:7571", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	token := os.Getenv("SCOW_API_TOKEN")

	if token != "" {
		md := metadata.New(map[string]string{"Authorization": "Bearer " + token})
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	// create a AccountService client (protos/server/account.proto)
	client := server.NewAccountServiceClient(conn)

	// call `GetAccounts` RPC to get all accounts under default tenant
	resp, err := client.GetAccounts(ctx, &server.GetAccountsRequest{
		TenantName: "default",
	})

	if err != nil {
		panic(err)
	}

	log.Printf("Account list: %+v", protojson.Format(resp))
}
