package worker

import (
	"sync"
	"time"

	pb "GoComputeFlow/internal/worker/proto"
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
