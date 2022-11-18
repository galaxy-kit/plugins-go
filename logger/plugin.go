package logger

import "github.com/galaxy-kit/galaxy-go/define"

var plugin = define.DefinePluginInterface[Logger]().ServicePluginInterface()

var Get = plugin.Get

var TryGet = plugin.TryGet
