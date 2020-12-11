package modules

import (
	"msgback/conf_struct"
	"msgback/core"
)

type HttpServerModuleImpl interface {

	Start()

	Driver()

	Stop()

}

type HttpServerModule struct {
	//
	core.ModuleImpl

	HttpServerModuleImpl

	Addr string

	Name string

	//
	Started chan bool

	Stoped chan bool

	Status int

	//
	Conf conf_struct.CoreConf

	Logger core.Logger

	Recorder core.Logger

	//
	Hooker struct{
		Listener func()
	}

}

func (self *HttpServerModule) Init_module(cycle *core.Cycle) {
	self.Logger = core.GetLogger()
	self.Recorder = core.GetRecorder()

	self.Started = make(chan bool)
	self.Stoped = make(chan bool)

	self.Conf = core.GetInjectModule(core.Conf{}).(core.ConfImpl).GetAll()

	self.Logger.Info("Http server module init")
}

func (self *HttpServerModule) Run_before_module() {}

func (self *HttpServerModule) Run_module() {
	go self.Driver()
	self.Start()
}

func (self *HttpServerModule) Run_after_module() {}

func (self *HttpServerModule) Close_module() {

	self.Stop()
}

func (self *HttpServerModule) Start() {
	<- self.Started
	self.Status = core.GTHREAD_START
}

func (self *HttpServerModule) Stop() {

	if self.Status != core.GTHREAD_STOP {
		self.Stoped <- true
		self.Status = core.GTHREAD_STOP
	}
}

func (self *HttpServerModule) Driver() {

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

			if self.Hooker.Listener != nil {

				self.Hooker.Listener()
			}

			//time.Sleep(3* time.Second)
		}
	}
}

