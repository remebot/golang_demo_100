package main

import (
	"context"
	"fmt"
	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/proxy/vless"
)

func addVlessUser(client command.HandlerServiceClient, user *UserInfo) error {
	_, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: user.InTag,
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Level: user.Level,
				Email: user.Email,
				Account: serial.ToTypedMessage(&vless.Account{
					Id:   user.Uuid,
					Flow: "xtls-rprx-vision",
				}),
			},
		}),
	})
	return err
}

func main_addRealityUser() {
	fmt.Println("即将添加用户")
	var (
		xrayCtl *XrayController
		cfg     = &BaseConfig{
			APIAddress: "34.94.216.7",
			APIPort:    32768,
		}
		user = UserInfo{
			Uuid:  "fac18615-f70a-4f78-98dd-6b1202b560a8",
			Level: 0,
			InTag: "VLESS-Vision-REALITY",
			Email: "add_by_api@xray.com",
		}
	)
	xrayCtl = new(XrayController)
	err := xrayCtl.Init(cfg)
	defer xrayCtl.CmdConn.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = addVlessUser(xrayCtl.HsClient, &user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("success")
}
