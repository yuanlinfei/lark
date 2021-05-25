// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetDriveRootFolderMeta
//
// **该篇文档有两个接口，一个接口用于获取指定文件夹的元信息，一个用于获取我的空间文件夹的元信息**
// 该接口用于根据 folderToken 获取该文件夹的元信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uAjNzUjLwYzM14CM2MTN
func (r *DriveService) GetDriveRootFolderMeta(ctx context.Context, request *GetDriveRootFolderMetaReq, options ...MethodOptionFunc) (*GetDriveRootFolderMetaResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveRootFolderMeta != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveRootFolderMeta mock enable")
		return r.cli.mock.mockDriveGetDriveRootFolderMeta(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Drive",
		API:                 "GetDriveRootFolderMeta",
		Method:              "GET",
		URL:                 "https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken/meta   ",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(getDriveRootFolderMetaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveGetDriveRootFolderMeta(f func(ctx context.Context, request *GetDriveRootFolderMetaReq, options ...MethodOptionFunc) (*GetDriveRootFolderMetaResp, *Response, error)) {
	r.mockDriveGetDriveRootFolderMeta = f
}

func (r *Mock) UnMockDriveGetDriveRootFolderMeta() {
	r.mockDriveGetDriveRootFolderMeta = nil
}

type GetDriveRootFolderMetaReq struct{}

type getDriveRootFolderMetaResp struct {
	Code int64                       `json:"code,omitempty"`
	Msg  string                      `json:"msg,omitempty"`
	Data *GetDriveRootFolderMetaResp `json:"data,omitempty"`
}

type GetDriveRootFolderMetaResp struct {
	Token  string `json:"token,omitempty"`   // 文件夹的 token
	ID     string `json:"id,omitempty"`      // 文件夹的 id
	UserID string `json:"user_id,omitempty"` // 文件夹的所有者 id
}
