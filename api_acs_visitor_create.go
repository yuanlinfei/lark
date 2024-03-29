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

// CreateACSVisitor 添加访客
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/visitor/create
func (r *ACSService) CreateACSVisitor(ctx context.Context, request *CreateACSVisitorReq, options ...MethodOptionFunc) (*CreateACSVisitorResp, *Response, error) {
	if r.cli.mock.mockACSCreateACSVisitor != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] ACS#CreateACSVisitor mock enable")
		return r.cli.mock.mockACSCreateACSVisitor(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "ACS",
		API:                 "CreateACSVisitor",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/acs/v1/visitors",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(createACSVisitorResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockACSCreateACSVisitor mock ACSCreateACSVisitor method
func (r *Mock) MockACSCreateACSVisitor(f func(ctx context.Context, request *CreateACSVisitorReq, options ...MethodOptionFunc) (*CreateACSVisitorResp, *Response, error)) {
	r.mockACSCreateACSVisitor = f
}

// UnMockACSCreateACSVisitor un-mock ACSCreateACSVisitor method
func (r *Mock) UnMockACSCreateACSVisitor() {
	r.mockACSCreateACSVisitor = nil
}

// CreateACSVisitorReq ...
type CreateACSVisitorReq struct {
	UserIDType *IDType                  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	User       *CreateACSVisitorReqUser `json:"user,omitempty"`         // 访客信息
}

// CreateACSVisitorReqUser ...
type CreateACSVisitorReqUser struct {
	UserType     int64   `json:"user_type,omitempty"`     // 用户类型, 示例值: 11, 可选值有: 1: 员工, 2: 部门, 10: 全体员工, 11: 访客
	UserID       *string `json:"user_id,omitempty"`       // 用户id, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	UserName     *string `json:"user_name,omitempty"`     // 用户名称, 示例值: "张三"
	PhoneNum     *string `json:"phone_num,omitempty"`     // 电话号码, 示例值: "1357890001"
	DepartmentID *string `json:"department_id,omitempty"` // 部门id, 示例值: "od-f7d44ab733f7602f5cc5194735fd9aaf"
}

// CreateACSVisitorResp ...
type CreateACSVisitorResp struct {
	VisitorID string `json:"visitor_id,omitempty"` // 访客的id
}

// createACSVisitorResp ...
type createACSVisitorResp struct {
	Code  int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                `json:"msg,omitempty"`  // 错误描述
	Data  *CreateACSVisitorResp `json:"data,omitempty"`
	Error *ErrorDetail          `json:"error,omitempty"`
}
