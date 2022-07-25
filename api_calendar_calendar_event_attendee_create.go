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

// CreateCalendarEventAttendee 批量给日程添加参与人。
//
// - 当前身份需要有日历的 writer 或 owner 权限, 并且日历的类型只能为 primary 或 shared。
// - 当前身份需要是日程的组织者, 或日程设置了「参与人可邀请其它参与人」权限。
// - 新添加的日程参与人必须与日程组织者在同一个企业内。
// - 使用该接口添加会议室后, 会议室会进入异步的预约流程, 请求结束不代表会议室预约成功, 需后续再查询预约状态。
// - 每个日程最多只能有 3000 名参与人。
// - 开启管理员能力后预约会议室可不受会议室预约范围的限制（当前不支持用管理员身份给其他人的日程预约会议室）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/create
func (r *CalendarService) CreateCalendarEventAttendee(ctx context.Context, request *CreateCalendarEventAttendeeReq, options ...MethodOptionFunc) (*CreateCalendarEventAttendeeResp, *Response, error) {
	if r.cli.mock.mockCalendarCreateCalendarEventAttendee != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#CreateCalendarEventAttendee mock enable")
		return r.cli.mock.mockCalendarCreateCalendarEventAttendee(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "CreateCalendarEventAttendee",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createCalendarEventAttendeeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarCreateCalendarEventAttendee mock CalendarCreateCalendarEventAttendee method
func (r *Mock) MockCalendarCreateCalendarEventAttendee(f func(ctx context.Context, request *CreateCalendarEventAttendeeReq, options ...MethodOptionFunc) (*CreateCalendarEventAttendeeResp, *Response, error)) {
	r.mockCalendarCreateCalendarEventAttendee = f
}

// UnMockCalendarCreateCalendarEventAttendee un-mock CalendarCreateCalendarEventAttendee method
func (r *Mock) UnMockCalendarCreateCalendarEventAttendee() {
	r.mockCalendarCreateCalendarEventAttendee = nil
}

// CreateCalendarEventAttendeeReq ...
type CreateCalendarEventAttendeeReq struct {
	CalendarID             string                                    `path:"calendar_id" json:"-"`                // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值: "feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID                string                                    `path:"event_id" json:"-"`                   // 日程ID。参见[日程ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/introduction), 示例值: "xxxxxxxxx_0"
	UserIDType             *IDType                                   `query:"user_id_type" json:"-"`              // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Attendees              []*CreateCalendarEventAttendeeReqAttendee `json:"attendees,omitempty"`                 // 新增参与人列表；, 单次请求会议室的数量限制为100。
	NeedNotification       *bool                                     `json:"need_notification,omitempty"`         // 是否给参与人发送bot通知 默认为true, 示例值: false
	InstanceStartTimeAdmin *string                                   `json:"instance_start_time_admin,omitempty"` // 使用管理员身份访问时要修改的实例(仅用于重复日程修改其中的一个实例, 非重复日程无需填此字段), 示例值: "1647320400"
	IsEnableAdmin          *bool                                     `json:"is_enable_admin,omitempty"`           // 是否启用管理员身份(需先在管理后台设置某人为会议室管理员), 示例值: false
}

// CreateCalendarEventAttendeeReqAttendee ...
type CreateCalendarEventAttendeeReqAttendee struct {
	Type                  *CalendarEventAttendeeType                                     `json:"type,omitempty"`                   // 参与人类型, 示例值: "user", 可选值有: user: 用户, chat: 群组, resource: 会议室, third_party: 邮箱
	IsOptional            *bool                                                          `json:"is_optional,omitempty"`            // 参与人是否为「可选参加」, 无法编辑群参与人的此字段, 示例值: true, 默认值: `false`
	UserID                *string                                                        `json:"user_id,omitempty"`                // 参与人的用户id, 依赖于user_id_type返回对应的取值, 当is_external为true时, 此字段只会返回open_id或者union_id, 参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 示例值: "ou_xxxxxxxx"
	ChatID                *string                                                        `json:"chat_id,omitempty"`                // chat类型参与人的群组chat_id, 参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: "oc_xxxxxxxxx"
	RoomID                *string                                                        `json:"room_id,omitempty"`                // resource类型参与人的会议室room_id, 示例值: "omm_xxxxxxxx"
	ThirdPartyEmail       *string                                                        `json:"third_party_email,omitempty"`      // third_party类型参与人的邮箱, 示例值: "wangwu@email.com"
	OperateID             *string                                                        `json:"operate_id,omitempty"`             // 如果日程是使用应用身份创建的, 在添加会议室的时候, 用来指定会议室的联系人, 在会议室视图展示。参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 示例值: "ou_xxxxxxxx"
	ResourceCustomization []*CreateCalendarEventAttendeeReqAttendeeResourceCustomization `json:"resource_customization,omitempty"` // 会议室的个性化配置
}

// CreateCalendarEventAttendeeReqAttendeeResourceCustomization ...
type CreateCalendarEventAttendeeReqAttendeeResourceCustomization struct {
	IndexKey     string                                                               `json:"index_key,omitempty"`     // 每个配置的唯一ID, 示例值: "16281481596100"
	InputContent *string                                                              `json:"input_content,omitempty"` // 当type类型为填空时, 该参数需要填入, 示例值: "xxx"
	Options      []*CreateCalendarEventAttendeeReqAttendeeResourceCustomizationOption `json:"options,omitempty"`       // 每个配置的选项
}

// CreateCalendarEventAttendeeReqAttendeeResourceCustomizationOption ...
type CreateCalendarEventAttendeeReqAttendeeResourceCustomizationOption struct {
	OptionKey     *string `json:"option_key,omitempty"`     // 每个选项的唯一ID, 示例值: "16281481596185"
	OthersContent *string `json:"others_content,omitempty"` // 当type类型为其它选项时, 该参数需要填入, 示例值: "xxx"
}

// CreateCalendarEventAttendeeResp ...
type CreateCalendarEventAttendeeResp struct {
	Attendees []*CreateCalendarEventAttendeeRespAttendee `json:"attendees,omitempty"` // 新增参与人后的日程所有参与人列表
}

// CreateCalendarEventAttendeeRespAttendee ...
type CreateCalendarEventAttendeeRespAttendee struct {
	Type                  CalendarEventAttendeeType                                       `json:"type,omitempty"`                   // 参与人类型, 可选值有: user: 用户, chat: 群组, resource: 会议室, third_party: 邮箱
	AttendeeID            string                                                          `json:"attendee_id,omitempty"`            // 参与人ID。参见[参与人ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/introduction#4998889c)
	RsvpStatus            string                                                          `json:"rsvp_status,omitempty"`            // 参与人RSVP状态, 可选值有: needs_action: 参与人尚未回复状态, 或表示会议室预约中, accept: 参与人回复接受, 或表示会议室预约成功, tentative: 参与人回复待定, decline: 参与人回复拒绝, 或表示会议室预约失败, removed: 参与人或会议室已经从日程中被移除
	IsOptional            bool                                                            `json:"is_optional,omitempty"`            // 参与人是否为「可选参加」, 无法编辑群参与人的此字段
	IsOrganizer           bool                                                            `json:"is_organizer,omitempty"`           // 参与人是否为日程组织者
	IsExternal            bool                                                            `json:"is_external,omitempty"`            // 参与人是否为外部参与人；外部参与人不支持编辑
	DisplayName           string                                                          `json:"display_name,omitempty"`           // 参与人名称
	ChatMembers           []*CreateCalendarEventAttendeeRespAttendeeChatMember            `json:"chat_members,omitempty"`           // 群中的群成员, 当type为Chat时有效；群成员不支持编辑
	UserID                string                                                          `json:"user_id,omitempty"`                // 参与人的用户id, 依赖于user_id_type返回对应的取值, 当is_external为true时, 此字段只会返回open_id或者union_id, 参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	ChatID                string                                                          `json:"chat_id,omitempty"`                // chat类型参与人的群组chat_id, 参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description)
	RoomID                string                                                          `json:"room_id,omitempty"`                // resource类型参与人的会议室room_id
	ThirdPartyEmail       string                                                          `json:"third_party_email,omitempty"`      // third_party类型参与人的邮箱
	OperateID             string                                                          `json:"operate_id,omitempty"`             // 如果日程是使用应用身份创建的, 在添加会议室的时候, 用来指定会议室的联系人, 在会议室视图展示。参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	ResourceCustomization []*CreateCalendarEventAttendeeRespAttendeeResourceCustomization `json:"resource_customization,omitempty"` // 会议室的个性化配置
}

// CreateCalendarEventAttendeeRespAttendeeChatMember ...
type CreateCalendarEventAttendeeRespAttendeeChatMember struct {
	RsvpStatus  string `json:"rsvp_status,omitempty"`  // 参与人RSVP状态, 可选值有: needs_action: 参与人尚未回复状态, 或表示会议室预约中, accept: 参与人回复接受, 或表示会议室预约成功, tentative: 参与人回复待定, decline: 参与人回复拒绝, 或表示会议室预约失败, removed: 参与人或会议室已经从日程中被移除
	IsOptional  bool   `json:"is_optional,omitempty"`  // 参与人是否为「可选参加」
	DisplayName string `json:"display_name,omitempty"` // 参与人名称
	IsOrganizer bool   `json:"is_organizer,omitempty"` // 参与人是否为日程组织者
	IsExternal  bool   `json:"is_external,omitempty"`  // 参与人是否为外部参与人
}

// CreateCalendarEventAttendeeRespAttendeeResourceCustomization ...
type CreateCalendarEventAttendeeRespAttendeeResourceCustomization struct {
	IndexKey     string                                                                `json:"index_key,omitempty"`     // 每个配置的唯一ID
	InputContent string                                                                `json:"input_content,omitempty"` // 当type类型为填空时, 该参数需要填入
	Options      []*CreateCalendarEventAttendeeRespAttendeeResourceCustomizationOption `json:"options,omitempty"`       // 每个配置的选项
}

// CreateCalendarEventAttendeeRespAttendeeResourceCustomizationOption ...
type CreateCalendarEventAttendeeRespAttendeeResourceCustomizationOption struct {
	OptionKey     string `json:"option_key,omitempty"`     // 每个选项的唯一ID
	OthersContent string `json:"others_content,omitempty"` // 当type类型为其它选项时, 该参数需要填入
}

// createCalendarEventAttendeeResp ...
type createCalendarEventAttendeeResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *CreateCalendarEventAttendeeResp `json:"data,omitempty"`
}
