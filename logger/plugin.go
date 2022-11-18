package logger

import "github.com/galaxy-kit/galaxy-go/define"

var plugin = define.DefinePluginInterface[Logger]().ServicePluginInterface()

var Context = plugin.Context
