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

// CreatePublicMailbox 创建一个公共邮箱。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/create
// new doc: https://open.feishu.cn/document/server-docs/mail-v1/public-mailbox/public_mailbox/create
func (r *MailService) CreatePublicMailbox(ctx context.Context, request *CreatePublicMailboxReq, options ...MethodOptionFunc) (*CreatePublicMailboxResp, *Response, error) {
	if r.cli.mock.mockMailCreatePublicMailbox != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Mail#CreatePublicMailbox mock enable")
		return r.cli.mock.mockMailCreatePublicMailbox(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "CreatePublicMailbox",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/public_mailboxes",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createPublicMailboxResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailCreatePublicMailbox mock MailCreatePublicMailbox method
func (r *Mock) MockMailCreatePublicMailbox(f func(ctx context.Context, request *CreatePublicMailboxReq, options ...MethodOptionFunc) (*CreatePublicMailboxResp, *Response, error)) {
	r.mockMailCreatePublicMailbox = f
}

// UnMockMailCreatePublicMailbox un-mock MailCreatePublicMailbox method
func (r *Mock) UnMockMailCreatePublicMailbox() {
	r.mockMailCreatePublicMailbox = nil
}

// CreatePublicMailboxReq ...
type CreatePublicMailboxReq struct {
	Email *string `json:"email,omitempty"` // 公共邮箱地址, 示例值: "test_public_mailbox@xxx.xx"
	Name  *string `json:"name,omitempty"`  // 公共邮箱名称, 示例值: "test public mailbox"
	Geo   *string `json:"geo,omitempty"`   // 数据驻留地, 示例值: "cn"
}

// CreatePublicMailboxResp ...
type CreatePublicMailboxResp struct {
	PublicMailboxID string `json:"public_mailbox_id,omitempty"` // 公共邮箱唯一标识
	Email           string `json:"email,omitempty"`             // 公共邮箱地址
	Name            string `json:"name,omitempty"`              // 公共邮箱名称
	Geo             string `json:"geo,omitempty"`               // 数据驻留地, 字段权限要求: 查看公共邮箱数据驻留地
}

// createPublicMailboxResp ...
type createPublicMailboxResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *CreatePublicMailboxResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
