package calculator

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/durationpb"

	"GoComputeFlow/internal/calculator/client"
	"GoComputeFlow/internal/database"
	"GoComputeFlow/internal/models"
	pb "GoComputeFlow/internal/worker/proto"
)

var GrpcClient pb.WorkerServiceClient

// CreateCalculators создает новый экземпляр структуры счётчика свободных вычислителей
func CreateCalculators() {
	connect, err := client.StartGPRCclient(GRPChost, GRPCport)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	GrpcClient = connect

	// Проверка базы данных, остались ли незавершённые задачи после отключения сервера
	data, err := database.GetAllUnfinishedTasks()
	if err != nil {
		log.Println("!!Error getting unfinished tasks: ", err)
		return
	}
	for _, v := range data {
		// Добавление незавершённых задач в очередь
		_ = AddExpressionToQueue(v.Expression, v.UserId, false, v.ID)
	}
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
			status := models.StatusCompleted
			if task.FlagError {
				status = models.StatusError
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
func AddExpressionToQueue(expression string, userId uint, newTask bool, expressionID uint) bool {
	// Парсим выражение
	tokens, err := ParseExpression(expression)
	if err != nil {
		log.Println("Error parsing expression: ", err)
		return false
	}

	// Если задача новая, то добавляем её в список вычислений юзера
	var ok bool
	if newTask {
		if expressionID, ok = database.AddExprssion(userId, expression); !ok {
			return false
		}
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

// SetTimeoutsOperations устанавливает время вычислений для каждой из операций
func SetTimeoutsOperations(add, subtract, multiply, divide string) (string, error) {
	timers := database.GetTimeouts()
	parseAndSetTimeout(add, &timers.AddTimeout)
	parseAndSetTimeout(subtract, &timers.SubtractTimeout)
	parseAndSetTimeout(multiply, &timers.MutiplyTimeout)
	parseAndSetTimeout(divide, &timers.DivideTimeout)

	_, err := GrpcClient.SetTimeouts(
		context.TODO(), &pb.TimeoutsRequest{
			Add:      durationpb.New(timers.AddTimeout),
			Subtract: durationpb.New(timers.SubtractTimeout),
			Multiply: durationpb.New(timers.MutiplyTimeout),
			Divide:   durationpb.New(timers.DivideTimeout),
		},
	)
	if err != nil {
		log.Println("Error set timeouts to grpc: ", err)
		return "", err
	}

	// Установка новых занчений таймаутов в бд
	database.SetTimeouts(timers.AddTimeout, timers.SubtractTimeout, timers.MutiplyTimeout, timers.DivideTimeout)
	return fmt.Sprintf(
		"Timeouts set: %s, %s, %s, %s", timers.AddTimeout, timers.SubtractTimeout, timers.MutiplyTimeout,
		timers.DivideTimeout,
	), nil
}

// GetWorkersTimeouts возвращает овтет воркеров о пингах воркеров
func GetWorkersTimeouts() (map[string]string, error) {
	pings, err := GrpcClient.GetPing(context.TODO(), new(empty.Empty))
	if err != nil {
		log.Println("Error get pings from grpc: ", err)
		return nil, err
	}

	result := make(map[string]string, len(pings.Ping))
	for _, v := range pings.Ping {
		result[v.Name] = fmt.Sprintf("%s", v.Ping)
	}
	return result, nil
}
