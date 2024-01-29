package wandb

type Run struct {
}

type Table struct{}

type Loggable interface {
	Table | ~int
}

func Init(entity string, project string) Run {
	return Run{}
}

func Log[T Loggable](key string, value T) {

}

func Finish() {

}
