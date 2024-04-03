package worker

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var worker *Worker

// CreateWorker создает новый экземпляр структуры вычислителя
func CreateWorker() {
	worker = &Worker{
		Count:           COUNTWORKERS,
		CountFree:       COUNTWORKERSFREE,
		Queue:           []TaskCalculate{},
		ResultQueue:     []float64{},
		taskChannel:     make(chan TaskCalculate),
		AddTimeout:      ADDTIMEOUT,
		SubtractTimeout: SUBTRACTTIMEOUT,
		MultiplyTimeout: MULTIPLYTIMEOUT,
		DivideTimeout:   DIVIDETIMEOUT,
		mu:              sync.Mutex{},
	}
	worker.PingTimeoutCalc = make([]time.Time, worker.Count)

	RunWorkers()   // Запускаем вычислители
	RunAllocator() // Распределитель вычислений
}

// RunAllocator Запускает распределитель вычислений
func RunAllocator() {
	go func() {
		for {
			if worker.CountFree > 0 {
				if len(worker.Queue) == 0 {
					log.Println("Задач в очереди нет, ожидание 2 секунды...")
					time.Sleep(2 * time.Second)
					continue
				}

				// Отправка задачи свободным вычислителям
				worker.mu.Lock()
				worker.CountFree--
				task := worker.Queue[0]
				// Перевод из очереди ожидания в очередь обработки
				worker.Queue = worker.Queue[1:]
				worker.mu.Unlock()
				worker.taskChannel <- task
			} else {
				// Ждать пока не появятся свободные вычислители
				log.Println("Нет свободных вычислителей, ожидание...")
				time.Sleep(2 * time.Second)
				continue
			}
		}
	}()
}

// RunWorkers Запускает вычислители в потоках
func RunWorkers() {
	for i := 0; i < worker.Count; i++ {
		go func(calcId int) {
			for {
				select {
				case tokens := <-worker.taskChannel:
					log.Printf("Вычислитель %d - получил задачу: %d\n", calcId, tokens.ID)
					result, flagError := worker.calculateValue(calcId, tokens.Expression)
					fmt.Println(result, flagError) // !ЗАГЛУШКА

					//worker.sendResult(tokens.ID, flagError, result)
					// TODO: Реализовать передачу результата по gRPC
					log.Println("Вычислитель отправил результат в бд: ", tokens.ID)

					// Переход в режим ожидания
					worker.CountFree++
					continue

				case <-time.After(3 * time.Second): // Пингуемся записывая текущее время в PingTimeoutCalc
					log.Println("Вычислитель пингуется: ", calcId)
					worker.mu.Lock()
					worker.PingTimeoutCalc[calcId] = time.Now()
					worker.mu.Unlock()
				}
			}
		}(i)
	}
}

// calculateValue вычисляет значение выражения
func (c *Worker) calculateValue(idCalc int, tokens []Token) (float64, bool) {
	var result float64
	flagError := false // Признак ошибки при выполнении операции
	if len(tokens) == 0 {
		flagError = true
		log.Println("Очередь задач пустая, вычисление невозможно")
		flagError = true
	} else {
		// Вычисление выражения
		stack := make([]float64, 0)
		for _, token := range tokens {
			if !token.IsOp {
				num, err := strconv.ParseFloat(token.Value, 64)
				if err != nil {
					log.Println("Ошибка при парсинге числа в вычислителе", err)
					flagError = true
					break
				}
				stack = append(stack, num)
			} else {
				if len(stack) < 2 {
					log.Println("Для операции необходимо два числа в стеке, ошибка в вычислителе")
					flagError = true
					break
				}
				num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]

				switch token.Value {
				case "+":
					stack = append(stack, num1+num2)
					time.Sleep(c.AddTimeout)
				case "-":
					stack = append(stack, num1-num2)
					time.Sleep(c.SubtractTimeout)
				case "*":
					stack = append(stack, num1*num2)
					time.Sleep(c.MultiplyTimeout)
				case "/":
					if num2 == 0 {
						log.Println("Деление на ноль")
						flagError = true
						break
					}
					stack = append(stack, num1/num2)
					time.Sleep(c.DivideTimeout)
				default:
					log.Println("Неизвестная операция в вычислителе")
					flagError = true
					break
				}

				// После кажого вычисления отправка пинга что вычислитель жив
				c.mu.Lock()
				c.PingTimeoutCalc[idCalc] = time.Now()
				c.mu.Unlock()
			}
		}
		if len(stack) != 1 {
			log.Println("Слишком много чисел в стеке, ошибка в вычислителе")
			flagError = true
		}
		result = stack[0]
	}
	return result, flagError
}
