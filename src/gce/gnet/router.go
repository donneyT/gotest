package gnet

import "gce/giface"

type BaseRouter struct {}
func (br *BaseRouter) PreHandle(request giface.IRequest){}
func (br *BaseRouter)Handle(request giface.IRequest){}
func (br *BaseRouter)PostHandle(request giface.IRequest){}