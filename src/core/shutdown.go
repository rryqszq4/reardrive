package core

type Shutdown struct {
	Name string
}

func (self *Shutdown) Init_module(cycle *Cycle) {
	GetLogger().Info("shutdown init")
}

func (self *Shutdown) Run_before_module() {}

func (self *Shutdown) Run_module() {
	GetLogger().Info("shutdown run")
}

func (self *Shutdown) Run_after_module() {}

func (self *Shutdown) Close_module() {
	GetLogger().Info("shutdown close")
}

