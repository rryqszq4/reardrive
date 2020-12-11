package modules

import (
	"msgback/conf_struct"
	"msgback/core"
	"github.com/apache/thrift/lib/go/thrift"
	"time"
)

type ThriftServerModuleImpl interface {

	Start()

	Driver()

	Stop()
}

type ThriftServerModule struct {
	//
	core.ModuleImpl

	ThriftServerModuleImpl

	//
	ProtocolFactory thrift.TProtocolFactory

	TransportFactory thrift.TTransportFactory

	Addr string

	Protocol string

	Secure *bool

	Transport thrift.TServerTransport

	Processor thrift.TProcessor

	Server *thrift.TSimpleServer

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
		GetProcessor func() thrift.TProcessor
	}
}

func (self *ThriftServerModule) Init_module(cycle *core.Cycle) {
	self.Logger = core.GetLogger()
	self.Recorder = core.GetRecorder()

	self.Started = make(chan bool)
	self.Stoped = make(chan bool)

	self.Conf = core.GetInjectModule(core.Conf{}).(core.ConfImpl).GetAll()

	self.Logger.Info("Thrift server module init")
}

func (self *ThriftServerModule) Run_before_module() {}

func (self *ThriftServerModule) Run_module() {
	go self.Driver()
	self.Start()
}

func (self *ThriftServerModule) Run_after_module() {}

func (self *ThriftServerModule) Close_module() {

	self.Server.Stop()
	self.Stop()
}

func (self *ThriftServerModule) Start() {
	<- self.Started
	self.Status = core.GTHREAD_START
}

func (self *ThriftServerModule) Stop() {

	if self.Status != core.GTHREAD_STOP {
		self.Stoped <- true
		self.Status = core.GTHREAD_STOP
	}
}

func (self *ThriftServerModule) Driver() {

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

			if self.Hooker.GetProcessor != nil {
				self.TransportFactory = thrift.NewTTransportFactory()
				self.Transport, _ = thrift.NewTServerSocket(self.Addr)
				self.ProtocolFactory = self.getProtocolFactory(&self.Protocol)

				self.Processor = self.Hooker.GetProcessor()

				self.Server = thrift.NewTSimpleServer4(self.Processor, self.Transport, self.TransportFactory, self.ProtocolFactory)

				self.Logger.Info("Thrift server serving...")
				self.Server.Serve()
			}

			time.Sleep(3* time.Second)
		}
	}
}

func (self *ThriftServerModule) getProtocolFactory(protocol *string) thrift.TProtocolFactory {
	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	}

	return protocolFactory
}

