package state

import (
	"context"
	"errors"

	"huaweiApi/pkg/rpc/protos"
)

func (service *Service) Ping(ctx context.Context, req *protos.PingRequest, reply *protos.PingReply) error {

	if req == nil || reply == nil {
		return errors.New("input parameters error")
	}

	reply.Message = req.Message
	return nil
}
