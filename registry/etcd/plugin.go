package etcd

import (
	"github.com/galaxy-kit/galaxy-go/define"
	"github.com/galaxy-kit/plugins-go/registry"
)

var plugin = define.DefinePlugin[registry.Registry, WithEtcdOption]().ServicePlugin(newRegistry)

var InstallTo = plugin.InstallTo

var UninstallFrom = plugin.UninstallFrom
