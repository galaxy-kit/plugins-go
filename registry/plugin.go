package registry

import (
	"github.com/galaxy-kit/galaxy-go/define"
)

var Plugin = define.DefinePluginInterface[Registry]().ServicePluginInterface()

var Get = Plugin.Get

var TryGet = Plugin.TryGet
