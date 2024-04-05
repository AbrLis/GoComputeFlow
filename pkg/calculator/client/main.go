package client

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "GoComputeFlow/pkg/worker/proto"
)

func StartGPRCclient(host, port string) (pb.WorkerServiceClient, error) {
	addr := fmt.Sprintf("%s%s", host, port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	log.Printf("Start gRPC client on: %s%s/n", addr, port)

	return pb.NewWorkerServiceClient(conn), nil
}
