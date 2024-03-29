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

// UpdateApplicationFeedback 更新应用的反馈数据
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-feedback/patch
// new doc: https://open.feishu.cn/document/server-docs/application-v6/application-feedback/patch
func (r *ApplicationService) UpdateApplicationFeedback(ctx context.Context, request *UpdateApplicationFeedbackReq, options ...MethodOptionFunc) (*UpdateApplicationFeedbackResp, *Response, error) {
	if r.cli.mock.mockApplicationUpdateApplicationFeedback != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Application#UpdateApplicationFeedback mock enable")
		return r.cli.mock.mockApplicationUpdateApplicationFeedback(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "UpdateApplicationFeedback",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v6/applications/:app_id/feedbacks/:feedback_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateApplicationFeedbackResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationUpdateApplicationFeedback mock ApplicationUpdateApplicationFeedback method
func (r *Mock) MockApplicationUpdateApplicationFeedback(f func(ctx context.Context, request *UpdateApplicationFeedbackReq, options ...MethodOptionFunc) (*UpdateApplicationFeedbackResp, *Response, error)) {
	r.mockApplicationUpdateApplicationFeedback = f
}

// UnMockApplicationUpdateApplicationFeedback un-mock ApplicationUpdateApplicationFeedback method
func (r *Mock) UnMockApplicationUpdateApplicationFeedback() {
	r.mockApplicationUpdateApplicationFeedback = nil
}

// UpdateApplicationFeedbackReq ...
type UpdateApplicationFeedbackReq struct {
	AppID      string  `path:"app_id" json:"-"`        // 目标应用 ID（本租户创建的所有应用）, 示例值: "cli_9f115af860f7901b"
	FeedbackID string  `path:"feedback_id" json:"-"`   // 应用反馈记录id, 示例值: "7057888018203574291"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Status     int64   `query:"status" json:"-"`       // 反馈处理状态, 示例值: 1, 可选值有: 0: 反馈未处理, 1: 反馈已处理, 2: 反馈处理中, 3: 反馈已关闭
	OperatorID string  `query:"operator_id" json:"-"`  // 反馈处理人员id, 租户内用户的唯一标识, ID值与查询参数中的user_id_type 对应, 示例值: ou_9565b69967831233761cc2f11b4c089f
}

// UpdateApplicationFeedbackResp ...
type UpdateApplicationFeedbackResp struct {
}

// updateApplicationFeedbackResp ...
type updateApplicationFeedbackResp struct {
	Code  int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                         `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateApplicationFeedbackResp `json:"data,omitempty"`
	Error *ErrorDetail                   `json:"error,omitempty"`
}
