package token

import "github.com/armoniax/eos-go"

func init() {
	eos.RegisterAction(AN("amax.token"), ActN("transfer"), Transfer{})
	eos.RegisterAction(AN("amax.token"), ActN("issue"), Issue{})
	eos.RegisterAction(AN("amax.token"), ActN("create"), Create{})
}
