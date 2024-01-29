package wandb

import "math/rand"

type Run struct {
}

type Table struct{}

type Loggable interface {
	Table | ~int | ~float64 | ~float32 | ~string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func generateRunId() string {
	return randStringBytes(10)
}

func Init(entity string, project string) Run {
	return Run{}
}

func Log[T Loggable](key string, value T) {

}

func Finish() {

}

func SetConfig(config map[string]any) {
}

// TODO: have a messenger object that groups requests etc.
// TODO: register signal handlers for SIGTERM and SIGINT to call messenger.finish()
