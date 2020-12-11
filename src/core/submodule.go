package core

type Submodule struct {
	Name string
}

func (self *Submodule) Init_module() {
	GetLogger().Info("Submodule init")
}

func (self *Submodule) Run_before_module() {}

func (self *Submodule) Run_module() {
	GetLogger().Info("Submodule run")
}

func (self *Submodule) Run_after_module() {}

func (self *Submodule) Close_module() {
	GetLogger().Info("Submodule close")
}