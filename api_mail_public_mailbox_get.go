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

// GetPublicMailbox 获取公共邮箱信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/get
// new doc: https://open.feishu.cn/document/server-docs/mail-v1/public-mailbox/public_mailbox/get
func (r *MailService) GetPublicMailbox(ctx context.Context, request *GetPublicMailboxReq, options ...MethodOptionFunc) (*GetPublicMailboxResp, *Response, error) {
	if r.cli.mock.mockMailGetPublicMailbox != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Mail#GetPublicMailbox mock enable")
		return r.cli.mock.mockMailGetPublicMailbox(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "GetPublicMailbox",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getPublicMailboxResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailGetPublicMailbox mock MailGetPublicMailbox method
func (r *Mock) MockMailGetPublicMailbox(f func(ctx context.Context, request *GetPublicMailboxReq, options ...MethodOptionFunc) (*GetPublicMailboxResp, *Response, error)) {
	r.mockMailGetPublicMailbox = f
}

// UnMockMailGetPublicMailbox un-mock MailGetPublicMailbox method
func (r *Mock) UnMockMailGetPublicMailbox() {
	r.mockMailGetPublicMailbox = nil
}

// GetPublicMailboxReq ...
type GetPublicMailboxReq struct {
	PublicMailboxID string `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值: "xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
}

// GetPublicMailboxResp ...
type GetPublicMailboxResp struct {
	PublicMailboxID string `json:"public_mailbox_id,omitempty"` // 公共邮箱唯一标识
	Email           string `json:"email,omitempty"`             // 公共邮箱地址
	Name            string `json:"name,omitempty"`              // 公共邮箱名称
	Geo             string `json:"geo,omitempty"`               // 数据驻留地, 字段权限要求: 查看公共邮箱数据驻留地
}

// getPublicMailboxResp ...
type getPublicMailboxResp struct {
	Code  int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                `json:"msg,omitempty"`  // 错误描述
	Data  *GetPublicMailboxResp `json:"data,omitempty"`
	Error *ErrorDetail          `json:"error,omitempty"`
}
