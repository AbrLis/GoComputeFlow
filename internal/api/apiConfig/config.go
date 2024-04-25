package apiConfig

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	ApiVersion         = "api/v1"
	AddExpressionPath  = "/add-expression"     // API для добавления арифметического выражения
	GetExpressionsPath = "/get-expressions"    // API для получения списка арифметических выражений
	GetValuePath       = "/get-value/:task_id" // API для получения значения выражения по его идентификатору
	GetOperationsPath  = "/get-operations"     // API для получения списка доступных операций со временем их выполнения
	SetOperationsPath  = "/set-operations"     // API для установки времени выполнения операции
	MonitoringPath     = "/monitoring"         // API для получения статуса вычислителей (времени последнего пинга)
	HostPath           = "0.0.0.0"             // Путь до хоста
	PortHost           = ":3000"               // Порт хоста
	RegisterPath       = "/register"           // API для регистрации пользователя
	LoginPath          = "/login"              // API для входа пользователя
)

const (
	DefaultCountExpressions = 100 // Возвращаемое количество выражений по умолчанию для общего вывода
)

var SECRETKEY string

func init() {
	_ = godotenv.Load(".env") // Переменная подгружается либо из .env файла, либо из переменных окружения контейнера
	SECRETKEY = os.Getenv("SECRETKEY")
	if SECRETKEY == "" {
		log.Fatal("SECRETKEY is not set")
	}
}
