package main

import (
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry/etcd"
	"golang.org/x/net/context"

	"huaweiApi/pkg/constants"
	ptest "huaweiApi/pkg/rpc/protos"
)

func runClient(service micro.Service) {
	testmymicro := ptest.NewMiniprogramService(constants.RPCServiceName, service.Client())
	rsp, err := testmymicro.GetMiniprogramByShopId(context.TODO(), &ptest.GetMiniprogramRequest{ShopId: 123})

	if err != nil {
		fmt.Println(err)
		return
	}
	// Print response
	fmt.Println("appId", rsp.AppId)
	fmt.Println("secret", rsp.Secret)
}
func main() {

	service := micro.NewService(
		micro.Registry(etcd.NewRegistry()),
	)
	service.Init()
	runClient(service)
}
