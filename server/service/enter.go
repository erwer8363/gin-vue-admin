package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	XkBBsServiceGroup   bbs.ServiceGroup
}
