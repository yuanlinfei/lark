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

// DeleteDriveFileVersion 删除文档版本。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-version/delete
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/file-version/delete
func (r *DriveService) DeleteDriveFileVersion(ctx context.Context, request *DeleteDriveFileVersionReq, options ...MethodOptionFunc) (*DeleteDriveFileVersionResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteDriveFileVersion != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#DeleteDriveFileVersion mock enable")
		return r.cli.mock.mockDriveDeleteDriveFileVersion(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteDriveFileVersion",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/versions/:version_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteDriveFileVersionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDeleteDriveFileVersion mock DriveDeleteDriveFileVersion method
func (r *Mock) MockDriveDeleteDriveFileVersion(f func(ctx context.Context, request *DeleteDriveFileVersionReq, options ...MethodOptionFunc) (*DeleteDriveFileVersionResp, *Response, error)) {
	r.mockDriveDeleteDriveFileVersion = f
}

// UnMockDriveDeleteDriveFileVersion un-mock DriveDeleteDriveFileVersion method
func (r *Mock) UnMockDriveDeleteDriveFileVersion() {
	r.mockDriveDeleteDriveFileVersion = nil
}

// DeleteDriveFileVersionReq ...
type DeleteDriveFileVersionReq struct {
	FileToken  string  `path:"file_token" json:"-"`    // 版本文档token, 如何获取文档Token可以参考[如何获取云文档相关token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "doxbcyvqZlSc9WlHvQMlSJwUrsb"
	VersionID  string  `path:"version_id" json:"-"`    // 版本文档版本号, 示例值: "file_version"
	ObjType    string  `query:"obj_type" json:"-"`     // 文档类型, 示例值: docx, 可选值有: docx: 新版文档
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// DeleteDriveFileVersionResp ...
type DeleteDriveFileVersionResp struct {
}

// deleteDriveFileVersionResp ...
type deleteDriveFileVersionResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *DeleteDriveFileVersionResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}
