package calculator

import (
	"GoComputeFlow/pkg/database"
	"fmt"
	"log"
	"sync"
	"time"
)

var Calc *FreeCalculators

// CreateCalculators создает новый экземпляр структуры счётчика свободных вычислителей
func CreateCalculators() {
	Calc = &FreeCalculators{
		Count:           5,
		CountFree:       5,
		Queue:           []TaskCalculate{},
		queueInProcess:  map[string]TaskCalculate{},
		taskChannel:     make(chan TaskCalculate),
		AddTimeout:      ADDTIMEOUT,
		SubtractTimeout: SUBTRACTTIMEOUT,
		MultiplyTimeout: MULTIPLYTIMEOUT,
		DivideTimeout:   DIVIDETIMEOUT,
		mu:              sync.Mutex{},
	}
	Calc.PingTimeoutCalc = make([]time.Time, Calc.Count)

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

	// Добавляю задачу в очередь
	Calc.mu.Lock()
	Calc.Queue = append(Calc.Queue, TaskCalculate{ID: userId, Expression: tokens})
	Calc.mu.Unlock()

	// Добавляю задачу в список вычислений юзера в базу данных
	if ok := database.AddExprssion(userId, expression); !ok {
		return false
	}

	return true
}

// GetTimeoutsOperations возвращает время вычислений для каждой из операций
func GetTimeoutsOperations() map[string]string {
	return map[string]string{
		"+": fmt.Sprintf("%.2f sec", Calc.AddTimeout.Seconds()),
		"-": fmt.Sprintf("%.2f sec", Calc.SubtractTimeout.Seconds()),
		"*": fmt.Sprintf("%.2f sec", Calc.MultiplyTimeout.Seconds()),
		"/": fmt.Sprintf("%.2f sec", Calc.DivideTimeout.Seconds()),
	}
}

//// RunCalculators запускает вычислители ожидающие очередь задач
//func (c *FreeCalculators) RunCalculators() {
//	for i := 0; i < c.Count; i++ {
//		go func(calcId int) {
//			for {
//				select {
//				case tokens := <-c.taskChannel:
//					log.Printf("Вычислитель %d - получил задачу: %s\n", calcId, tokens.ID)
//					result, flagError := c.calculateValue(calcId, tokens.Expression)
//
//					c.sendResult(tokens.ID, flagError, result)
//					log.Println("Вычислитель отправил результат в бд: ", tokens.ID)
//
//					// Переход в режим ожидания
//					c.CountFree++
//					continue
//
//				case <-time.After(3 * time.Second): // Пингуемся записывая текущее время в PingTimeoutCalc
//					log.Println("Вычислитель пингуется: ", calcId)
//					c.mu.Lock()
//					c.PingTimeoutCalc[calcId] = time.Now()
//					c.mu.Unlock()
//				case <-done:
//					// Завершение операций
//					log.Println("Вычислитель завершил работу: ", calcId)
//					return
//				}
//			}
//		}(i)
//	}
//}
//
//// sendResult - Отправка результата на оркестратор
//func (c *FreeCalculators) sendResult(idCalc string, flagError bool, result float64) {
//	// Отправка результатов и переход в режим ожидания
//	textResult := "error parse or calculate"
//	status := database.StatusError
//	if !flagError {
//		textResult = strconv.FormatFloat(result, 'f', -1, 64)
//		status = database.StatusCompleted
//	}
//
//	// Запись результатов обработки в базу данных
//	err := c.db.SetTaskResult(idCalc, status, textResult)
//	if err != nil {
//		log.Printf("Ошибка записи результата в базу данных: %s\n", err)
//		c.mu.Lock()
//		c.Queue = append(c.Queue, c.queueInProcess[idCalc])
//		c.mu.Unlock()
//		log.Println("Ошибочная операция перенесена в конец очереди...")
//	}
//
//	// Удаление задачи из очереди обработки
//	delete(c.queueInProcess, idCalc)
//}
//
//// calculateValue вычисляет значение выражения
//func (c *FreeCalculators) calculateValue(idCalc int, tokens []Token) (float64, bool) {
//	var result float64
//	flagError := false // Признак ошибки при выполнении операции
//	if len(tokens) == 0 {
//		flagError = true
//		log.Println("Очередь задач пустая, вычисление невозможно")
//		flagError = true
//	} else {
//		// Вычисление выражения
//		stack := make([]float64, 0)
//		for _, token := range tokens {
//			if !token.IsOp {
//				num, err := strconv.ParseFloat(token.Value, 64)
//				if err != nil {
//					log.Println("Ошибка при парсинге числа в вычислителе", err)
//					flagError = true
//					break
//				}
//				stack = append(stack, num)
//			} else {
//				if len(stack) < 2 {
//					log.Println("Для операции необходимо два числа в стеке, ошибка в вычислителе")
//					flagError = true
//					break
//				}
//				num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
//				stack = stack[:len(stack)-2]
//
//				switch token.Value {
//				case "+":
//					stack = append(stack, num1+num2)
//					time.Sleep(c.AddTimeout)
//				case "-":
//					stack = append(stack, num1-num2)
//					time.Sleep(c.SubtractTimeout)
//				case "*":
//					stack = append(stack, num1*num2)
//					time.Sleep(c.MultiplyTimeout)
//				case "/":
//					if num2 == 0 {
//						log.Println("Деление на ноль")
//						flagError = true
//						break
//					}
//					stack = append(stack, num1/num2)
//					time.Sleep(c.DivideTimeout)
//				default:
//					log.Println("Неизвестная операция в вычислителе")
//					flagError = true
//					break
//				}
//
//				// После кажого вычисления отправка пинга что вычислитель жив
//				c.mu.Lock()
//				c.PingTimeoutCalc[idCalc] = time.Now()
//				c.mu.Unlock()
//			}
//		}
//		if len(stack) != 1 {
//			log.Println("Слишком много чисел в стеке, ошибка в вычислителе")
//			flagError = true
//		}
//		result = stack[0]
//	}
//	return result, flagError
//}
