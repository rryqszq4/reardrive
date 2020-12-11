package tutorial_server

import (
	"reardrive/core"
	"reardrive/core/modules"
	"reardrive/thrift_server/library/tutorial"
	"github.com/apache/thrift/lib/go/thrift"
)

type TutorialServer struct {
	modules.ThriftServerModule

}

func (self *TutorialServer) Run_before_module() {
	self.init()
}


func (self *TutorialServer) init() {

	_thrift := core.GetInjectModule(core.HotConfig{}).(*core.HotConfig).GetMap("thrift_server")

	_addr := _thrift["addr"]
	self.Addr = _addr.(string)

	_protocol := _thrift["protocol"]
	self.Protocol = _protocol.(string)

	self.Hooker.GetProcessor = self.getProcessor
}

func (self *TutorialServer) getProcessor() thrift.TProcessor {
	handler := NewCalculatorHandler()

	return tutorial.NewCalculatorProcessor(handler)
}

