package calculator

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"

	"GoComputeFlow/pkg/calculator/client"
	"GoComputeFlow/pkg/database"
	pb "GoComputeFlow/pkg/worker/proto"
)

var GrpcClient pb.WorkerServiceClient

// CreateCalculators создает новый экземпляр структуры счётчика свободных вычислителей
func CreateCalculators() {
	connect, err := client.StartGPRCclient(GRPChost, GRPCport)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	GrpcClient = connect

	// TODO: Добавить запуск распределителя вычислений в своём потоке, которой будет следить за очередью задач и
	// передавать задачи вычислителям, так же будет получать от них ответы и заносить результаты в бд
}

// AddExpressionToQueue добавляет выражение в очередь задач
func AddExpressionToQueue(expression string, userId uint) bool {
	// Парсим выражение
	tokens, err := ParseExpression(expression)
	if err != nil {
		log.Println("Error parsing expression: ", err)
		return false
	}

	// Передача задачи вычислителю
	_, err = GrpcClient.SetTask(
		context.TODO(), &pb.TaskRequest{
			UserId:     int32(userId),
			Expression: tokens,
		},
	)
	if err != nil {
		log.Println("Error set task to grpc: ", err)
		return false
	}

	// Добавляю задачу в список вычислений юзера в базу данных
	if ok := database.AddExprssion(userId, expression); !ok {
		return false
	}

	return true
}

// GetTimeoutsOperations возвращает время вычислений для каждой из операций
func GetTimeoutsOperations() map[string]string {
	timeouts, err := GrpcClient.GetTimeouts(context.TODO(), new(empty.Empty))
	if err != nil {
		log.Println("Error get timeouts from grpc: ", err)
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{
		"+": fmt.Sprintf("%s sec", timeouts.Add),
		"-": fmt.Sprintf("%s sec", timeouts.Subtract),
		"*": fmt.Sprintf("%s sec", timeouts.Multiply),
		"/": fmt.Sprintf("%s sec", timeouts.Divide),
	}
}
