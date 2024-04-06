package worker

import (
	"sync"
	"time"

	pb "GoComputeFlow/pkg/worker/proto"
)

// Константы таймаутов вычислений по умолчанию
const (
	ADDTIMEOUT      = 5 * time.Second
	SUBTRACTTIMEOUT = 3 * time.Second
	MULTIPLYTIMEOUT = 4 * time.Second
	DIVIDETIMEOUT   = 6 * time.Second
)

const (
	COUNTWORKERS     = 5
	COUNTWORKERSFREE = 5
)

type Worker struct {
	Count           int
	CountFree       int
	Queue           []pb.TaskRequest
	ResultQueue     []pb.TaskRespons
	taskChannel     chan pb.TaskRequest
	PingTimeoutCalc []time.Time
	AddTimeout      time.Duration
	SubtractTimeout time.Duration
	MultiplyTimeout time.Duration
	DivideTimeout   time.Duration
	Mu              sync.Mutex
}
