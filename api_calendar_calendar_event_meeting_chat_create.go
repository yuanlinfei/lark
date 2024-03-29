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

// CreateCalendarEventMeetingChat 该接口用于以当前身份（应用/用户）给日程创建一个会议群。
//
// 身份由 Header Authorization 的 Token 类型决定。
// - 日历需要是当前身份的主日历, 且具有writer权限。
// - 日程至少需要2个参与人, 且不隐藏参与人列表
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-meeting_chat/create
func (r *CalendarService) CreateCalendarEventMeetingChat(ctx context.Context, request *CreateCalendarEventMeetingChatReq, options ...MethodOptionFunc) (*CreateCalendarEventMeetingChatResp, *Response, error) {
	if r.cli.mock.mockCalendarCreateCalendarEventMeetingChat != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Calendar#CreateCalendarEventMeetingChat mock enable")
		return r.cli.mock.mockCalendarCreateCalendarEventMeetingChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "CreateCalendarEventMeetingChat",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/meeting_chat",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createCalendarEventMeetingChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarCreateCalendarEventMeetingChat mock CalendarCreateCalendarEventMeetingChat method
func (r *Mock) MockCalendarCreateCalendarEventMeetingChat(f func(ctx context.Context, request *CreateCalendarEventMeetingChatReq, options ...MethodOptionFunc) (*CreateCalendarEventMeetingChatResp, *Response, error)) {
	r.mockCalendarCreateCalendarEventMeetingChat = f
}

// UnMockCalendarCreateCalendarEventMeetingChat un-mock CalendarCreateCalendarEventMeetingChat method
func (r *Mock) UnMockCalendarCreateCalendarEventMeetingChat() {
	r.mockCalendarCreateCalendarEventMeetingChat = nil
}

// CreateCalendarEventMeetingChatReq ...
type CreateCalendarEventMeetingChatReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID, 示例值: "feishu.cn_xxx@group.calendar.feishu.cn"
	EventID    string `path:"event_id" json:"-"`    // 日程ID, 示例值: "75d28f9b-e35c-4230-8a83-123_0"
}

// CreateCalendarEventMeetingChatResp ...
type CreateCalendarEventMeetingChatResp struct {
	MeetingChatID string `json:"meeting_chat_id,omitempty"` // 会议群ID
	Applink       string `json:"applink,omitempty"`         // 群分享链接
}

// createCalendarEventMeetingChatResp ...
type createCalendarEventMeetingChatResp struct {
	Code  int64                               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                              `json:"msg,omitempty"`  // 错误描述
	Data  *CreateCalendarEventMeetingChatResp `json:"data,omitempty"`
	Error *ErrorDetail                        `json:"error,omitempty"`
}
