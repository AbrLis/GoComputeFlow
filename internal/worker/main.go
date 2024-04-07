package worker

import (
	"log"
	"strconv"
	"sync"
	"time"

	pb "GoComputeFlow/internal/worker/proto"
)

var DataWorker *Worker

// CreateWorker создает новый экземпляр структуры вычислителя
func CreateWorker(add, sub, mult, div time.Duration) {
	DataWorker = &Worker{
		Count:           COUNTWORKERS,
		CountFree:       COUNTWORKERSFREE,
		Queue:           make([]pb.TaskRequest, 0),
		ResultQueue:     make([]pb.TaskRespons, 0),
		taskChannel:     make(chan pb.TaskRequest),
		AddTimeout:      add,
		SubtractTimeout: sub,
		MultiplyTimeout: mult,
		DivideTimeout:   div,
		Mu:              sync.Mutex{},
	}
	DataWorker.PingTimeoutCalc = make([]time.Time, DataWorker.Count)

	RunWorkers()   // Запускаем вычислители
	RunAllocator() // Распределитель вычислений
}

// RunAllocator Запускает распределитель вычислений
func RunAllocator() {
	go func() {
		for {
			if DataWorker.CountFree > 0 {
				if len(DataWorker.Queue) == 0 {
					// Задач в очереди нет, ожидание 2 секунды...
					time.Sleep(2 * time.Second)
					continue
				}

				// Отправка задачи свободным вычислителям
				DataWorker.Mu.Lock()
				DataWorker.CountFree--
				task := DataWorker.Queue[0]
				// Перевод из очереди ожидания в очередь обработки
				DataWorker.Queue = DataWorker.Queue[1:]
				DataWorker.Mu.Unlock()
				DataWorker.taskChannel <- task
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
	for i := 0; i < DataWorker.Count; i++ {
		go func(calcId int) {
			for {
				select {
				case tokens := <-DataWorker.taskChannel:
					log.Printf("Вычислитель %d - получил задачу: %d\n", calcId, tokens.UserId)
					result, flagError := DataWorker.calculateValue(calcId, tokens.Expression)
					DataWorker.Mu.Lock()
					DataWorker.ResultQueue = append(
						DataWorker.ResultQueue,
						pb.TaskRespons{
							UserId: tokens.UserId, ExpressionId: tokens.ExpressionId, FlagError: flagError,
							Value: float32(result),
						},
					)
					DataWorker.Mu.Unlock()
					log.Println("Вычислитель отправил результат в очередь результатов: ", tokens.UserId)

					// Переход в режим ожидания
					DataWorker.CountFree++
					continue

				case <-time.After(3 * time.Second): // Пингуемся записывая текущее время в PingTimeoutCalc
					DataWorker.Mu.Lock()
					DataWorker.PingTimeoutCalc[calcId] = time.Now()
					DataWorker.Mu.Unlock()
				}
			}
		}(i)
	}
}

// calculateValue вычисляет значение выражения
func (c *Worker) calculateValue(idCalc int, tokens []*pb.Token) (float64, bool) {
	var result float64
	flagError := false // Признак ошибки при выполнении операции
	if len(tokens) == 0 {
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
				c.Mu.Lock()
				c.PingTimeoutCalc[idCalc] = time.Now()
				c.Mu.Unlock()
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
