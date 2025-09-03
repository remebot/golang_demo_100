package main

import "fmt"

func main() {
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
	users, err := getInboundUsers(xrayCtl.HsClient)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}
