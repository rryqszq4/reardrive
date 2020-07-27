package boostrap

import (
	"reflect"
	"reardrive/src/core"
)

var (
	GlobalModuleString = []string {
		"core.Startup",
		"core.Conf",
		"core.HotConfig",
		"tutorial_server.TutorialServer",
	}

	GlobalStructs []interface{}

	GlobalModuleStruct map[string]reflect.Type

	GlobalModule map[string]*core.ModuleFrame
)