package core

const(
	GTHREAD_START = 0
	GTHREAD_STOP = 1
	GTHREAD_RUNNING = 2
	GTHREAD_YIELD = 3
	GTHREAD_RESUME = 4
)

type GoThread struct {
	name string

	started chan bool
	stoped chan bool

	status int
	rotate int

	Doom func()
	Routine func()
}

/*
====================
Create
	线程的创建
====================
*/
func (self *GoThread) Create() {

	// 创建无缓冲通道
	self.started = make(chan bool)
	self.stoped = make(chan bool)
	GetLogger().Debug("GoThread(", &self, ") was created")
}

/*
====================
Create
	线程启动并等待线程结束
====================
*/
func (self *GoThread) Join() {

	self.started <- true

	for {
		GetLogger().Debug("GoThread(", &self, ") is running")
		select {
		case <- self.started:
			GetLogger().Debug("GoThread(", &self, ") started")
		case <- self.stoped:
			if self.Doom != nil {
				self.Doom()
			}
			GetLogger().Debug("GoThread(", &self, ") stoped")
			return
		default:
			if self.Routine != nil {
				self.Routine()
			}
		}
	}
}

/*
====================
Create
	线程开始
====================
*/
func (self *GoThread) Start() {
	<- self.started
	self.status = GTHREAD_START
}

/*
====================
Create
	线程结束
====================
*/
func (self *GoThread) Stop() {
	if self.status != GTHREAD_STOP {
		self.stoped <- true
		self.status = GTHREAD_STOP
	}
}

/*
====================
Create
	线程结束并清理
====================
*/
func (self *GoThread) Close() {
	close(self.started)
	close(self.stoped)
	GetLogger().Debug("GoThread(", &self, ") closed")
}

