package logger

import "github.com/galaxy-kit/galaxy-go/define"

var Plugin = define.DefinePluginInterface[Logger]().ServicePluginInterface()

var Get = Plugin.Get

var TryGet = Plugin.TryGet
