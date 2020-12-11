package core

import (
	"fmt"
	"reflect"
)

type injector struct {
	values map[reflect.Type]reflect.Value
}

func InterfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic("Called inject.InterfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
	}

	return t
}

func NewInjector() *injector {
	return &injector{
		values: make(map[reflect.Type]reflect.Value),
	}
}

func (self *injector) Invoke(f interface{}) ([]reflect.Value, error) {
	t := reflect.TypeOf(f)

	var in = make([]reflect.Value, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		argType := t.In(i)
		val := self.Get(argType)
		if !val.IsValid() {
			return nil, fmt.Errorf("Value not found for type %v", argType)
		}

		in[i] = val
	}

	return reflect.ValueOf(f).Call(in), nil
}

func (self *injector) Apply(val interface{}) error {
	v := reflect.ValueOf(val)

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	t := v.Type()

	for i:=0; i < v.NumField(); i++ {
		f := v.Field(i)
		structField := t.Field(i)
		if f.CanSet() && (structField.Tag == "inject" || structField.Tag.Get("inject") != "") {
			ft := f.Type()
			v := self.Get(ft)
			if !v.IsValid() {
				return fmt.Errorf("Value not found for type %v", ft)
			}

			f.Set(v)
		}
	}

	return nil
}

func (self *injector) Get(t reflect.Type) reflect.Value {
	val := self.values[t]

	if val.IsValid() {
		return val
	}

	return val

}

func (self *injector) Map(val interface{}) {
	self.values[reflect.TypeOf(val)] = reflect.ValueOf(val)
}

func (self *injector) MapTo(val interface{}, ifacePtr interface{}) {
	self.values[InterfaceOf(ifacePtr)] = reflect.ValueOf(val)
}

func (self *injector) Set(t reflect.Type, val reflect.Value) {
	self.values[t] = val
}
