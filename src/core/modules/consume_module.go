package modules

import (
	"msgback/conf_struct"
	"msgback/core"
)

type ConsumeModuleImpl interface {

	Init_after_module()

	Consumer()

	Handler()
}

type ConsumeModule struct {

	ConsumeModuleImpl

	Started chan bool

	Stoped chan bool

	Status int

	Conf conf_struct.CoreConf

	Logger core.Logger

	Recorder core.Logger

	Hooker struct {
		ConsumerHook func() *string

		HandlerHook func(string2 *string)

		ClearHook func()
	}
}

/*func (self *ConsumeModule) throw(e interface{}) {
	panic(e)
}

func (self *ConsumeModule) catch() {
	switch err := recover(); err {
	case nil:
		self.logger.Info("no panic")
	default:
		self.logger.Error(err)
	}

	time.Sleep(1*time.Second)

}*/

func (self *ConsumeModule) Init_module(cycle *core.Cycle) {

	self.Logger = core.GetLogger()
	self.Recorder = core.GetRecorder()

	self.Started = make(chan bool)
	self.Stoped = make(chan bool)

	self.Conf = core.GetInjectModule(core.Conf{}).(core.ConfImpl).GetAll()

	self.Logger.Info("Consume module init")

}

func (self *ConsumeModule) Run_before_module() {}

func (self *ConsumeModule) Run_module() {
	go self.Consumer()
	self.Start()

	self.Logger.Info("Consume module run")
}

func (self *ConsumeModule) Run_after_module() {}

func (self *ConsumeModule) Close_module() {

	self.Stop()

	if self.Hooker.ClearHook != nil {
		self.Hooker.ClearHook()
	}

	self.Logger.Info("Consume module close")
}

func (self *ConsumeModule) Start() {
	<- self.Started
	self.Status = core.GTHREAD_START
}

func (self *ConsumeModule) Stop() {

	if self.Status != core.GTHREAD_STOP {
		self.Stoped <- true
		self.Status = core.GTHREAD_STOP
	}
}

func (self *ConsumeModule) Consumer() {

	self.Started <- true

	for {
		self.Logger.Debug("Routine running")
		select {
		case <- self.Started:
			self.Logger.Info("Routine started")
		case <- self.Stoped:
			self.Logger.Info("Routine stoped")
			return
		default:
			// The sleep use of test.
			//time.Sleep(3 * time.Second)
			//self.Logger.Info(self.Hooker)
			if self.Hooker.ConsumerHook != nil{
				result := self.Hooker.ConsumerHook()
				if result != nil {
					self.Handler(result)
				}
			}
		}
	}
}

func (self *ConsumeModule) Handler(s *string) {
	if self.Hooker.HandlerHook != nil {
		self.Hooker.HandlerHook(s)
	}
}