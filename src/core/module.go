package core

type ModuleImpl interface {

	Init_module()
	Run_before_module()
	Run_module()
	Run_after_module()
	Close_module()
}

type ModuleFrame struct {
	Name string
	Category int
	Object interface{}
	Init func()
	Handler func()
	Clear func()
}

func (self *ModuleFrame) Init_module() {
	switch true {
	case self.Init != nil:
		self.Init()
	case self.Object != nil:
		self.Object.(ModuleImpl).Init_module()
	}
}

func (self *ModuleFrame) Run_before_module()  {
	if self.Object != nil {
		self.Object.(ModuleImpl).Run_before_module()
		return
	}
}

func (self *ModuleFrame) Run_module()  {
	if self.Handler != nil {
		self.Handler()
		return
	}

	if self.Object != nil {
		self.Object.(ModuleImpl).Run_module()
		return
	}
}

func (self *ModuleFrame) Run_after_module()  {
	if self.Object != nil {
		self.Object.(ModuleImpl).Run_after_module()
		return
	}
}

func (self *ModuleFrame) Close_module() {
	if self.Clear != nil {
		self.Clear()
		return
	}

	if self.Object != nil {
		self.Object.(ModuleImpl).Close_module()
		return
	}
}

