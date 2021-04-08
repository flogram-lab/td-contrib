// Code generated by 'yaegi extract github.com/gotd/td/telegram/dcs'. DO NOT EDIT.

package yaegi

import (
	"context"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/transport"
	"net"
	"reflect"
)

func init() {
	Symbols["github.com/gotd/td/telegram/dcs"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"FindDCs":         reflect.ValueOf(dcs.FindDCs),
		"FindPrimaryDCs":  reflect.ValueOf(dcs.FindPrimaryDCs),
		"MTProxyResolver": reflect.ValueOf(dcs.MTProxyResolver),
		"PlainResolver":   reflect.ValueOf(dcs.PlainResolver),
		"ProdDCs":         reflect.ValueOf(dcs.ProdDCs),
		"StagingDCs":      reflect.ValueOf(dcs.StagingDCs),

		// type definitions
		"MTProxyOptions": reflect.ValueOf((*dcs.MTProxyOptions)(nil)),
		"PlainOptions":   reflect.ValueOf((*dcs.PlainOptions)(nil)),
		"Resolver":       reflect.ValueOf((*dcs.Resolver)(nil)),
		"Transport":      reflect.ValueOf((*dcs.Transport)(nil)),

		// interface wrapper definitions
		"_Resolver":  reflect.ValueOf((*_github_com_gotd_td_telegram_dcs_Resolver)(nil)),
		"_Transport": reflect.ValueOf((*_github_com_gotd_td_telegram_dcs_Transport)(nil)),
	}
}

// _github_com_gotd_td_telegram_dcs_Resolver is an interface wrapper for Resolver type
type _github_com_gotd_td_telegram_dcs_Resolver struct {
	WCDN       func(ctx context.Context, dc int, dcOptions []tg.DCOption) (transport.Conn, error)
	WMediaOnly func(ctx context.Context, dc int, dcOptions []tg.DCOption) (transport.Conn, error)
	WPrimary   func(ctx context.Context, dc int, dcOptions []tg.DCOption) (transport.Conn, error)
}

func (W _github_com_gotd_td_telegram_dcs_Resolver) CDN(ctx context.Context, dc int, dcOptions []tg.DCOption) (transport.Conn, error) {
	return W.WCDN(ctx, dc, dcOptions)
}
func (W _github_com_gotd_td_telegram_dcs_Resolver) MediaOnly(ctx context.Context, dc int, dcOptions []tg.DCOption) (transport.Conn, error) {
	return W.WMediaOnly(ctx, dc, dcOptions)
}
func (W _github_com_gotd_td_telegram_dcs_Resolver) Primary(ctx context.Context, dc int, dcOptions []tg.DCOption) (transport.Conn, error) {
	return W.WPrimary(ctx, dc, dcOptions)
}

// _github_com_gotd_td_telegram_dcs_Transport is an interface wrapper for Transport type
type _github_com_gotd_td_telegram_dcs_Transport struct {
	WCodec     func() transport.Codec
	WHandshake func(conn net.Conn) (transport.Conn, error)
}

func (W _github_com_gotd_td_telegram_dcs_Transport) Codec() transport.Codec { return W.WCodec() }
func (W _github_com_gotd_td_telegram_dcs_Transport) Handshake(conn net.Conn) (transport.Conn, error) {
	return W.WHandshake(conn)
}
