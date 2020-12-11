package core

type Bootstrap struct {
	Name string
}

func (self *Bootstrap) Init_module(cycle *Cycle) {
	GetLogger().Info("bootstrap init")
}

func (self *Bootstrap) Run_before_module() {}

func (self *Bootstrap) Run_module() {
	GetLogger().Info("bootstrap run")
}

func (self *Bootstrap) Run_after_module() {}

func (self *Bootstrap) Close_module() {
	GetLogger().Info("bootstrap close")
}
