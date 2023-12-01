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

// BatchGetUserByID 通过该接口, 可使用手机号/邮箱获取用户的 ID 信息, 具体获取支持的 ID 类型包括 open_id、user_id、union_id, 可通过查询参数指定。
//
// 不返回用户 ID 的情况说明:
// - 查询的手机号或者邮箱不存在。
// - 未开通 获取用户 user ID 权限, 将无法返回用户的 user_id。
// - 无权限查看用户信息。你可以在开发者后台应用详情页的 权限管理 > 数据权限 > 通讯录权限范围 中, 指定用户权限范围。
// - 使用企业邮箱查询将无法返回用户 ID, 需要使用用户的邮箱地址。
// - 在请求头 Authorization 中, 传入的 Token 有误。例如, Token 对应的应用与实际所需应用不一致。
// - 所查询的用户已离职。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/batch_get_id
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/user/batch_get_id
func (r *ContactService) BatchGetUserByID(ctx context.Context, request *BatchGetUserByIDReq, options ...MethodOptionFunc) (*BatchGetUserByIDResp, *Response, error) {
	if r.cli.mock.mockContactBatchGetUserByID != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#BatchGetUserByID mock enable")
		return r.cli.mock.mockContactBatchGetUserByID(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "BatchGetUserByID",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/users/batch_get_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetUserByIDResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactBatchGetUserByID mock ContactBatchGetUserByID method
func (r *Mock) MockContactBatchGetUserByID(f func(ctx context.Context, request *BatchGetUserByIDReq, options ...MethodOptionFunc) (*BatchGetUserByIDResp, *Response, error)) {
	r.mockContactBatchGetUserByID = f
}

// UnMockContactBatchGetUserByID un-mock ContactBatchGetUserByID method
func (r *Mock) UnMockContactBatchGetUserByID() {
	r.mockContactBatchGetUserByID = nil
}

// BatchGetUserByIDReq ...
type BatchGetUserByIDReq struct {
	UserIDType      *IDType  `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值: user_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Emails          []string `json:"emails,omitempty"`           // 要查询的用户邮箱（不支持企业邮箱）, 最多 50 条, 注意, emails与mobiles相互独立, 每条用户邮箱返回对应的用户ID, 本接口返回的用户ID数量为emails数量与mobiles数量的和, 示例值: ["abc@z.com"], 最大长度: `50`
	Mobiles         []string `json:"mobiles,omitempty"`          // 要查询的用户手机号, 最多 50 条, 注意, 1. emails与mobiles相互独立, 每条用户手机号返回对应的用户ID, 2.  非中国大陆地区的手机号需要添加以 “+” 开头的国家 / 地区代码, 示例值: ["13812345678"], 最大长度: `50`
	IncludeResigned *bool    `json:"include_resigned,omitempty"` // 查询结果是否包含离职员工。取值为 true 后可查询离职用户的 ID, 示例值: true, 默认值: `false`
}

// BatchGetUserByIDResp ...
type BatchGetUserByIDResp struct {
	UserList []*BatchGetUserByIDRespUser `json:"user_list,omitempty"` // 手机号或者邮箱对应的用户id信息
}

// BatchGetUserByIDRespUser ...
type BatchGetUserByIDRespUser struct {
	UserID string `json:"user_id,omitempty"` // 用户id, 值为user_id_type所指定的类型。如果查询的手机号、邮箱不存在, 或者无权限查看对应的用户, 则不返回此项。
	Mobile string `json:"mobile,omitempty"`  // 手机号, 通过手机号查询时返回
	Email  string `json:"email,omitempty"`   // 邮箱, 通过邮箱查询时返回
}

// batchGetUserByIDResp ...
type batchGetUserByIDResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *BatchGetUserByIDResp `json:"data,omitempty"`
}
