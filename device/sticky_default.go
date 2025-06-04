//go:build !linux

package device

import (
	"github.com/chickencoding123/wireguard-go-nanovms/wireguard/conn"
	"github.com/chickencoding123/wireguard-go-nanovms/wireguard/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
