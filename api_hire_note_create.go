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

// CreateHireNote 创建备注信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/note/create
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/candidate-management/note/create
func (r *HireService) CreateHireNote(ctx context.Context, request *CreateHireNoteReq, options ...MethodOptionFunc) (*CreateHireNoteResp, *Response, error) {
	if r.cli.mock.mockHireCreateHireNote != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#CreateHireNote mock enable")
		return r.cli.mock.mockHireCreateHireNote(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "CreateHireNote",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/notes",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createHireNoteResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireCreateHireNote mock HireCreateHireNote method
func (r *Mock) MockHireCreateHireNote(f func(ctx context.Context, request *CreateHireNoteReq, options ...MethodOptionFunc) (*CreateHireNoteResp, *Response, error)) {
	r.mockHireCreateHireNote = f
}

// UnMockHireCreateHireNote un-mock HireCreateHireNote method
func (r *Mock) UnMockHireCreateHireNote() {
	r.mockHireCreateHireNote = nil
}

// CreateHireNoteReq ...
type CreateHireNoteReq struct {
	UserIDType    *IDType `query:"user_id_type" json:"-"`   // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以people_admin_id来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	TalentID      string  `json:"talent_id,omitempty"`      // 人才ID, 示例值: "6916472453069883661"
	ApplicationID *string `json:"application_id,omitempty"` // 投递ID, 示例值: "6891565253964859661"
	CreatorID     *string `json:"creator_id,omitempty"`     // 创建人ID, 示例值: "ou_f476cb099ac9227c9bae09ce46112579"
	Content       string  `json:"content,omitempty"`        // 内容, 示例值: "测试"
	Privacy       *int64  `json:"privacy,omitempty"`        // 备注私密属性（默认为公开）, 示例值: 1, 可选值有: 1: 私密, 2: 公开
}

// CreateHireNoteResp ...
type CreateHireNoteResp struct {
	Note *CreateHireNoteRespNote `json:"note,omitempty"` // 备注信息
}

// CreateHireNoteRespNote ...
type CreateHireNoteRespNote struct {
	ID            string `json:"id,omitempty"`             // 备注ID
	TalentID      string `json:"talent_id,omitempty"`      // 人才ID
	ApplicationID string `json:"application_id,omitempty"` // 投递ID
	IsPrivate     bool   `json:"is_private,omitempty"`     // 是否私密
	CreateTime    int64  `json:"create_time,omitempty"`    // 创建时间
	ModifyTime    int64  `json:"modify_time,omitempty"`    // 更新时间
	CreatorID     string `json:"creator_id,omitempty"`     // 创建人ID
	Content       string `json:"content,omitempty"`        // 内容
}

// createHireNoteResp ...
type createHireNoteResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *CreateHireNoteResp `json:"data,omitempty"`
}
