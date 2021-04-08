// Code generated by 'yaegi extract github.com/gotd/td/telegram/uploader'. DO NOT EDIT.

package yaegi

import (
	"context"
	"github.com/gotd/td/telegram/uploader"
	"github.com/gotd/td/tg"
	"go/constant"
	"go/token"
	"io/fs"
	"reflect"
)

func init() {
	Symbols["github.com/gotd/td/telegram/uploader"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"MaximumPartSize": reflect.ValueOf(constant.MakeFromLiteral("524288", token.INT, 0)),
		"NewUpload":       reflect.ValueOf(uploader.NewUpload),
		"NewUploader":     reflect.ValueOf(uploader.NewUploader),

		// type definitions
		"Client":        reflect.ValueOf((*uploader.Client)(nil)),
		"File":          reflect.ValueOf((*uploader.File)(nil)),
		"Progress":      reflect.ValueOf((*uploader.Progress)(nil)),
		"ProgressState": reflect.ValueOf((*uploader.ProgressState)(nil)),
		"Upload":        reflect.ValueOf((*uploader.Upload)(nil)),
		"Uploader":      reflect.ValueOf((*uploader.Uploader)(nil)),

		// interface wrapper definitions
		"_Client":   reflect.ValueOf((*_github_com_gotd_td_telegram_uploader_Client)(nil)),
		"_File":     reflect.ValueOf((*_github_com_gotd_td_telegram_uploader_File)(nil)),
		"_Progress": reflect.ValueOf((*_github_com_gotd_td_telegram_uploader_Progress)(nil)),
	}
}

// _github_com_gotd_td_telegram_uploader_Client is an interface wrapper for Client type
type _github_com_gotd_td_telegram_uploader_Client struct {
	WUploadSaveBigFilePart func(ctx context.Context, request *tg.UploadSaveBigFilePartRequest) (bool, error)
	WUploadSaveFilePart    func(ctx context.Context, request *tg.UploadSaveFilePartRequest) (bool, error)
}

func (W _github_com_gotd_td_telegram_uploader_Client) UploadSaveBigFilePart(ctx context.Context, request *tg.UploadSaveBigFilePartRequest) (bool, error) {
	return W.WUploadSaveBigFilePart(ctx, request)
}
func (W _github_com_gotd_td_telegram_uploader_Client) UploadSaveFilePart(ctx context.Context, request *tg.UploadSaveFilePartRequest) (bool, error) {
	return W.WUploadSaveFilePart(ctx, request)
}

// _github_com_gotd_td_telegram_uploader_File is an interface wrapper for File type
type _github_com_gotd_td_telegram_uploader_File struct {
	WRead func(p []byte) (n int, err error)
	WStat func() (fs.FileInfo, error)
}

func (W _github_com_gotd_td_telegram_uploader_File) Read(p []byte) (n int, err error) {
	return W.WRead(p)
}
func (W _github_com_gotd_td_telegram_uploader_File) Stat() (fs.FileInfo, error) { return W.WStat() }

// _github_com_gotd_td_telegram_uploader_Progress is an interface wrapper for Progress type
type _github_com_gotd_td_telegram_uploader_Progress struct {
	WChunk func(ctx context.Context, state uploader.ProgressState) error
}

func (W _github_com_gotd_td_telegram_uploader_Progress) Chunk(ctx context.Context, state uploader.ProgressState) error {
	return W.WChunk(ctx, state)
}
