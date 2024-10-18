package queue

const (
	QueueTestQueue  = "test:queue"
	QueueTestQueue2 = "test:queue:2"
)

var AllQueues = map[string]int{
	QueueTestQueue:  1,
	QueueTestQueue2: 1,
}
