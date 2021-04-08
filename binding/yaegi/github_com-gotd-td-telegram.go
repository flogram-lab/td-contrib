// Code generated by 'yaegi extract github.com/gotd/td/telegram'. DO NOT EDIT.

package yaegi

import (
	"context"
	"github.com/gotd/td/bin"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/transport"
	"go/constant"
	"go/token"
	"net"
	"reflect"
)

func init() {
	Symbols["github.com/gotd/td/telegram"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AddrProduction":         reflect.ValueOf(constant.MakeFromLiteral("\"149.154.167.50:443\"", token.STRING, 0)),
		"AddrTest":               reflect.ValueOf(constant.MakeFromLiteral("\"149.154.167.40:443\"", token.STRING, 0)),
		"AsFloodWait":            reflect.ValueOf(telegram.AsFloodWait),
		"BotFromEnvironment":     reflect.ValueOf(telegram.BotFromEnvironment),
		"ClientFromEnvironment":  reflect.ValueOf(telegram.ClientFromEnvironment),
		"CodeOnlyAuth":           reflect.ValueOf(telegram.CodeOnlyAuth),
		"ConstantAuth":           reflect.ValueOf(telegram.ConstantAuth),
		"EnvAuth":                reflect.ValueOf(telegram.EnvAuth),
		"ErrFloodWait":           reflect.ValueOf(constant.MakeFromLiteral("\"FLOOD_WAIT\"", token.STRING, 0)),
		"ErrPasswordAuthNeeded":  reflect.ValueOf(&telegram.ErrPasswordAuthNeeded).Elem(),
		"ErrPasswordNotProvided": reflect.ValueOf(&telegram.ErrPasswordNotProvided).Elem(),
		"NewAuth":                reflect.ValueOf(telegram.NewAuth),
		"NewClient":              reflect.ValueOf(telegram.NewClient),
		"OptionsFromEnvironment": reflect.ValueOf(telegram.OptionsFromEnvironment),
		"Port":                   reflect.ValueOf(constant.MakeFromLiteral("443", token.INT, 0)),
		"RunUntilCanceled":       reflect.ValueOf(telegram.RunUntilCanceled),
		"TestAppHash":            reflect.ValueOf(constant.MakeFromLiteral("\"344583e45741c457fe1862106095a5eb\"", token.STRING, 0)),
		"TestAppID":              reflect.ValueOf(constant.MakeFromLiteral("17349", token.INT, 0)),
		"TestAuth":               reflect.ValueOf(telegram.TestAuth),
		"TestClient":             reflect.ValueOf(telegram.TestClient),

		// type definitions
		"AuthFlow":              reflect.ValueOf((*telegram.AuthFlow)(nil)),
		"AuthFlowClient":        reflect.ValueOf((*telegram.AuthFlowClient)(nil)),
		"AuthStatus":            reflect.ValueOf((*telegram.AuthStatus)(nil)),
		"Client":                reflect.ValueOf((*telegram.Client)(nil)),
		"CloseInvoker":          reflect.ValueOf((*telegram.CloseInvoker)(nil)),
		"CodeAuthenticator":     reflect.ValueOf((*telegram.CodeAuthenticator)(nil)),
		"CodeAuthenticatorFunc": reflect.ValueOf((*telegram.CodeAuthenticatorFunc)(nil)),
		"DeviceConfig":          reflect.ValueOf((*telegram.DeviceConfig)(nil)),
		"Error":                 reflect.ValueOf((*telegram.Error)(nil)),
		"FileSessionStorage":    reflect.ValueOf((*telegram.FileSessionStorage)(nil)),
		"Options":               reflect.ValueOf((*telegram.Options)(nil)),
		"SendCodeOptions":       reflect.ValueOf((*telegram.SendCodeOptions)(nil)),
		"SessionStorage":        reflect.ValueOf((*telegram.SessionStorage)(nil)),
		"SignUp":                reflect.ValueOf((*telegram.SignUp)(nil)),
		"SignUpRequired":        reflect.ValueOf((*telegram.SignUpRequired)(nil)),
		"Transport":             reflect.ValueOf((*telegram.Transport)(nil)),
		"UpdateHandler":         reflect.ValueOf((*telegram.UpdateHandler)(nil)),
		"UserAuthenticator":     reflect.ValueOf((*telegram.UserAuthenticator)(nil)),
		"UserInfo":              reflect.ValueOf((*telegram.UserInfo)(nil)),

		// interface wrapper definitions
		"_AuthFlowClient":    reflect.ValueOf((*_github_com_gotd_td_telegram_AuthFlowClient)(nil)),
		"_CloseInvoker":      reflect.ValueOf((*_github_com_gotd_td_telegram_CloseInvoker)(nil)),
		"_CodeAuthenticator": reflect.ValueOf((*_github_com_gotd_td_telegram_CodeAuthenticator)(nil)),
		"_SessionStorage":    reflect.ValueOf((*_github_com_gotd_td_telegram_SessionStorage)(nil)),
		"_Transport":         reflect.ValueOf((*_github_com_gotd_td_telegram_Transport)(nil)),
		"_UpdateHandler":     reflect.ValueOf((*_github_com_gotd_td_telegram_UpdateHandler)(nil)),
		"_UserAuthenticator": reflect.ValueOf((*_github_com_gotd_td_telegram_UserAuthenticator)(nil)),
	}
}

// _github_com_gotd_td_telegram_AuthFlowClient is an interface wrapper for AuthFlowClient type
type _github_com_gotd_td_telegram_AuthFlowClient struct {
	WAuthPassword func(ctx context.Context, password string) error
	WAuthSendCode func(ctx context.Context, phone string, options telegram.SendCodeOptions) (codeHash string, err error)
	WAuthSignIn   func(ctx context.Context, phone string, code string, codeHash string) error
	WAuthSignUp   func(ctx context.Context, s telegram.SignUp) error
}

func (W _github_com_gotd_td_telegram_AuthFlowClient) AuthPassword(ctx context.Context, password string) error {
	return W.WAuthPassword(ctx, password)
}
func (W _github_com_gotd_td_telegram_AuthFlowClient) AuthSendCode(ctx context.Context, phone string, options telegram.SendCodeOptions) (codeHash string, err error) {
	return W.WAuthSendCode(ctx, phone, options)
}
func (W _github_com_gotd_td_telegram_AuthFlowClient) AuthSignIn(ctx context.Context, phone string, code string, codeHash string) error {
	return W.WAuthSignIn(ctx, phone, code, codeHash)
}
func (W _github_com_gotd_td_telegram_AuthFlowClient) AuthSignUp(ctx context.Context, s telegram.SignUp) error {
	return W.WAuthSignUp(ctx, s)
}

// _github_com_gotd_td_telegram_CloseInvoker is an interface wrapper for CloseInvoker type
type _github_com_gotd_td_telegram_CloseInvoker struct {
	WClose     func() error
	WInvokeRaw func(ctx context.Context, input bin.Encoder, output bin.Decoder) error
}

func (W _github_com_gotd_td_telegram_CloseInvoker) Close() error { return W.WClose() }
func (W _github_com_gotd_td_telegram_CloseInvoker) InvokeRaw(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	return W.WInvokeRaw(ctx, input, output)
}

// _github_com_gotd_td_telegram_CodeAuthenticator is an interface wrapper for CodeAuthenticator type
type _github_com_gotd_td_telegram_CodeAuthenticator struct {
	WCode func(ctx context.Context) (string, error)
}

func (W _github_com_gotd_td_telegram_CodeAuthenticator) Code(ctx context.Context) (string, error) {
	return W.WCode(ctx)
}

// _github_com_gotd_td_telegram_SessionStorage is an interface wrapper for SessionStorage type
type _github_com_gotd_td_telegram_SessionStorage struct {
	WLoadSession  func(ctx context.Context) ([]byte, error)
	WStoreSession func(ctx context.Context, data []byte) error
}

func (W _github_com_gotd_td_telegram_SessionStorage) LoadSession(ctx context.Context) ([]byte, error) {
	return W.WLoadSession(ctx)
}
func (W _github_com_gotd_td_telegram_SessionStorage) StoreSession(ctx context.Context, data []byte) error {
	return W.WStoreSession(ctx, data)
}

// _github_com_gotd_td_telegram_Transport is an interface wrapper for Transport type
type _github_com_gotd_td_telegram_Transport struct {
	WCodec     func() transport.Codec
	WHandshake func(conn net.Conn) (transport.Conn, error)
}

func (W _github_com_gotd_td_telegram_Transport) Codec() transport.Codec { return W.WCodec() }
func (W _github_com_gotd_td_telegram_Transport) Handshake(conn net.Conn) (transport.Conn, error) {
	return W.WHandshake(conn)
}

// _github_com_gotd_td_telegram_UpdateHandler is an interface wrapper for UpdateHandler type
type _github_com_gotd_td_telegram_UpdateHandler struct {
	WHandle      func(ctx context.Context, u *tg.Updates) error
	WHandleShort func(ctx context.Context, u *tg.UpdateShort) error
}

func (W _github_com_gotd_td_telegram_UpdateHandler) Handle(ctx context.Context, u *tg.Updates) error {
	return W.WHandle(ctx, u)
}
func (W _github_com_gotd_td_telegram_UpdateHandler) HandleShort(ctx context.Context, u *tg.UpdateShort) error {
	return W.WHandleShort(ctx, u)
}

// _github_com_gotd_td_telegram_UserAuthenticator is an interface wrapper for UserAuthenticator type
type _github_com_gotd_td_telegram_UserAuthenticator struct {
	WAcceptTermsOfService func(ctx context.Context, tos tg.HelpTermsOfService) error
	WCode                 func(ctx context.Context) (string, error)
	WPassword             func(ctx context.Context) (string, error)
	WPhone                func(ctx context.Context) (string, error)
	WSignUp               func(ctx context.Context) (telegram.UserInfo, error)
}

func (W _github_com_gotd_td_telegram_UserAuthenticator) AcceptTermsOfService(ctx context.Context, tos tg.HelpTermsOfService) error {
	return W.WAcceptTermsOfService(ctx, tos)
}
func (W _github_com_gotd_td_telegram_UserAuthenticator) Code(ctx context.Context) (string, error) {
	return W.WCode(ctx)
}
func (W _github_com_gotd_td_telegram_UserAuthenticator) Password(ctx context.Context) (string, error) {
	return W.WPassword(ctx)
}
func (W _github_com_gotd_td_telegram_UserAuthenticator) Phone(ctx context.Context) (string, error) {
	return W.WPhone(ctx)
}
func (W _github_com_gotd_td_telegram_UserAuthenticator) SignUp(ctx context.Context) (telegram.UserInfo, error) {
	return W.WSignUp(ctx)
}
