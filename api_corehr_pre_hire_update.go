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

// UpdateCoreHRPreHire 更新待入职信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/pre_hire/patch
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/pre_hire/patch
func (r *CoreHRService) UpdateCoreHRPreHire(ctx context.Context, request *UpdateCoreHRPreHireReq, options ...MethodOptionFunc) (*UpdateCoreHRPreHireResp, *Response, error) {
	if r.cli.mock.mockCoreHRUpdateCoreHRPreHire != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#UpdateCoreHRPreHire mock enable")
		return r.cli.mock.mockCoreHRUpdateCoreHRPreHire(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "UpdateCoreHRPreHire",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/pre_hires/:pre_hire_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateCoreHRPreHireResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRUpdateCoreHRPreHire mock CoreHRUpdateCoreHRPreHire method
func (r *Mock) MockCoreHRUpdateCoreHRPreHire(f func(ctx context.Context, request *UpdateCoreHRPreHireReq, options ...MethodOptionFunc) (*UpdateCoreHRPreHireResp, *Response, error)) {
	r.mockCoreHRUpdateCoreHRPreHire = f
}

// UnMockCoreHRUpdateCoreHRPreHire un-mock CoreHRUpdateCoreHRPreHire method
func (r *Mock) UnMockCoreHRUpdateCoreHRPreHire() {
	r.mockCoreHRUpdateCoreHRPreHire = nil
}

// UpdateCoreHRPreHireReq ...
type UpdateCoreHRPreHireReq struct {
	PreHireID        string                                  `path:"pre_hire_id" json:"-"`         // 待入职ID, 示例值: "1616161616"
	ClientToken      *string                                 `query:"client_token" json:"-"`       // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	AtsApplicationID *string                                 `json:"ats_application_id,omitempty"` // 招聘投递 ID, 详细信息可以通过招聘的【获取投递信息】接口查询获得, 示例值: "4719168654814483759"
	HireDate         *string                                 `json:"hire_date,omitempty"`          // 入职日期, 示例值: "2020-01-01"
	EmployeeType     *UpdateCoreHRPreHireReqEmployeeType     `json:"employee_type,omitempty"`      // 雇佣类型
	WorkerID         *string                                 `json:"worker_id,omitempty"`          // 人员编号, 示例值: "1245646"
	EmployeeTypeID   *string                                 `json:"employee_type_id,omitempty"`   // 雇佣类型, 示例值: "正式"
	PersonID         *string                                 `json:"person_id,omitempty"`          // 引用Person ID, 示例值: "656464648662"
	CustomFields     []*UpdateCoreHRPreHireReqCustomField    `json:"custom_fields,omitempty"`      // 自定义字段
	CostCenterRate   []*UpdateCoreHRPreHireReqCostCenterRate `json:"cost_center_rate,omitempty"`   // 成本中心分摊信息
	OnboardingStatus *UpdateCoreHRPreHireReqOnboardingStatus `json:"onboarding_status,omitempty"`  // 入职状态, 待入职(preboarding), 已删除(deleted), 准备就绪(day_one), 已撤销(withdrawn), 已完成(completed)
}

// UpdateCoreHRPreHireReqCostCenterRate ...
type UpdateCoreHRPreHireReqCostCenterRate struct {
	CostCenterID *string `json:"cost_center_id,omitempty"` // 支持的成本中心id, 示例值: "6950635856373745165"
	Rate         *int64  `json:"rate,omitempty"`           // 分摊比例, 示例值: 100
}

// UpdateCoreHRPreHireReqCustomField ...
type UpdateCoreHRPreHireReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// UpdateCoreHRPreHireReqEmployeeType ...
type UpdateCoreHRPreHireReqEmployeeType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// UpdateCoreHRPreHireReqOnboardingStatus ...
type UpdateCoreHRPreHireReqOnboardingStatus struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// UpdateCoreHRPreHireResp ...
type UpdateCoreHRPreHireResp struct {
	PreHire *UpdateCoreHRPreHireRespPreHire `json:"pre_hire,omitempty"` // 待入职数据
}

// UpdateCoreHRPreHireRespPreHire ...
type UpdateCoreHRPreHireRespPreHire struct {
	AtsApplicationID string                                          `json:"ats_application_id,omitempty"` // 招聘投递 ID, 详细信息可以通过招聘的【获取投递信息】接口查询获得
	ID               string                                          `json:"id,omitempty"`                 // 待入职ID
	HireDate         string                                          `json:"hire_date,omitempty"`          // 入职日期
	EmployeeType     *UpdateCoreHRPreHireRespPreHireEmployeeType     `json:"employee_type,omitempty"`      // 雇佣类型
	WorkerID         string                                          `json:"worker_id,omitempty"`          // 人员编号
	EmployeeTypeID   string                                          `json:"employee_type_id,omitempty"`   // 雇佣类型
	PersonID         string                                          `json:"person_id,omitempty"`          // 引用Person ID
	CustomFields     []*UpdateCoreHRPreHireRespPreHireCustomField    `json:"custom_fields,omitempty"`      // 自定义字段
	CostCenterRate   []*UpdateCoreHRPreHireRespPreHireCostCenterRate `json:"cost_center_rate,omitempty"`   // 成本中心分摊信息
	OnboardingStatus *UpdateCoreHRPreHireRespPreHireOnboardingStatus `json:"onboarding_status,omitempty"`  // 入职状态, 待入职(preboarding), 已删除(deleted), 准备就绪(day_one), 已撤销(withdrawn), 已完成(completed)
}

// UpdateCoreHRPreHireRespPreHireCostCenterRate ...
type UpdateCoreHRPreHireRespPreHireCostCenterRate struct {
	CostCenterID string `json:"cost_center_id,omitempty"` // 支持的成本中心id
	Rate         int64  `json:"rate,omitempty"`           // 分摊比例
}

// UpdateCoreHRPreHireRespPreHireCustomField ...
type UpdateCoreHRPreHireRespPreHireCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// UpdateCoreHRPreHireRespPreHireEmployeeType ...
type UpdateCoreHRPreHireRespPreHireEmployeeType struct {
	EnumName string                                               `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRPreHireRespPreHireEmployeeTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRPreHireRespPreHireEmployeeTypeDisplay ...
type UpdateCoreHRPreHireRespPreHireEmployeeTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHRPreHireRespPreHireOnboardingStatus ...
type UpdateCoreHRPreHireRespPreHireOnboardingStatus struct {
	EnumName string                                                   `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRPreHireRespPreHireOnboardingStatusDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRPreHireRespPreHireOnboardingStatusDisplay ...
type UpdateCoreHRPreHireRespPreHireOnboardingStatusDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// updateCoreHRPreHireResp ...
type updateCoreHRPreHireResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateCoreHRPreHireResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
