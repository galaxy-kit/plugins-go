package registry

import (
	"github.com/galaxy-kit/galaxy-go/define"
)

var plugin = define.DefinePluginInterface[Registry]().ServicePluginInterface()

var Context = plugin.Context
