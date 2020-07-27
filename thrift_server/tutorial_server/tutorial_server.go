package tutorial_server

import (
	"code.sohuno.com/apollo/msgback/core"
	"code.sohuno.com/apollo/msgback/core/modules"
	"code.sohuno.com/apollo/msgback/thrift_server/tutorial"
	"github.com/apache/thrift/lib/go/thrift"
)

type TutorialServer struct {
	modules.ThriftServerModule

}

func (self *TutorialServer) Run_before_module() {
	self.init()
}


func (self *TutorialServer) init() {

	_thrift := core.GetHotConfig().GetMap("thrift_server")

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

