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

// CreateApprovalComment 在某审批实例下创建、修改评论或评论回复（不包含审批同意、拒绝、转交等附加的理由或意见）。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance-comment/create
// new doc: https://open.feishu.cn/document/server-docs/approval-v4/instance-comment/create
func (r *ApprovalService) CreateApprovalComment(ctx context.Context, request *CreateApprovalCommentReq, options ...MethodOptionFunc) (*CreateApprovalCommentResp, *Response, error) {
	if r.cli.mock.mockApprovalCreateApprovalComment != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Approval#CreateApprovalComment mock enable")
		return r.cli.mock.mockApprovalCreateApprovalComment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "CreateApprovalComment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/instances/:instance_id/comments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createApprovalCommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalCreateApprovalComment mock ApprovalCreateApprovalComment method
func (r *Mock) MockApprovalCreateApprovalComment(f func(ctx context.Context, request *CreateApprovalCommentReq, options ...MethodOptionFunc) (*CreateApprovalCommentResp, *Response, error)) {
	r.mockApprovalCreateApprovalComment = f
}

// UnMockApprovalCreateApprovalComment un-mock ApprovalCreateApprovalComment method
func (r *Mock) UnMockApprovalCreateApprovalComment() {
	r.mockApprovalCreateApprovalComment = nil
}

// CreateApprovalCommentReq ...
type CreateApprovalCommentReq struct {
	InstanceID      string                            `path:"instance_id" json:"-"`        // 审批实例code（或租户自定义审批实例ID）, 示例值: "6A123516-FB88-470D-A428-9AF58B71B3C0"
	UserIDType      *IDType                           `query:"user_id_type" json:"-"`      // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	UserID          string                            `query:"user_id" json:"-"`           // 用户ID, 示例值: "e5286g26"
	Content         *string                           `json:"content,omitempty"`           // 评论内容, 包含艾特人、附件等, 示例值: "{\"text\":\"来自小程序的评论111我带附件中有extra \", \"files\":[{\"url\":\"xxx\", \"fileSize\":155149, \"title\":\"9a9fedc5cfb01a4a20c715098.png\", \"type\":\"image\", \"extra\":\"\"}]}"
	AtInfoList      []*CreateApprovalCommentReqAtInfo `json:"at_info_list,omitempty"`      // 评论中艾特人信息
	ParentCommentID *string                           `json:"parent_comment_id,omitempty"` // 父评论ID, 如果是回复评论, 需要传, 示例值: "7081516627711524883"
	CommentID       *string                           `json:"comment_id,omitempty"`        // 评论ID, 如果是编辑、删除一条评论, 需要传, 示例值: "7081516627711524883"
	DisableBot      *bool                             `json:"disable_bot,omitempty"`       // disable_bot=true只同步数据, 不触发bot, 示例值: false
	Extra           *string                           `json:"extra,omitempty"`             // 附加字段, 示例值: "{\"a\":\"a\"}"
}

// CreateApprovalCommentReqAtInfo ...
type CreateApprovalCommentReqAtInfo struct {
	UserID string `json:"user_id,omitempty"` // 被艾特人的ID, 示例值: "579fd9c4"
	Name   string `json:"name,omitempty"`    // 被艾特人的姓名, 示例值: "张某"
	Offset string `json:"offset,omitempty"`  // 被艾特人在评论中的位置, 从0开始, 示例值: "1"
}

// CreateApprovalCommentResp ...
type CreateApprovalCommentResp struct {
	CommentID string `json:"comment_id,omitempty"` // 保存成功的comment_id
}

// createApprovalCommentResp ...
type createApprovalCommentResp struct {
	Code  int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 错误描述
	Data  *CreateApprovalCommentResp `json:"data,omitempty"`
	Error *ErrorDetail               `json:"error,omitempty"`
}
