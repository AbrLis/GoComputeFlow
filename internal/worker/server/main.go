package worker

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"GoComputeFlow/internal/worker"
	pb "GoComputeFlow/internal/worker/proto"
)

type Server struct {
	pb.WorkerServiceServer
}

func NewServer() *Server {
	return &Server{}
}

// StartGRPCServerWorker запускает gRPC-сервер
func StartGRPCServerWorker(host, port string) {
	go func() {
		addr := fmt.Sprintf("%s%s", host, port)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal("Error start listener gRPC worker: ", err)
		}
		log.Printf("Start server gRPC worker on: %s%s", addr, port)

		grpcServer := grpc.NewServer()
		myServiceDriver := NewServer()
		pb.RegisterWorkerServiceServer(grpcServer, myServiceDriver)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("Error start server gRPC worker: ", err)
		}
		log.Printf("I`m stop?")
	}()
}

// GetTimeouts возвращает таймауты вычислений в текстовом виде
func (s *Server) GetTimeouts(context.Context, *empty.Empty) (*pb.TimeoutsMessage, error) {
	return &pb.TimeoutsMessage{
		Add:      fmt.Sprintf("%.2f sec", worker.DataWorker.AddTimeout.Seconds()),
		Subtract: fmt.Sprintf("%.2f sec", worker.DataWorker.SubtractTimeout.Seconds()),
		Multiply: fmt.Sprintf("%.2f sec", worker.DataWorker.MultiplyTimeout.Seconds()),
		Divide:   fmt.Sprintf("%.2f sec", worker.DataWorker.DivideTimeout.Seconds()),
	}, nil
}

// SetTimeouts устанавливает таймауты вычислений
func (s *Server) SetTimeouts(_ context.Context, req *pb.TimeoutsRequest) (*empty.Empty, error) {
	worker.DataWorker.AddTimeout = req.Add.AsDuration()
	worker.DataWorker.SubtractTimeout = req.Subtract.AsDuration()
	worker.DataWorker.MultiplyTimeout = req.Multiply.AsDuration()
	worker.DataWorker.DivideTimeout = req.Divide.AsDuration()

	return &empty.Empty{}, nil
}

// SetTask Добавляет задачу в очередь исполнения
func (s *Server) SetTask(_ context.Context, req *pb.TaskRequest) (*empty.Empty, error) {
	worker.DataWorker.Mu.Lock()
	worker.DataWorker.Queue = append(
		worker.DataWorker.Queue, pb.TaskRequest{
			UserId:       req.UserId,
			ExpressionId: req.ExpressionId,
			Expression:   req.Expression,
		},
	)
	worker.DataWorker.Mu.Unlock()

	return &empty.Empty{}, nil
}

// GetResult возвращает результат вычисления
func (s *Server) GetResult(context.Context, *empty.Empty) (*pb.TaskRespons, error) {
	if worker.DataWorker.ResultQueue == nil || len(worker.DataWorker.ResultQueue) == 0 {
		return nil, status.Error(codes.NotFound, "empty results")
	}
	//
	worker.DataWorker.Mu.Lock()
	data := worker.DataWorker.ResultQueue[0]
	worker.DataWorker.ResultQueue = worker.DataWorker.ResultQueue[1:]
	worker.DataWorker.Mu.Unlock()

	return &data, nil
}

// GetPing возвращает пинг воркеров
func (s *Server) GetPing(context.Context, *empty.Empty) (*pb.PingResponse, error) {
	worker.DataWorker.Mu.Lock()
	defer worker.DataWorker.Mu.Unlock()

	var ping *pb.PingMessage
	var pingData []*pb.PingMessage

	for i, v := range worker.DataWorker.PingTimeoutCalc {
		name := fmt.Sprintf("worker %d", i)

		ping = &pb.PingMessage{Name: name, Ping: fmt.Sprintf("%.2f sec", time.Since(v).Seconds())}
		pingData = append(pingData, ping)
	}

	return &pb.PingResponse{Ping: pingData}, nil
}
