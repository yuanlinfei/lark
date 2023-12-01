// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// UpdateDrivePermissionPublicPassword 该接口用于根据 filetoken 刷新云文档的密码。
//
// 注意: 刷新密码, 需要先通过”云文档“-“权限”-“设置”-“更新云文档权限设置”的接口更新元文档为互联网上获得链接的任何人可阅读/编辑
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-public-password/update
// new doc: https://open.feishu.cn/document/server-docs/docs/permission/permission-public/permission-public-password/update
func (r *DriveService) UpdateDrivePermissionPublicPassword(ctx context.Context, request *UpdateDrivePermissionPublicPasswordReq, options ...MethodOptionFunc) (*UpdateDrivePermissionPublicPasswordResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateDrivePermissionPublicPassword != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateDrivePermissionPublicPassword mock enable")
		return r.cli.mock.mockDriveUpdateDrivePermissionPublicPassword(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateDrivePermissionPublicPassword",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/permissions/:token/public/password",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateDrivePermissionPublicPasswordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateDrivePermissionPublicPassword mock DriveUpdateDrivePermissionPublicPassword method
func (r *Mock) MockDriveUpdateDrivePermissionPublicPassword(f func(ctx context.Context, request *UpdateDrivePermissionPublicPasswordReq, options ...MethodOptionFunc) (*UpdateDrivePermissionPublicPasswordResp, *Response, error)) {
	r.mockDriveUpdateDrivePermissionPublicPassword = f
}

// UnMockDriveUpdateDrivePermissionPublicPassword un-mock DriveUpdateDrivePermissionPublicPassword method
func (r *Mock) UnMockDriveUpdateDrivePermissionPublicPassword() {
	r.mockDriveUpdateDrivePermissionPublicPassword = nil
}

// UpdateDrivePermissionPublicPasswordReq ...
type UpdateDrivePermissionPublicPasswordReq struct {
	Token string `path:"token" json:"-"` // 文件的 token, 获取方式见 [如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "doccnBKgoMyY5OMbUG6FioTXuBe"
	Type  string `query:"type" json:"-"` // 文件类型, 需要与文件的 token 相匹配, 示例值: doc, 可选值有: doc: 文档, sheet: 电子表格, file: 云空间文件, wiki: 知识库节点, bitable: 多维表格, docx: 新版文档, mindnote: 思维笔记, minutes: 妙计（暂不支持）
}

// UpdateDrivePermissionPublicPasswordResp ...
type UpdateDrivePermissionPublicPasswordResp struct {
	Password string `json:"password,omitempty"` // 密码
}

// updateDrivePermissionPublicPasswordResp ...
type updateDrivePermissionPublicPasswordResp struct {
	Code int64                                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                   `json:"msg,omitempty"`  // 错误描述
	Data *UpdateDrivePermissionPublicPasswordResp `json:"data,omitempty"`
}
