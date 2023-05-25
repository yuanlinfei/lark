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

// GetChatMemberList 获取用户/机器人所在群的群成员列表。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 机器人或授权用户必须在群组中
// - 该接口不会返回群内的机器人成员
// - 由于返回的群成员列表会过滤掉机器人成员, 因此返回的群成员个数可能会小于指定的page_size
// - 如果有同一时间加入群的群成员, 会一次性返回, 这会导致返回的群成员个数可能会大于指定的page_size
// - 获取内部群信息时, 操作者须与群组在同一租户下
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/get
func (r *ChatService) GetChatMemberList(ctx context.Context, request *GetChatMemberListReq, options ...MethodOptionFunc) (*GetChatMemberListResp, *Response, error) {
	if r.cli.mock.mockChatGetChatMemberList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#GetChatMemberList mock enable")
		return r.cli.mock.mockChatGetChatMemberList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "GetChatMemberList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getChatMemberListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatGetChatMemberList mock ChatGetChatMemberList method
func (r *Mock) MockChatGetChatMemberList(f func(ctx context.Context, request *GetChatMemberListReq, options ...MethodOptionFunc) (*GetChatMemberListResp, *Response, error)) {
	r.mockChatGetChatMemberList = f
}

// UnMockChatGetChatMemberList un-mock ChatGetChatMemberList method
func (r *Mock) UnMockChatGetChatMemberList() {
	r.mockChatGetChatMemberList = nil
}

// GetChatMemberListReq ...
type GetChatMemberListReq struct {
	ChatID       string  `path:"chat_id" json:"-"`         // 群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: "oc_a0553eda9014c201e6969b478895c230"
	MemberIDType *IDType `query:"member_id_type" json:"-"` // 群成员 用户 ID 类型, 详情参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 示例值: "open_id", 可选值有: user_id: 以 user_id 来识别成员, union_id: 以 union_id 来识别成员, open_id: 以 open_id 来识别成员
	PageSize     *int64  `query:"page_size" json:"-"`      // 分页大小, 示例值: 20, 默认值: `20`, 最大值: `100`
	PageToken    *string `query:"page_token" json:"-"`     // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "WWxHTStrOEs5WHZpNktGbU94bUcvMWlxdDUzTWt1OXNrRmlLaGRNVG0yaz0="
}

// GetChatMemberListResp ...
type GetChatMemberListResp struct {
	Items       []*GetChatMemberListRespItem `json:"items,omitempty"`        // 成员列表
	PageToken   string                       `json:"page_token,omitempty"`   // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore     bool                         `json:"has_more,omitempty"`     // 是否还有更多项
	MemberTotal int64                        `json:"member_total,omitempty"` // 成员总数
}

// GetChatMemberListRespItem ...
type GetChatMemberListRespItem struct {
	MemberIDType IDType `json:"member_id_type,omitempty"` // 成员的用户 ID 类型, 与查询参数中的 member_id_type 相同。取值为: `open_id`、`user_id`、`union_id`其中之一。
	MemberID     string `json:"member_id,omitempty"`      // 成员的用户ID, ID值与查询参数中的 member_id_type 对应, 不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Name         string `json:"name,omitempty"`           // 名字
	TenantKey    string `json:"tenant_key,omitempty"`     // 租户Key, 为租户在飞书上的唯一标识, 用来换取对应的tenant_access_token, 也可以用作租户在应用中的唯一标识
}

// getChatMemberListResp ...
type getChatMemberListResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetChatMemberListResp `json:"data,omitempty"`
}
