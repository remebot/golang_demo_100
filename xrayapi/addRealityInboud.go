package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/xtls/xray-core/app/proxyman"
	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/proxy/vless"
	"github.com/xtls/xray-core/proxy/vless/inbound"
	"github.com/xtls/xray-core/transport/internet"
	"github.com/xtls/xray-core/transport/internet/reality"
	_ "github.com/xtls/xray-core/transport/internet/tcp"
	"golang.org/x/crypto/curve25519"
)

// 添加一个 VLESS-Vision-REALITY 入站（保留的默认实现）
// func addVlessRealityInbound(client command.HandlerServiceClient) error {
// 	p := RealityInboundParams{
// 		Tag:         "VLESS-Vision-REALITY",
// 		Port:        443,
// 		ListenAnyIP: true,
// 		UserUUID:    "b8eb278f-8685-4a7a-a7ab-3cad172c230a",
// 		UserEmail:   "vless@xray",
// 		UserFlow:    "xtls-rprx-vision",
// 		Dest:        "dl.google.com:443",
// 		ServerNames: []string{"dl.google.com"},
// 		ShowReality: false,
// 		// PrivateKey/ShortIds 留空则自动生成
// 	}
// 	_, _, _, err := addVlessRealityInboundWithParams(client, p)
// 	return err
// }

// 参数结构体：用于参数化创建 VLESS-Vision-REALITY 入站
type RealityInboundParams struct {
	Tag         string
	Port        uint16
	ListenAnyIP bool
	ListenIP    string

	UserUUID  string
	UserEmail string
	UserFlow  string

	Dest        string
	ServerNames []string
	ShowReality bool

	// 二选一输入：提供 PrivateKey 或 PrivateKeyHex；都为空时自动生成
	PrivateKey    []byte
	PrivateKeyHex string
	// 二选一输入：提供 ShortIds 或 ShortIdsHex；都为空时自动生成一个8位hex
	ShortIds    [][]byte
	ShortIdsHex []string
}

// 生成 X25519 私钥（32字节，已 clamp）
func generateX25519PrivateKey() ([]byte, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	// clamp per RFC7748
	b[0] &= 248
	b[31] &= 127
	b[31] |= 64
	return b, nil
}

// 从私钥推导 X25519 公钥（32字节）
func deriveX25519PublicKey(private []byte) ([]byte, error) {
	return curve25519.X25519(private, curve25519.Basepoint)
}

// 生成一个 8位 hex 的 shortId
func generateShortId8Hex() ([]byte, error) {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}
	enc := make([]byte, hex.EncodedLen(len(buf)))
	hex.Encode(enc, buf)
	return enc, nil // 返回的是ASCII-hex字节，例如 []byte("a1b2c3d4")
}

// 参数化添加入站
// 返回：使用的 PrivateKey、PublicKey、ShortIds
func addVlessRealityInboundWithParams(client command.HandlerServiceClient, p RealityInboundParams) ([]byte, []byte, [][]byte, error) {
	// 合法化默认值
	if p.Tag == "" {
		p.Tag = "VLESS-Vision-REALITY"
	}
	if p.Port == 0 {
		p.Port = 443
	}
	if p.UserFlow == "" {
		p.UserFlow = "xtls-rprx-vision"
	}
	if p.UserEmail == "" {
		p.UserEmail = "vless@xray"
	}
	if p.Dest == "" {
		p.Dest = "dl.google.com:443"
	}
	if len(p.ServerNames) == 0 {
		p.ServerNames = []string{"dl.google.com"}
	}

	// 准备私钥
	var privKey []byte
	if len(p.PrivateKey) > 0 {
		privKey = p.PrivateKey
	} else if p.PrivateKeyHex != "" {
		decoded, err := hex.DecodeString(p.PrivateKeyHex)
		if err != nil {
			return nil, nil, nil, err
		}
		privKey = decoded
	} else {
		var err error
		privKey, err = generateX25519PrivateKey()
		if err != nil {
			return nil, nil, nil, err
		}
	}

	// 推导公钥
	pubKey, err := deriveX25519PublicKey(privKey)
	if err != nil {
		return nil, nil, nil, err
	}

	// 准备 ShortIds
	var shortIds [][]byte
	if len(p.ShortIds) > 0 {
		shortIds = p.ShortIds
	} else if len(p.ShortIdsHex) > 0 {
		for _, hx := range p.ShortIdsHex {
			b, err := hex.DecodeString(hx)
			if err != nil {
				return nil, nil, nil, err
			}
			shortIds = append(shortIds, b)
		}
	} else {
		b, err := generateShortId8Hex()
		if err != nil {
			return nil, nil, nil, err
		}
		shortIds = [][]byte{b}
	}

	// VLESS Inbound 配置
	vlessCfg := &inbound.Config{
		Clients: []*protocol.User{
			{
				Level: 0,
				Email: p.UserEmail,
				Account: serial.ToTypedMessage(&vless.Account{
					Id:   p.UserUUID,
					Flow: p.UserFlow,
				}),
			},
		},
		Decryption: "none",
	}

	// 传输与 REALITY 安全层配置
	stream := &internet.StreamConfig{
		ProtocolName: "tcp",
		SecurityType: serial.GetMessageType(&reality.Config{}),
		SecuritySettings: []*serial.TypedMessage{
			serial.ToTypedMessage(&reality.Config{
				Show:        p.ShowReality,
				Dest:        p.Dest,
				Xver:        0,
				ServerNames: p.ServerNames,
				PrivateKey:  privKey,
				ShortIds:    shortIds,
			}),
		},
	}

	// 监听设置
	receiver := &proxyman.ReceiverConfig{
		PortList: &net.PortList{Range: []*net.PortRange{net.SinglePortRange(net.Port(p.Port))}},
		SniffingSettings: &proxyman.SniffingConfig{
			Enabled:             true,
			DestinationOverride: []string{"http", "tls"},
		},
		StreamSettings: stream,
	}
	if p.ListenAnyIP {
		receiver.Listen = net.NewIPOrDomain(net.AnyIP)
	} else if p.ListenIP != "" {
		receiver.Listen = net.NewIPOrDomain(net.ParseAddress(p.ListenIP))
	} else {
		receiver.Listen = net.NewIPOrDomain(net.AnyIP)
	}

	_, err = client.AddInbound(context.Background(), &command.AddInboundRequest{
		Inbound: &core.InboundHandlerConfig{
			Tag:              p.Tag,
			ReceiverSettings: serial.ToTypedMessage(receiver),
			ProxySettings:    serial.ToTypedMessage(vlessCfg),
		},
	})
	if err != nil {
		return nil, nil, nil, err
	}
	return privKey, pubKey, shortIds, nil
}

// 在当前文件中添加一个用于测试的 main 函数
// 注意：此 main 仅用于本地测试 API 是否可调用
// 运行该文件会尝试通过 gRPC 调用正在运行的 Xray 实例的 API
// 请先确保 Xray 已启用 API 并监听到对应地址与端口
func main() {
	var (
		xrayCtl *XrayController
		cfg     = &BaseConfig{
			APIAddress: "34.94.216.7",
			APIPort:    32768,
		}
	)

	xrayCtl = new(XrayController)
	if err := xrayCtl.Init(cfg); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer xrayCtl.CmdConn.Close()

	// 使用参数化版本：未提供 PrivateKey/ShortIds 时会自动生成
	params := RealityInboundParams{
		Tag:         "VLESS-Vision-REALITY",
		Port:        443,
		ListenAnyIP: true,
		UserUUID:    "b8eb278f-8685-4a7a-a7ab-3cad172c230a",
		UserEmail:   "vless@xray",
		UserFlow:    "xtls-rprx-vision",
		Dest:        "dl.google.com:443",
		ServerNames: []string{"dl.google.com"},
		ShowReality: false,
	}

	priv, pub, sids, err := addVlessRealityInboundWithParams(xrayCtl.HsClient, params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success")
	// 打印客户端连接所需参数（示例：v2rayN/xray-core 客户端）
	fmt.Println("\n客户端参数：")
	fmt.Printf("server: %s\n", "<你的服务器IP>")
	fmt.Printf("port: %d\n", params.Port)
	fmt.Printf("uuid: %s\n", params.UserUUID)
	fmt.Printf("flow: %s\n", params.UserFlow)
	fmt.Printf("sni(serverName): %s\n", params.ServerNames[0])
	fmt.Printf("public-key: %s\n", hex.EncodeToString(pub))
	if len(sids) > 0 {
		fmt.Printf("short-id: %s\n", string(sids[0]))
	}
	_ = priv // 若需要也可打印私钥：hex.EncodeToString(priv)
}
