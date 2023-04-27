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
	"net"

	"github.com/PKUHPC/scow-go-demo/gen/go/hook"
	"google.golang.org/grpc"
)

type MyHookServer struct{}

func (s *MyHookServer) OnEvent(ctx context.Context, req *hook.OnEventRequest) (*hook.OnEventResponse, error) {

	log.Printf("Received event: %v", req)

	return &hook.OnEventResponse{}, nil
}

func StartHookServer() {

	addr := "0.0.0.0:5000"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	server := MyHookServer{}

	hook.RegisterHookServiceServer(grpcServer, &server)
	print("Listening at " + addr)

	grpcServer.Serve(lis)

}
