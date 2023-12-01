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

// GetMessageSpecialFocusUnread 支持按单聊类型和群聊类型获取用户的特别关注未读消息数。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/special_focus/unread
//
// Deprecated
func (r *MessageService) GetMessageSpecialFocusUnread(ctx context.Context, request *GetMessageSpecialFocusUnreadReq, options ...MethodOptionFunc) (*GetMessageSpecialFocusUnreadResp, *Response, error) {
	if r.cli.mock.mockMessageGetMessageSpecialFocusUnread != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#GetMessageSpecialFocusUnread mock enable")
		return r.cli.mock.mockMessageGetMessageSpecialFocusUnread(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Message",
		API:                 "GetMessageSpecialFocusUnread",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/im/v1/special_focus/unread",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(getMessageSpecialFocusUnreadResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageGetMessageSpecialFocusUnread mock MessageGetMessageSpecialFocusUnread method
func (r *Mock) MockMessageGetMessageSpecialFocusUnread(f func(ctx context.Context, request *GetMessageSpecialFocusUnreadReq, options ...MethodOptionFunc) (*GetMessageSpecialFocusUnreadResp, *Response, error)) {
	r.mockMessageGetMessageSpecialFocusUnread = f
}

// UnMockMessageGetMessageSpecialFocusUnread un-mock MessageGetMessageSpecialFocusUnread method
func (r *Mock) UnMockMessageGetMessageSpecialFocusUnread() {
	r.mockMessageGetMessageSpecialFocusUnread = nil
}

// GetMessageSpecialFocusUnreadReq ...
type GetMessageSpecialFocusUnreadReq struct {
	MemberIDType *IDType   `query:"member_id_type" json:"-"` // 指定请求体中 [id_list] 的ID类型, 示例值: open_id, 可选值有: user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), app_id: 以app_id来识别成员。[了解更多: 如何获取应用的App ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-app-id), 默认值: `open_id`
	IDList       []string  `json:"id_list,omitempty"`        // 特别关注ID列表。请根据 [member_id_type] 参数填入`open_id`、`user_id`、`union_id`或`app_id`类型的ID, 示例值: ["ou_e167f0c694cb1c77bb040857dd963024"]
	ChatMode     *ChatMode `json:"chat_mode,omitempty"`      // 按群模式过滤特别关注, 注意: 当指定 [chat_mode] 为 [p2p]时, 将返回用户与被特别关注者之间单聊的未读消息数, 当指定[chat_mode] 为[group]时, 将返回被用户特别关注者在群聊中At用户且用户未读的消息数, 示例值: "group", 可选值有: group: 群聊, p2p: 单聊, 默认值: `group`
}

// GetMessageSpecialFocusUnreadResp ...
type GetMessageSpecialFocusUnreadResp struct {
	SpecialFocusUnread []*GetMessageSpecialFocusUnreadRespSpecialFocusUnread `json:"special_focus_unread,omitempty"` // 特别关注中未读的消息数
}

// GetMessageSpecialFocusUnreadRespSpecialFocusUnread ...
type GetMessageSpecialFocusUnreadRespSpecialFocusUnread struct {
	ID          string `json:"id,omitempty"`           // 成员ID
	IDType      IDType `json:"id_type,omitempty"`      // 成员ID类型。根据 [member_id_type] 参数返回`open_id`、`user_id`或`union_id`类型的用户ID；机器人返回`app_id`, 可选值有: user_id: 以user_id来识别用户；需要有获取用户 userID的权限 ([什么是 User ID？](https://open.feishu.cn/document/home/user-identity-introduction/user-id)), union_id: 以union_id来识别用户([什么是 Union ID？](https://open.feishu.cn/document/home/user-identity-introduction/union-id)), open_id: 以open_id来识别用户([什么是 Open ID？](https://open.feishu.cn/document/home/user-identity-introduction/open-id)), app_id: 以app_id来识别成员([什么是App ID？](https://open.feishu.cn/document/ukTMukTMukTM/uYTM5UjL2ETO14iNxkTN/terminology#b047be0c))
	UnreadCount string `json:"unread_count,omitempty"` // 未读数
}

// getMessageSpecialFocusUnreadResp ...
type getMessageSpecialFocusUnreadResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *GetMessageSpecialFocusUnreadResp `json:"data,omitempty"`
}
