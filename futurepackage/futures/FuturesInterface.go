package futures

import (
	"time"
)

type Future interface {
	Cancel()
	Get() Result
	GetWithTimeout(duration time.Duration) Result
	IsCancelled() bool
	IsDone() bool
}


type Result struct {
	Value interface{}
	Error error
}