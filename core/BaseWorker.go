package core

// go语言的问题，莫得方法重载，用worker 来代替

type IWorker interface {
	OnEnter(tick *Tick)
	OnOpen(tick *Tick)
	OnClose(tick *Tick)
	OnExit(tick *Tick)
	OnTick(tick *Tick) NodeStatus
}

type Worker struct {
}

func (w *Worker) OnEnter(tick *Tick) {
}

func (w *Worker) OnOpen(tick *Tick) {
}

func (w *Worker) OnTick(tick *Tick) NodeStatus {
	// 默认成功
	return STATUS_SUCCESS
}

func (w *Worker) OnClose(tick *Tick) {
}

func (w *Worker) OnExit(tick *Tick) {
}
