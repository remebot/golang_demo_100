package main

import (
	"context"
	"fmt"
	"github.com/xtls/xray-core/app/proxyman/command"
)

/*
获取 Tag 下用户数据，如下
users:{email:"love@xray.com" account:{type:"xray.proxy.vmess.Account"  value:"\n$10354ac4-9ec1-4864-ba3e-f5fd35869ef8\x1a\x02\x08\x04"}}
Email 为空时返回所有用户
*/

func getInboundUsers(client command.HandlerServiceClient) (users *command.GetInboundUserResponse, err error) {
	users, err = client.GetInboundUsers(context.Background(), &command.GetInboundUserRequest{
		Tag: "VLESS-Vision-REALITY",
		//Email:"love@xray.com",
	})

	return
}

func main_getUsers() {
	// 先指定 API 端口和地址
	var (
		xrayCtl *XrayController
		cfg     = &BaseConfig{
			APIAddress: "34.94.216.7",
			APIPort:    32768,
		}
	)

	xrayCtl = new(XrayController)
	err := xrayCtl.Init(cfg)
	defer xrayCtl.CmdConn.Close()
	if err != nil {
		fmt.Println(err)
	}

	// 获取用户信息
	fmt.Println("获取用户信息：")
	users, err := getInboundUsers(xrayCtl.HsClient)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}
