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

// MergeForwardMessage 将来自同一个群聊中的多条消息合并转发给指定用、群聊或话题。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 向用户合并转发消息, 需要机器人对用户有[可用性](https://open.feishu.cn/document/home/introduction-to-scope-and-authorization/availability)
// - 向群组合并转发消息, 需要机器人在群组中
// - 对于要合并转发的消息与转发到的对象, 本接口有以下限制:
// - 不支持合并转发系统消息（system）
// - 不支持合并转发来自不同群聊中的消息
// - 不支持同时合并转发来自多个话题中的消息
// - 不支持同时合并转发普通消息与话题中的消息
// - 不支持再次合并转发合并转发消息中的子消息（含有[upper_message_id]字段的消息）
// - 合并转发生成的新消息的消息内容为固定值[Merged and Forwarded Message], 其子消息可以使用[获取指定消息的内容](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/get)接口获取
// - 为避免对用户造成打扰, 向同一用户发送消息的限频为 [5 QPS], 向同一群组发送消息的限频为群内机器人共享 [5 QPS]
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/merge_forward
// new doc: https://open.feishu.cn/document/server-docs/im-v1/message/merge_forward
func (r *MessageService) MergeForwardMessage(ctx context.Context, request *MergeForwardMessageReq, options ...MethodOptionFunc) (*MergeForwardMessageResp, *Response, error) {
	if r.cli.mock.mockMessageMergeForwardMessage != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Message#MergeForwardMessage mock enable")
		return r.cli.mock.mockMessageMergeForwardMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "MergeForwardMessage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/merge_forward",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(mergeForwardMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageMergeForwardMessage mock MessageMergeForwardMessage method
func (r *Mock) MockMessageMergeForwardMessage(f func(ctx context.Context, request *MergeForwardMessageReq, options ...MethodOptionFunc) (*MergeForwardMessageResp, *Response, error)) {
	r.mockMessageMergeForwardMessage = f
}

// UnMockMessageMergeForwardMessage un-mock MessageMergeForwardMessage method
func (r *Mock) UnMockMessageMergeForwardMessage() {
	r.mockMessageMergeForwardMessage = nil
}

// MergeForwardMessageReq ...
type MergeForwardMessageReq struct {
	ReceiveIDType IDType   `query:"receive_id_type" json:"-"` // 消息接收者id类型 open_id/user_id/union_id/email/chat_id, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), email: 以用户的真实邮箱来标识用户。, chat_id: 以群ID来标识群聊。[了解更多: 如何获取群ID ](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), thread_id: 以话题ID来标识话题。了解更多: [话题介绍](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/thread-introduction),
	UUID          *string  `query:"uuid" json:"-"`            // 由开发者生成的唯一字符串序列, 用于合并转发消息请求去重；持有相同uuid的请求在1小时内向同一个目标的合并转发只可成功一次, 示例值: b13g2t38-1jd2-458b-8djf-dtbca5104204, 最大长度: `50` 字符
	ReceiveID     string   `json:"receive_id,omitempty"`      // 依据`receive_id_type`的值, 填写对应的合并转发目标的ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
	MessageIDList []string `json:"message_id_list,omitempty"` // 要转发的消息ID列表, 详情参见[消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2), 示例值: ["om_dc13264520392913993dd051dba21dcf"], 长度范围: `1` ～ `100`
}

// MergeForwardMessageResp ...
type MergeForwardMessageResp struct {
	Message              *MergeForwardMessageRespMessage `json:"message,omitempty"`                 // 合并转发生成的新消息
	InvalidMessageIDList []string                        `json:"invalid_message_id_list,omitempty"` // 无效的消息ID列表, 如不存在的消息、已被撤回的消息、不可见的历史消息、不支持的消息类型等。
}

// MergeForwardMessageRespMessage ...
type MergeForwardMessageRespMessage struct {
	MessageID      string       `json:"message_id,omitempty"`       // 消息id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	RootID         string       `json:"root_id,omitempty"`          // 根消息id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	ParentID       string       `json:"parent_id,omitempty"`        // 父消息的id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	ThreadID       string       `json:"thread_id,omitempty"`        // 消息所属的话题 ID（不返回说明该消息非话题消息）, 说明参见: [话题介绍](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/thread-introduction)
	MsgType        MsgType      `json:"msg_type,omitempty"`         // 消息类型 包括: text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等, 类型定义请参考[接收消息Content](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content)
	CreateTime     string       `json:"create_time,omitempty"`      // 消息生成的时间戳（毫秒）
	UpdateTime     string       `json:"update_time,omitempty"`      // 消息更新的时间戳（毫秒）
	Deleted        bool         `json:"deleted,omitempty"`          // 消息是否被撤回
	Updated        bool         `json:"updated,omitempty"`          // 消息是否被更新
	ChatID         string       `json:"chat_id,omitempty"`          // 所属的群
	Sender         *Sender      `json:"sender,omitempty"`           // 发送者, 可以是用户或应用
	Body           *MessageBody `json:"body,omitempty"`             // 消息内容
	Mentions       []*Mention   `json:"mentions,omitempty"`         // 被@的用户或机器人的id列表
	UpperMessageID string       `json:"upper_message_id,omitempty"` // 合并转发消息中, 上一层级的消息id message_id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
}

// mergeForwardMessageResp ...
type mergeForwardMessageResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *MergeForwardMessageResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
