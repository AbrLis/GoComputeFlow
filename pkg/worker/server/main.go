package worker

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	"GoComputeFlow/pkg/worker"
	pb "GoComputeFlow/pkg/worker/proto"
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
		Add:      fmt.Sprintf("%.2f sec", worker.ADDTIMEOUT.Seconds()),
		Subtract: fmt.Sprintf("%.2f sec", worker.SUBTRACTTIMEOUT.Seconds()),
		Multiply: fmt.Sprintf("%.2f sec", worker.MULTIPLYTIMEOUT.Seconds()),
		Divide:   fmt.Sprintf("%.2f sec", worker.DIVIDETIMEOUT.Seconds()),
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
	expression := make([]worker.Token, len(req.Expression))
	for i, token := range req.Expression {
		expression[i] = worker.Token{
			Value: token.Value,
			IsOp:  token.IsOp,
		}
	}

	worker.DataWorker.Queue = append(
		worker.DataWorker.Queue, worker.TaskCalculate{
			ID:         uint(req.UserId),
			Expression: expression,
		},
	)

	return &empty.Empty{}, nil
}

// GetResult возвращает результат вычисления
func (s *Server) GetResult(context.Context, *empty.Empty) (*pb.TaskRespons, error) {
	if worker.DataWorker.ResultQueue == nil || len(worker.DataWorker.ResultQueue) == 0 {
		return nil, fmt.Errorf("no results")
	}

	worker.DataWorker.Mu.Lock()
	data := worker.DataWorker.ResultQueue[0]
	worker.DataWorker.ResultQueue = worker.DataWorker.ResultQueue[1:]
	worker.DataWorker.Mu.Unlock()

	return &pb.TaskRespons{UserId: int32(data.ID), Value: float32(data.Result), FlagError: data.FlagError}, nil
}
