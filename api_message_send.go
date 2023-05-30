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

// SendRawMessage 给指定用户或者会话发送消息, 支持文本、富文本、可交互的[消息卡片](https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN)、群名片、个人名片、图片、视频、音频、文件、表情包。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 给用户发送消息, 需要机器人对用户有[可用性](https://open.feishu.cn/document/home/introduction-to-scope-and-authorization/availability)
// - 给群组发送消息, 需要机器人在群组中
// - 为避免对用户造成打扰, 向同一用户发送消息的限频为 [5 QPS], 向同一群组发送消息的限频为群内机器人共享 [5 QPS]
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create
func (r *MessageService) SendRawMessage(ctx context.Context, request *SendRawMessageReq, options ...MethodOptionFunc) (*SendRawMessageResp, *Response, error) {
	if r.cli.mock.mockMessageSendRawMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#SendRawMessage mock enable")
		return r.cli.mock.mockMessageSendRawMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "SendRawMessage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(sendRawMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageSendRawMessage mock MessageSendRawMessage method
func (r *Mock) MockMessageSendRawMessage(f func(ctx context.Context, request *SendRawMessageReq, options ...MethodOptionFunc) (*SendRawMessageResp, *Response, error)) {
	r.mockMessageSendRawMessage = f
}

// UnMockMessageSendRawMessage un-mock MessageSendRawMessage method
func (r *Mock) UnMockMessageSendRawMessage() {
	r.mockMessageSendRawMessage = nil
}

// SendRawMessageReq ...
type SendRawMessageReq struct {
	ReceiveIDType IDType  `query:"receive_id_type" json:"-"` // 消息接收者id类型 open_id/user_id/union_id/email/chat_id, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), email: 以用户的真实邮箱来标识用户。, chat_id: 以群ID来标识群聊。[了解更多: 如何获取群ID ](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description)
	ReceiveID     string  `json:"receive_id,omitempty"`      // 消息接收者的ID, ID类型应与查询参数[receive_id_type] 对应；推荐使用 OpenID, 获取方式可参考文档[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), 示例值: "ou_7d8a6e6df7621556ce0d21922b676706ccs"
	MsgType       MsgType `json:"msg_type,omitempty"`        // 消息类型 包括: text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等, 类型定义请参考[发送消息Content](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json), 示例值: "text"
	Content       string  `json:"content,omitempty"`         // 消息内容, JSON结构序列化后的字符串。不同msg_type对应不同内容, 具体格式说明参考: [发送消息Content](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json), 注意: JSON字符串需进行转义, 如换行符转义后为`\\n`, 文本消息请求体最大不能超过150KB, 卡片及富文本消息请求体最大不能超过30KB, 示例值: "{\"text\":\"test content\"}"
	UUID          *string `json:"uuid,omitempty"`            // 由开发者生成的唯一字符串序列, 用于发送消息请求去重；持有相同uuid的请求1小时内至多成功发送一条消息, 示例值: "a0d69e20-1dd1-458b-k525-dfeca4015204", 最大长度: `50` 字符
}

// SendRawMessageResp ...
type SendRawMessageResp struct {
	MessageID      string       `json:"message_id,omitempty"`       // 消息id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	RootID         string       `json:"root_id,omitempty"`          // 根消息id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	ParentID       string       `json:"parent_id,omitempty"`        // 父消息的id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
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

// sendRawMessageResp ...
type sendRawMessageResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *SendRawMessageResp `json:"data,omitempty"`
}
