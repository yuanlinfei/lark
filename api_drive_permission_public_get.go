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

// GetDrivePublicPermission 该接口用于根据 filetoken 获取云文档的权限设置。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-public/get
// new doc: https://open.feishu.cn/document/server-docs/docs/permission/permission-public/get
func (r *DriveService) GetDrivePublicPermission(ctx context.Context, request *GetDrivePublicPermissionReq, options ...MethodOptionFunc) (*GetDrivePublicPermissionResp, *Response, error) {
	if r.cli.mock.mockDriveGetDrivePublicPermission != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDrivePublicPermission mock enable")
		return r.cli.mock.mockDriveGetDrivePublicPermission(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDrivePublicPermission",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/permissions/:token/public",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDrivePublicPermissionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDrivePublicPermission mock DriveGetDrivePublicPermission method
func (r *Mock) MockDriveGetDrivePublicPermission(f func(ctx context.Context, request *GetDrivePublicPermissionReq, options ...MethodOptionFunc) (*GetDrivePublicPermissionResp, *Response, error)) {
	r.mockDriveGetDrivePublicPermission = f
}

// UnMockDriveGetDrivePublicPermission un-mock DriveGetDrivePublicPermission method
func (r *Mock) UnMockDriveGetDrivePublicPermission() {
	r.mockDriveGetDrivePublicPermission = nil
}

// GetDrivePublicPermissionReq ...
type GetDrivePublicPermissionReq struct {
	Token string `path:"token" json:"-"` // 文件的 token, 获取方式见 [如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "doccnBKgoMyY5OMbUG6FioTXuBe"
	Type  string `query:"type" json:"-"` // 文件类型, 需要与文件的 token 相匹配, 示例值: doc, 可选值有: doc: 文档, sheet: 电子表格, file: 云空间文件, wiki: 知识库节点, bitable: 多维表格, docx: 新版文档, mindnote: 思维笔记, minutes: 妙记
}

// GetDrivePublicPermissionResp ...
type GetDrivePublicPermissionResp struct {
	PermissionPublic *GetDrivePublicPermissionRespPermissionPublic `json:"permission_public,omitempty"` // 返回的文档权限设置
}

// GetDrivePublicPermissionRespPermissionPublic ...
type GetDrivePublicPermissionRespPermissionPublic struct {
	ExternalAccess  bool   `json:"external_access,omitempty"`   // 允许内容被分享到组织外, 可选值有: `true`: 允许, `false`: 不允许
	SecurityEntity  string `json:"security_entity,omitempty"`   // 谁可以复制内容、创建副本、打印、下载, 可选值有: anyone_can_view: 拥有可阅读权限的用户, anyone_can_edit: 拥有可编辑权限的用户, only_full_access: 拥有可管理权限（包括我）的用户
	CommentEntity   string `json:"comment_entity,omitempty"`    // 可评论设置, 可选值有: anyone_can_view: 拥有可阅读权限的用户, anyone_can_edit: 拥有可编辑权限的用户
	ShareEntity     string `json:"share_entity,omitempty"`      // 谁可以添加和管理协作者, 可选值有: anyone: 所有可阅读或编辑此文档的用户, same_tenant: 组织内所有可阅读或编辑此文档的用户, only_full_access: 拥有可管理权限（包括我）的用户
	LinkShareEntity string `json:"link_share_entity,omitempty"` // 链接分享设置, 可选值有: tenant_readable: 组织内获得链接的人可阅读, tenant_editable: 组织内获得链接的人可编辑, anyone_readable: 互联网上获得链接的任何人可阅读, anyone_editable: 互联网上获得链接的任何人可编辑, closed: 关闭链接分享
	InviteExternal  bool   `json:"invite_external,omitempty"`   // 允许非「可管理权限」的人分享到组织外
	LockSwitch      bool   `json:"lock_switch,omitempty"`       // 节点加锁状态
}

// getDrivePublicPermissionResp ...
type getDrivePublicPermissionResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *GetDrivePublicPermissionResp `json:"data,omitempty"`
}
