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

// SetVCScopeConfig 该接口可以用来设置某个会议层级范围下或者某个会议室的配置。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/scope_config/create
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/scope_config/create
func (r *VCService) SetVCScopeConfig(ctx context.Context, request *SetVCScopeConfigReq, options ...MethodOptionFunc) (*SetVCScopeConfigResp, *Response, error) {
	if r.cli.mock.mockVCSetVCScopeConfig != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] VC#SetVCScopeConfig mock enable")
		return r.cli.mock.mockVCSetVCScopeConfig(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "SetVCScopeConfig",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/scope_config",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(setVCScopeConfigResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCSetVCScopeConfig mock VCSetVCScopeConfig method
func (r *Mock) MockVCSetVCScopeConfig(f func(ctx context.Context, request *SetVCScopeConfigReq, options ...MethodOptionFunc) (*SetVCScopeConfigResp, *Response, error)) {
	r.mockVCSetVCScopeConfig = f
}

// UnMockVCSetVCScopeConfig un-mock VCSetVCScopeConfig method
func (r *Mock) UnMockVCSetVCScopeConfig() {
	r.mockVCSetVCScopeConfig = nil
}

// SetVCScopeConfigReq ...
type SetVCScopeConfigReq struct {
	UserIDType  *IDType                         `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 长度范围: `0` ～ `10` 字符, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ScopeType   int64                           `json:"scope_type,omitempty"`   // 查询节点范围, 示例值: 1, 可选值有: 1: 会议室层级, 2: 会议室
	ScopeID     string                          `json:"scope_id,omitempty"`     // 查询节点ID: 如果scope_type为1, 则为层级ID, 如果scope_type为2, 则为会议室ID, 示例值: "omm_608d34d82d531b27fa993902d350a307"
	ScopeConfig *SetVCScopeConfigReqScopeConfig `json:"scope_config,omitempty"` // 节点配置
}

// SetVCScopeConfigReqScopeConfig ...
type SetVCScopeConfigReqScopeConfig struct {
	RoomBackground        *string                                              `json:"room_background,omitempty"`          // 飞书会议室背景图, 示例值: "https://lf1-ttcdn-tos.pstatp.com/obj/xxx"
	DisplayBackground     *string                                              `json:"display_background,omitempty"`       // 飞书签到板背景图, 示例值: "https://lf1-ttcdn-tos.pstatp.com/obj/xxx"
	DigitalSignage        *SetVCScopeConfigReqScopeConfigDigitalSignage        `json:"digital_signage,omitempty"`          // 飞书会议室数字标牌
	RoomBoxDigitalSignage *SetVCScopeConfigReqScopeConfigRoomBoxDigitalSignage `json:"room_box_digital_signage,omitempty"` // 飞书投屏盒子数字标牌
	RoomStatus            *SetVCScopeConfigReqScopeConfigRoomStatus            `json:"room_status,omitempty"`              // 会议室状态
}

// SetVCScopeConfigReqScopeConfigDigitalSignage ...
type SetVCScopeConfigReqScopeConfigDigitalSignage struct {
	IfCoverChildScope *bool                                                   `json:"if_cover_child_scope,omitempty"` // 是否覆盖子层级及会议室, 示例值: true
	Enable            *bool                                                   `json:"enable,omitempty"`               // 是否开启数字标牌功能, 示例值: true
	Mute              *bool                                                   `json:"mute,omitempty"`                 // 是否静音播放, 示例值: true
	StartDisplay      *int64                                                  `json:"start_display,omitempty"`        // 在会议结束n分钟后开始播放, 取值1~720（仅对飞书会议室数字标牌生效）, 示例值: 3
	StopDisplay       *int64                                                  `json:"stop_display,omitempty"`         // 在日程会议开始前n分钟停止播放, 取值1~720（仅对飞书会议室数字标牌生效）, 示例值: 3
	Materials         []*SetVCScopeConfigReqScopeConfigDigitalSignageMaterial `json:"materials,omitempty"`            // 素材列表
}

// SetVCScopeConfigReqScopeConfigDigitalSignageMaterial ...
type SetVCScopeConfigReqScopeConfigDigitalSignageMaterial struct {
	ID           *string `json:"id,omitempty"`            // 素材ID, 当设置新素材时, 无需传递该字段, 示例值: "7847784676276"
	Name         *string `json:"name,omitempty"`          // 素材名称, 示例值: "name"
	MaterialType *int64  `json:"material_type,omitempty"` // 素材类型, 示例值: 0, 可选值有: 1: 图片, 2: 视频, 3: GIF
	URL          *string `json:"url,omitempty"`           // 素材url, 示例值: "url"
	Duration     *int64  `json:"duration,omitempty"`      // 播放时长（单位sec）, 取值1~43200, 示例值: 15
	Cover        *string `json:"cover,omitempty"`         // 素材封面url, 示例值: "url"
	Md5          *string `json:"md5,omitempty"`           // 素材文件md5, 示例值: "md5"
	Vid          *string `json:"vid,omitempty"`           // 素材文件vid, 示例值: "vid"
	Size         *string `json:"size,omitempty"`          // 素材文件大小（单位byte）, 示例值: "100"
}

// SetVCScopeConfigReqScopeConfigRoomBoxDigitalSignage ...
type SetVCScopeConfigReqScopeConfigRoomBoxDigitalSignage struct {
	IfCoverChildScope *bool                                                          `json:"if_cover_child_scope,omitempty"` // 是否覆盖子层级及会议室, 示例值: true
	Enable            *bool                                                          `json:"enable,omitempty"`               // 是否开启数字标牌功能, 示例值: true
	Mute              *bool                                                          `json:"mute,omitempty"`                 // 是否静音播放, 示例值: true
	StartDisplay      *int64                                                         `json:"start_display,omitempty"`        // 在会议结束n分钟后开始播放, 取值1~720（仅对飞书会议室数字标牌生效）, 示例值: 3
	StopDisplay       *int64                                                         `json:"stop_display,omitempty"`         // 在日程会议开始前n分钟停止播放, 取值1~720（仅对飞书会议室数字标牌生效）, 示例值: 3
	Materials         []*SetVCScopeConfigReqScopeConfigRoomBoxDigitalSignageMaterial `json:"materials,omitempty"`            // 素材列表
}

// SetVCScopeConfigReqScopeConfigRoomBoxDigitalSignageMaterial ...
type SetVCScopeConfigReqScopeConfigRoomBoxDigitalSignageMaterial struct {
	ID           *string `json:"id,omitempty"`            // 素材ID, 当设置新素材时, 无需传递该字段, 示例值: "7847784676276"
	Name         *string `json:"name,omitempty"`          // 素材名称, 示例值: "name"
	MaterialType *int64  `json:"material_type,omitempty"` // 素材类型, 示例值: 0, 可选值有: 1: 图片, 2: 视频, 3: GIF
	URL          *string `json:"url,omitempty"`           // 素材url, 示例值: "url"
	Duration     *int64  `json:"duration,omitempty"`      // 播放时长（单位sec）, 取值1~43200, 示例值: 15
	Cover        *string `json:"cover,omitempty"`         // 素材封面url, 示例值: "url"
	Md5          *string `json:"md5,omitempty"`           // 素材文件md5, 示例值: "md5"
	Vid          *string `json:"vid,omitempty"`           // 素材文件vid, 示例值: "v039b2g10000ca89uj3c77u5pfdkfvpg"
	Size         *string `json:"size,omitempty"`          // 素材文件大小（单位byte）, 示例值: "100"
}

// SetVCScopeConfigReqScopeConfigRoomStatus ...
type SetVCScopeConfigReqScopeConfigRoomStatus struct {
	Status           bool     `json:"status,omitempty"`             // 是否启用会议室, 示例值: true
	ScheduleStatus   *bool    `json:"schedule_status,omitempty"`    // 会议室未来状态为启用或禁用（请忽略, 该字段用于查询接口的返回值）, 示例值: true
	DisableStartTime *string  `json:"disable_start_time,omitempty"` // 禁用开始时间（unix时间, 单位sec）, 示例值: "1652356050"
	DisableEndTime   *string  `json:"disable_end_time,omitempty"`   // 禁用结束时间（unix时间, 单位sec, 数值0表示永久禁用）, 示例值: "1652442450"
	DisableReason    *string  `json:"disable_reason,omitempty"`     // 禁用原因, 示例值: "测试占用"
	ContactIDs       []string `json:"contact_ids,omitempty"`        // 联系人列表, id类型由user_id_type参数决定, 示例值: ["ou_3ec3f6a28a0d08c45d895276e8e5e19b"]
	DisableNotice    *bool    `json:"disable_notice,omitempty"`     // 是否在禁用时发送通知给预定了该会议室的员工, 示例值: true
	ResumeNotice     *bool    `json:"resume_notice,omitempty"`      // 是否在恢复启用时发送通知给联系人, 示例值: true
}

// SetVCScopeConfigResp ...
type SetVCScopeConfigResp struct {
}

// setVCScopeConfigResp ...
type setVCScopeConfigResp struct {
	Code  int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                `json:"msg,omitempty"`  // 错误描述
	Data  *SetVCScopeConfigResp `json:"data,omitempty"`
	Error *ErrorDetail          `json:"error,omitempty"`
}
