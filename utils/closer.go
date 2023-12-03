package utils

import "sync"

// Closer 用于资源回收的信号控制
type Closer struct {
	waiting     sync.WaitGroup
	closeSignal chan struct{}
}

func NewCloser(i int) *Closer {
	closer := &Closer{waiting: sync.WaitGroup{}}
	closer.waiting.Add(i)
	closer.closeSignal = make(chan struct{})
	return closer
}

// Close 上游通知下游协程进行资源回收, 并等待协程通知回收完毕
func (c *Closer) Close() {
	close(c.closeSignal)
	c.waiting.Wait()
}

// Done 标识协程已经完成额资源回收, 通知上游正式关闭
func (c *Closer) Done() {
	c.waiting.Done()
}

// Wait 返回关闭信号
func (c *Closer) Wait() chan struct{} {
	return c.closeSignal
}
