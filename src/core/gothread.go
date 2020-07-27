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

func (self *GoThread) Create() {

	self.started = make(chan bool)
	self.stoped = make(chan bool)
	GetLogger().Debug("GoThread was created")

	self.started <- true

	for {
		GetLogger().Debug("GoThread is running")
		select {
		case <- self.started:
			GetLogger().Debug("GoThread started")
		case <- self.stoped:
			if self.Doom != nil {
				self.Doom()
			}
			GetLogger().Debug("GoThread stoped")
			return
		default:
			if self.Routine != nil {
				self.Routine()
			}
		}
	}
}

func (self *GoThread) Start() {
	<- self.started
	self.status = GTHREAD_START
}

func (self *GoThread) Stop() {
	if self.status != GTHREAD_STOP {
		self.stoped <- true
		self.status = GTHREAD_STOP
	}
}

