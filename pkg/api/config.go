package api

const (
	apiVersion         = "api/v1"
	addExpressionPath  = "/add-expression"  // API для добавления арифметического выражения
	getExpressionsPath = "/get-expressions" // API для получения списка арифметических выражений
	getValuePath       = "/get-value/"      // API для получения значения выражения по его идентификатору
	getOperationsPath  = "/get-operations"  // API для получения списка доступных операций со временем их выполнения
	monitoring         = "/monitoring"      // API для получения статуса вычислителей (времени последнего пинга)
	HostPath           = "localhost"        // Путь до хоста
	PortHost           = ":3000"            // Порт хоста
	registerPath       = "/register"        // API для регистрации пользователя
	LoginPath          = "/login"           // API для входа пользователя
)
