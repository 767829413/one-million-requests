package plan

var (
	Queue     chan interface{}
	MAX_QUEUE = 1000
)

func Run() {
	Queue = make(chan interface{}, MAX_QUEUE)
	go func() {
		for {
			select {
			case job, ok := <-Queue:
				if !ok {
					break
				}
				_ = job
				// do something
			}
		}
	}()
}

func Handler2(payloads []interface{}) {
	Run()
	for _, payload := range payloads {
		Queue <- payload
	}
}
