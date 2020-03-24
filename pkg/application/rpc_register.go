package application

import (
	"github.com/sirupsen/logrus"
	"huaweiApi/pkg/rpc/protos"
	"huaweiApi/pkg/rpc/state"
)

func (a *app) setRegisterRPCControllers() {
	var err error
	err = protos.RegisterStatusControllerHandler(a.rpcService.Server(), new(state.Service))
	if err != nil {
		logrus.WithField("error", err).Errorf("register state error")
	}
}
