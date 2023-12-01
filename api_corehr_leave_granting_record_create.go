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

// CreateCoreHRLeaveGrantingRecord 向飞书人事休假系统写入假期发放记录。
//
// 仅飞书人事企业版可用
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/leave_granting_record/create
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/leave/create
func (r *CoreHRService) CreateCoreHRLeaveGrantingRecord(ctx context.Context, request *CreateCoreHRLeaveGrantingRecordReq, options ...MethodOptionFunc) (*CreateCoreHRLeaveGrantingRecordResp, *Response, error) {
	if r.cli.mock.mockCoreHRCreateCoreHRLeaveGrantingRecord != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHR#CreateCoreHRLeaveGrantingRecord mock enable")
		return r.cli.mock.mockCoreHRCreateCoreHRLeaveGrantingRecord(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "CreateCoreHRLeaveGrantingRecord",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/leave_granting_records",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHRLeaveGrantingRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRCreateCoreHRLeaveGrantingRecord mock CoreHRCreateCoreHRLeaveGrantingRecord method
func (r *Mock) MockCoreHRCreateCoreHRLeaveGrantingRecord(f func(ctx context.Context, request *CreateCoreHRLeaveGrantingRecordReq, options ...MethodOptionFunc) (*CreateCoreHRLeaveGrantingRecordResp, *Response, error)) {
	r.mockCoreHRCreateCoreHRLeaveGrantingRecord = f
}

// UnMockCoreHRCreateCoreHRLeaveGrantingRecord un-mock CoreHRCreateCoreHRLeaveGrantingRecord method
func (r *Mock) UnMockCoreHRCreateCoreHRLeaveGrantingRecord() {
	r.mockCoreHRCreateCoreHRLeaveGrantingRecord = nil
}

// CreateCoreHRLeaveGrantingRecordReq ...
type CreateCoreHRLeaveGrantingRecordReq struct {
	UserIDType       *IDType                                     `query:"user_id_type" json:"-"`      // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的ID来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	LeaveTypeID      string                                      `json:"leave_type_id,omitempty"`     // 假期类型 ID, 枚举值可通过【获取假期类型列表】接口获取（若假期类型下存在假期子类, 此处仅支持传入假期子类的 ID）, 示例值: "7111688079785723436"
	EmploymentID     string                                      `json:"employment_id,omitempty"`     // 员工 ID, 示例值: "6982509313466189342"
	GrantingQuantity string                                      `json:"granting_quantity,omitempty"` // 发放数量, 示例值: "0.5"
	GrantingUnit     int64                                       `json:"granting_unit,omitempty"`     // 发放时长单位, 可选值有: 1: 天, 2: 小时, 示例值: 1
	EffectiveDate    string                                      `json:"effective_date,omitempty"`    // 生效时间, 示例值: "2022-01-01"
	Reason           []*CreateCoreHRLeaveGrantingRecordReqReason `json:"reason,omitempty"`            // 发放原因
	ExternalID       *string                                     `json:"external_id,omitempty"`       // 自定义外部 ID, 可用于避免数据重复写入（不能超过 64 字符）, 示例值: "111"
}

// CreateCoreHRLeaveGrantingRecordReqReason ...
type CreateCoreHRLeaveGrantingRecordReqReason struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// CreateCoreHRLeaveGrantingRecordResp ...
type CreateCoreHRLeaveGrantingRecordResp struct {
	LeaveGrantingRecord *CreateCoreHRLeaveGrantingRecordRespLeaveGrantingRecord `json:"leave_granting_record,omitempty"` // 假期发放记录
}

// CreateCoreHRLeaveGrantingRecordRespLeaveGrantingRecord ...
type CreateCoreHRLeaveGrantingRecordRespLeaveGrantingRecord struct {
	ID               string                                                          `json:"id,omitempty"`                // 假期发放记录 ID
	EmploymentID     string                                                          `json:"employment_id,omitempty"`     // 员工 ID
	LeaveTypeID      string                                                          `json:"leave_type_id,omitempty"`     // 假期类型 ID
	GrantingQuantity string                                                          `json:"granting_quantity,omitempty"` // 发放数量
	GrantingUnit     int64                                                           `json:"granting_unit,omitempty"`     // 发放时长单位, 可选值有: 1: 天, 2: 小时
	EffectiveDate    string                                                          `json:"effective_date,omitempty"`    // 生效时间
	ExpirationDate   string                                                          `json:"expiration_date,omitempty"`   // 失效时间（根据休假规则自动计算）
	GrantedBy        int64                                                           `json:"granted_by,omitempty"`        // 发放来源, 可选值有: 1: 系统发放, 2: 手动发放, 3: 外部系统发放
	Reason           []*CreateCoreHRLeaveGrantingRecordRespLeaveGrantingRecordReason `json:"reason,omitempty"`            // 发放原因
	CreatedAt        string                                                          `json:"created_at,omitempty"`        // 发放记录的创建时间
	CreatedBy        string                                                          `json:"created_by,omitempty"`        // 发放记录的创建人, 值为创建人的员工 ID
	UpdatedAt        string                                                          `json:"updated_at,omitempty"`        // 发放记录的更新时间
	UpdatedBy        string                                                          `json:"updated_by,omitempty"`        // 发放记录的更新人, 值为更新人的员工 ID
}

// CreateCoreHRLeaveGrantingRecordRespLeaveGrantingRecordReason ...
type CreateCoreHRLeaveGrantingRecordRespLeaveGrantingRecordReason struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// createCoreHRLeaveGrantingRecordResp ...
type createCoreHRLeaveGrantingRecordResp struct {
	Code int64                                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                               `json:"msg,omitempty"`  // 错误描述
	Data *CreateCoreHRLeaveGrantingRecordResp `json:"data,omitempty"`
}
