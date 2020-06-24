package futures

import (
	"time"
	"sync"
)

type FutureTask struct {
	channel chan Result
	complete bool
	cancelled bool
	result Result
	mutex sync.Mutex
}
func (futureTask *FutureTask) Cancel() {
	if !futureTask.IsComplete() && !futureTask.IsCancelled() {
		futureTask.channel <- Result{Error: &InterruptError{desc: "Cancelled manually"}}
		futureTask.cancelled = true
	}
}
func (futureTask *FutureTask) IsCancelled() bool{
	return futureTask.cancelled
}
func (futureTask *FutureTask) IsComplete() bool{
	return futureTask.complete
}
func (futureTask *FutureTask) Get() Result {
	defer futureTask.Cancel() 
	futureTask.mutex.Lock()
	defer futureTask.mutex.Unlock()
	if futureTask.complete || futureTask.cancelled{
		return futureTask.result
	}
	futureTask.result = <- futureTask.channel
	futureTask.complete = true
	return futureTask.result
	
}
func (futureTask *FutureTask) GetWithTimeout(duration time.Duration) Result {
	ch := make(chan Result)
	go func(){
		ch <- futureTask.Get()
	}()
	select {
		case <-time.After(duration): return Result{Value: nil, Error: &TimeoutError{desc:"TimeOut Occured"}}
		case <- ch : return futureTask.result
	}
}

func MakeFuture( function func() Result) *FutureTask {
	task := FutureTask{channel : make(chan Result)}

	go func(){
		y := function()
		task.channel <- y
	}()
	return &task
}
