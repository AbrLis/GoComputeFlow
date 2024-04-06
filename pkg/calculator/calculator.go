package calculator

import (
	"context"
	"fmt"
	"log"
	"time"

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

	// Запуск горутины что будет собирать выполненные задачи
	runCollector()
}

// runCollector запускает горутину собирающую выполненные задачи и записывает их в бд
func runCollector() {
	go func() {
		for {
			// Ждем задачу
			task, err := GrpcClient.GetResult(context.TODO(), &empty.Empty{})
			if task == nil {
				// Если нет задачи - ждем 2 секунды
				time.Sleep(2 * time.Second)
				continue
			}
			if err != nil {
				log.Println("!!Error getting result in grpc!!: ", err)
				continue
			}
			status := database.StatusCompleted
			if task.FlagError {
				status = database.StatusError
			}
			database.SetTaskResult(
				int(task.UserId),
				int(task.ExpressionId),
				status,
				task.Value,
			)
		}
	}()
}

// AddExpressionToQueue добавляет выражение в очередь задач
func AddExpressionToQueue(expression string, userId uint) bool {
	// Парсим выражение
	tokens, err := ParseExpression(expression)
	if err != nil {
		log.Println("Error parsing expression: ", err)
		return false
	}

	// Добавляю задачу в список вычислений юзера в базу данных
	expressionID, ok := database.AddExprssion(userId, expression)
	if !ok {
		return false
	}

	// Передача задачи вычислителю
	_, err = GrpcClient.SetTask(
		context.TODO(), &pb.TaskRequest{
			UserId:       uint32(userId),
			ExpressionId: uint32(expressionID),
			Expression:   tokens,
		},
	)
	if err != nil {
		log.Println("Error set task to grpc: ", err)
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
		"+": fmt.Sprintf("%s", timeouts.Add),
		"-": fmt.Sprintf("%s", timeouts.Subtract),
		"*": fmt.Sprintf("%s", timeouts.Multiply),
		"/": fmt.Sprintf("%s", timeouts.Divide),
	}
}
