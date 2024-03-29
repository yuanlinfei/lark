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

// GetMessageList 获取会话（包括单聊、群组）的历史消息（聊天记录）。
//
// 接口级别权限默认只能获取单聊（p2p）消息, 如果需要获取群组（group）消息, 应用还必须拥有 [获取群组中所有消息] 权限
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 获取消息时, 机器人必须在群组中
// - 对于普通对话群中的话题消息, 通过 "chat" 容器类型仅能获取到话题的根消息, 可通过指定容器类型为 "thread" 获取话题回复中的所有消息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/list
// new doc: https://open.feishu.cn/document/server-docs/im-v1/message/list
func (r *MessageService) GetMessageList(ctx context.Context, request *GetMessageListReq, options ...MethodOptionFunc) (*GetMessageListResp, *Response, error) {
	if r.cli.mock.mockMessageGetMessageList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Message#GetMessageList mock enable")
		return r.cli.mock.mockMessageGetMessageList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "GetMessageList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMessageListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageGetMessageList mock MessageGetMessageList method
func (r *Mock) MockMessageGetMessageList(f func(ctx context.Context, request *GetMessageListReq, options ...MethodOptionFunc) (*GetMessageListResp, *Response, error)) {
	r.mockMessageGetMessageList = f
}

// UnMockMessageGetMessageList un-mock MessageGetMessageList method
func (r *Mock) UnMockMessageGetMessageList() {
	r.mockMessageGetMessageList = nil
}

// GetMessageListReq ...
type GetMessageListReq struct {
	ContainerIDType ContainerIDType `query:"container_id_type" json:"-"` // 容器类型, 可选值有: `chat`: 包含单聊（p2p）和群聊（group）, `thread`: 话题, 示例值: chat
	ContainerID     string          `query:"container_id" json:"-"`      // 容器的ID, 可填写: chat_id, 群聊或单聊的 ID, 参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), thread_id, 话题 ID, 参见[话题介绍](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/thread-introduction), 示例值: oc_234jsi43d3ssi993d43545f
	StartTime       *string         `query:"start_time" json:"-"`        // 历史信息的起始时间（秒级时间戳）, 注意: thread 容器类型暂不支持获取指定时间范围内的消息, 示例值: 1608594809
	EndTime         *string         `query:"end_time" json:"-"`          // 历史信息的结束时间（秒级时间戳）, 注意: thread 容器类型暂不支持获取指定时间范围内的消息, 示例值: 1609296809
	SortType        *string         `query:"sort_type" json:"-"`         // 消息排序方式, 示例值: ByCreateTimeAsc, 可选值有: ByCreateTimeAsc: 按消息创建时间升序排列, ByCreateTimeDesc: 按消息创建时间降序排列, 默认值: `ByCreateTimeAsc`
	PageSize        *int64          `query:"page_size" json:"-"`         // 分页大小, 示例值: 20, 默认值: `20`, 取值范围: `1` ～ `50`
	PageToken       *string         `query:"page_token" json:"-"`        // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ[
}

// GetMessageListResp ...
type GetMessageListResp struct {
	HasMore   bool                      `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                    `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetMessageListRespItem `json:"items,omitempty"`      // message[]
}

// GetMessageListRespItem ...
type GetMessageListRespItem struct {
	MessageID      string       `json:"message_id,omitempty"`       // 消息id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	RootID         string       `json:"root_id,omitempty"`          // 根消息id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	ParentID       string       `json:"parent_id,omitempty"`        // 父消息的id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	ThreadID       string       `json:"thread_id,omitempty"`        // 消息所属的话题 ID
	MsgType        MsgType      `json:"msg_type,omitempty"`         // 消息类型 包括: text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等, 类型定义请参考[接收消息内容](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content)
	CreateTime     string       `json:"create_time,omitempty"`      // 消息生成的时间戳（毫秒）
	UpdateTime     string       `json:"update_time,omitempty"`      // 消息更新的时间戳（毫秒）
	Deleted        bool         `json:"deleted,omitempty"`          // 消息是否被撤回或删除
	Updated        bool         `json:"updated,omitempty"`          // 消息是否被更新
	ChatID         string       `json:"chat_id,omitempty"`          // 所属的群
	Sender         *Sender      `json:"sender,omitempty"`           // 发送者, 可以是用户或应用
	Body           *MessageBody `json:"body,omitempty"`             // 消息内容
	Mentions       []*Mention   `json:"mentions,omitempty"`         // 被@的用户或机器人的id列表
	UpperMessageID string       `json:"upper_message_id,omitempty"` // 合并转发消息中, 上一层级的消息id message_id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
}

// getMessageListResp ...
type getMessageListResp struct {
	Code  int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string              `json:"msg,omitempty"`  // 错误描述
	Data  *GetMessageListResp `json:"data,omitempty"`
	Error *ErrorDetail        `json:"error,omitempty"`
}
