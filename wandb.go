package wandb

type Run struct {
}

type Table struct{}

type Loggable interface {
	Table | ~int | ~float64 | ~float32 | ~string
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
