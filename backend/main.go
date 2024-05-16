package main

import (
	"backend/db"
	"backend/env"
	"backend/handler"
	"backend/logger"
	"backend/proto"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

func startGrpcServer(wg *sync.WaitGroup, db db.DB) {
	defer wg.Done()
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	proto.RegisterStreamingToolsServiceServer(s, &handler.StreamingToolsService{DB: db})
	// Serve gRPC Server
	lis, err := net.Listen("tcp", env.ServerCfgs.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Serving gRPC on 0.0.0.0:" + env.ServerCfgs.GrpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startHttpServer(wg *sync.WaitGroup) {
	defer wg.Done()
	// 2. 启动 HTTP 服务
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.Dial(
		"localhost:9001",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = proto.RegisterStreamingToolsServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    env.ServerCfgs.HttpPort,
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:" + env.ServerCfgs.HttpPort)
	log.Fatalln(gwServer.ListenAndServe())
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 初始化 log
	logger.InitSlog(env.LogCfgs)
	// 链接数据库
	database, err := db.Connect(env.DBcfgs)
	if err != nil {
		log.Println("Failed to connect database ", "reason: ", err.Error())
		return
	}
	// 启动 Grpc 服务
	go startGrpcServer(&wg, database)
	// 启动 Http 服务
	go startHttpServer(&wg)
	wg.Wait()
}
