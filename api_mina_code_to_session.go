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

// MinaCodeToSession 通过 [login](https://open.feishu.cn/document/uYjL24iN/uYzMuYzMuYzM)接口获取到登录凭证`code`后, 开发者可以通过服务器发送请求的方式获取 session_key 和 用户凭证信息。
//
// 本接口适用于 [小程序登录](https://open.feishu.cn/document/uYjL24iN/uETO5QjLxkTO04SM5kDN) 及[小组件登录](https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/block/guide/open-ability/block-login)。
//
// doc: https://open.feishu.cn/document/uYjL24iN/ukjM04SOyQjL5IDN
// new doc: https://open.feishu.cn/document/client-docs/gadget/-web-app-api/open-ability/login/code2session
//
// Deprecated
func (r *MinaService) MinaCodeToSession(ctx context.Context, request *MinaCodeToSessionReq, options ...MethodOptionFunc) (*MinaCodeToSessionResp, *Response, error) {
	if r.cli.mock.mockMinaMinaCodeToSession != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Mina#MinaCodeToSession mock enable")
		return r.cli.mock.mockMinaMinaCodeToSession(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:              "Mina",
		API:                "MinaCodeToSession",
		Method:             "POST",
		URL:                r.cli.openBaseURL + "/open-apis/mina/v2/tokenLoginValidate",
		Body:               request,
		MethodOption:       newMethodOption(options),
		NeedAppAccessToken: true,
	}
	resp := new(minaCodeToSessionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMinaMinaCodeToSession mock MinaMinaCodeToSession method
func (r *Mock) MockMinaMinaCodeToSession(f func(ctx context.Context, request *MinaCodeToSessionReq, options ...MethodOptionFunc) (*MinaCodeToSessionResp, *Response, error)) {
	r.mockMinaMinaCodeToSession = f
}

// UnMockMinaMinaCodeToSession un-mock MinaMinaCodeToSession method
func (r *Mock) UnMockMinaMinaCodeToSession() {
	r.mockMinaMinaCodeToSession = nil
}

// MinaCodeToSessionReq ...
type MinaCodeToSessionReq struct {
	Code string `json:"code,omitempty"` // [登录](https://open.feishu.cn/document/uYjL24iN/uYzMuYzMuYzM)时获取的 code
}

// MinaCodeToSessionResp ...
type MinaCodeToSessionResp struct {
	OpenID       string `json:"open_id,omitempty"`       // 用户的[Open ID](https://open.feishu.cn/document/home/user-identity-introduction/open-id), 用于在同一个应用中对用户进行标识
	EmployeeID   string `json:"employee_id,omitempty"`   // 用户的[User ID](https://open.feishu.cn/document/home/user-identity-introduction/user-id), 在职员工在企业内的唯一标识, 仅当开通以下权限后 返回该字段: 获取用户 user ID
	SessionKey   string `json:"session_key,omitempty"`   // 会话密钥
	TenantKey    string `json:"tenant_key,omitempty"`    // 用户所在企业唯一标识
	AccessToken  string `json:"access_token,omitempty"`  // [user_access_token](https://open.feishu.cn/document/ukTMukTMukTM/uMTNz4yM1MjLzUzM), 用户身份访问凭证
	ExpiresIn    int64  `json:"expires_in,omitempty"`    // [user_access_token](https://open.feishu.cn/document/ukTMukTMukTM/uMTNz4yM1MjLzUzM)过期时间戳
	RefreshToken string `json:"refresh_token,omitempty"` // 刷新用户 access_token 时使用的 token, 过期时间为30天。刷新access_token 接口说明请查看[文档](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/authen/refresh_access_token)
}

// minaCodeToSessionResp ...
type minaCodeToSessionResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *MinaCodeToSessionResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
