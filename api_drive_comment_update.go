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

// UpdateDriveComment 更新云文档中的某条回复的内容。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment-reply/update
// new doc: https://open.feishu.cn/document/server-docs/docs/CommentAPI/update
func (r *DriveService) UpdateDriveComment(ctx context.Context, request *UpdateDriveCommentReq, options ...MethodOptionFunc) (*UpdateDriveCommentResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateDriveComment != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#UpdateDriveComment mock enable")
		return r.cli.mock.mockDriveUpdateDriveComment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateDriveComment",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/comments/:comment_id/replies/:reply_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateDriveCommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateDriveComment mock DriveUpdateDriveComment method
func (r *Mock) MockDriveUpdateDriveComment(f func(ctx context.Context, request *UpdateDriveCommentReq, options ...MethodOptionFunc) (*UpdateDriveCommentResp, *Response, error)) {
	r.mockDriveUpdateDriveComment = f
}

// UnMockDriveUpdateDriveComment un-mock DriveUpdateDriveComment method
func (r *Mock) UnMockDriveUpdateDriveComment() {
	r.mockDriveUpdateDriveComment = nil
}

// UpdateDriveCommentReq ...
type UpdateDriveCommentReq struct {
	FileToken  string                        `path:"file_token" json:"-"`    // 文档 Token, 示例值: "doxbcdl03Vsxhm7Qmnj110abcef"
	CommentID  string                        `path:"comment_id" json:"-"`    // 评论 ID, 示例值: "6916106822734578184"
	ReplyID    string                        `path:"reply_id" json:"-"`      // 回复 ID, 示例值: "6916106822734594568"
	FileType   FileType                      `query:"file_type" json:"-"`    // 文档类型, 示例值: doc, 可选值有: doc: 文档, sheet: 表格, file: 文件, docx: 新版文档
	UserIDType *IDType                       `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Content    *UpdateDriveCommentReqContent `json:"content,omitempty"`      // 回复内容
}

// UpdateDriveCommentReqContent ...
type UpdateDriveCommentReqContent struct {
	Elements []*UpdateDriveCommentReqContentElement `json:"elements,omitempty"` // 回复的内容
}

// UpdateDriveCommentReqContentElement ...
type UpdateDriveCommentReqContentElement struct {
	Type     string                                       `json:"type,omitempty"`      // 回复的内容元素, 示例值: "text_run", 可选值有: text_run: 普通文本, docs_link: at 云文档链接, person: at 联系人
	TextRun  *UpdateDriveCommentReqContentElementTextRun  `json:"text_run,omitempty"`  // 文本内容
	DocsLink *UpdateDriveCommentReqContentElementDocsLink `json:"docs_link,omitempty"` // 添加云文档链接
	Person   *UpdateDriveCommentReqContentElementPerson   `json:"person,omitempty"`    // 添加用户的 user_id
}

// UpdateDriveCommentReqContentElementDocsLink ...
type UpdateDriveCommentReqContentElementDocsLink struct {
	URL string `json:"url,omitempty"` // 回复 at 云文档, 示例值: "https://example.feishu.cn/docs/doccnHh7U87HOFpii5u5Gabcef"
}

// UpdateDriveCommentReqContentElementPerson ...
type UpdateDriveCommentReqContentElementPerson struct {
	UserID string `json:"user_id,omitempty"` // 添加用户的 user_id 以@用户, 示例值: "ou_cc19b2bfb93f8a44db4b4d6eababcef"
}

// UpdateDriveCommentReqContentElementTextRun ...
type UpdateDriveCommentReqContentElementTextRun struct {
	Text string `json:"text,omitempty"` // 回复 普通文本, 示例值: "comment text"
}

// UpdateDriveCommentResp ...
type UpdateDriveCommentResp struct {
}

// updateDriveCommentResp ...
type updateDriveCommentResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateDriveCommentResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}
