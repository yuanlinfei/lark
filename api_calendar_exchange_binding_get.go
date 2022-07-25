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

// GetCalendarExchangeBinding 本接口获取Exchange账户的绑定状态, 包括exchange日历是否同步完成。
//
// 操作用户需要是企业超级管理员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/get
func (r *CalendarService) GetCalendarExchangeBinding(ctx context.Context, request *GetCalendarExchangeBindingReq, options ...MethodOptionFunc) (*GetCalendarExchangeBindingResp, *Response, error) {
	if r.cli.mock.mockCalendarGetCalendarExchangeBinding != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#GetCalendarExchangeBinding mock enable")
		return r.cli.mock.mockCalendarGetCalendarExchangeBinding(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Calendar",
		API:                 "GetCalendarExchangeBinding",
		Method:              "GET",
		URL:                 r.cli.openBaseURL + "/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(getCalendarExchangeBindingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarGetCalendarExchangeBinding mock CalendarGetCalendarExchangeBinding method
func (r *Mock) MockCalendarGetCalendarExchangeBinding(f func(ctx context.Context, request *GetCalendarExchangeBindingReq, options ...MethodOptionFunc) (*GetCalendarExchangeBindingResp, *Response, error)) {
	r.mockCalendarGetCalendarExchangeBinding = f
}

// UnMockCalendarGetCalendarExchangeBinding un-mock CalendarGetCalendarExchangeBinding method
func (r *Mock) UnMockCalendarGetCalendarExchangeBinding() {
	r.mockCalendarGetCalendarExchangeBinding = nil
}

// GetCalendarExchangeBindingReq ...
type GetCalendarExchangeBindingReq struct {
	ExchangeBindingID string  `path:"exchange_binding_id" json:"-"` // exchange绑定唯一标识id。参见[exchange绑定ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/introduction#12533d5e), 示例值: "ZW1haWxfYWRtaW5fZXhhbXBsZUBvdXRsb29rLmNvbSBlbWFpbF9hY2NvdW50X2V4YW1wbGVAb3V0bG9vay5jb20="
	UserIDType        *IDType `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetCalendarExchangeBindingResp ...
type GetCalendarExchangeBindingResp struct {
	AdminAccount      string `json:"admin_account,omitempty"`       // admin账户, 字段权限要求: 获取用户邮箱信息
	ExchangeAccount   string `json:"exchange_account,omitempty"`    // 用户绑定的exchange账户, 字段权限要求: 获取用户邮箱信息
	UserID            string `json:"user_id,omitempty"`             // exchange账户绑定user唯一标识id, 参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Status            string `json:"status,omitempty"`              // exchange账户同步状态, 可选值有: doing: 日历正在同步, cal_done: 日历同步完成, timespan_done: 近期时间段同步完成, done: 日程同步完成, err: 同步错误
	ExchangeBindingID string `json:"exchange_binding_id,omitempty"` // exchange绑定唯一标识id。参见[exchange绑定ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/introduction#12533d5e)
}

// getCalendarExchangeBindingResp ...
type getCalendarExchangeBindingResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarExchangeBindingResp `json:"data,omitempty"`
}
