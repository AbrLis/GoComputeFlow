package api

const (
	apiVersion         = "api/v1"
	addExpressionPath  = "/add-expression"     // API для добавления арифметического выражения
	getExpressionsPath = "/get-expressions"    // API для получения списка арифметических выражений
	getValuePath       = "/get-value/:task_id" // API для получения значения выражения по его идентификатору
	getOperationsPath  = "/get-operations"     // API для получения списка доступных операций со временем их выполнения
	setOperationsPath  = "/set-operations"     // API для установки времени выполнения операции
	monitoringPath     = "/monitoring"         // API для получения статуса вычислителей (времени последнего пинга)
	HostPath           = "0.0.0.0"             // Путь до хоста
	PortHost           = ":3000"               // Порт хоста
	registerPath       = "/register"           // API для регистрации пользователя
	LoginPath          = "/login"              // API для входа пользователя
)

const (
	DefaultCountExpressions = 100 // Возвращаемое количество выражений по умолчанию для общего вывода
)
