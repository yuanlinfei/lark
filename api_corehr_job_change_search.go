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

// SearchCoreHRJobChange 搜索异动信息, 该接口会按照应用拥有的「员工数据」的权限范围返回数据, 请确定在「开发者后台 - 权限管理 - 数据权限」中有申请「员工资源」权限范围
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/job_change/search
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/job_change/search
func (r *CoreHRService) SearchCoreHRJobChange(ctx context.Context, request *SearchCoreHRJobChangeReq, options ...MethodOptionFunc) (*SearchCoreHRJobChangeResp, *Response, error) {
	if r.cli.mock.mockCoreHRSearchCoreHRJobChange != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#SearchCoreHRJobChange mock enable")
		return r.cli.mock.mockCoreHRSearchCoreHRJobChange(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "SearchCoreHRJobChange",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/job_changes/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchCoreHRJobChangeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRSearchCoreHRJobChange mock CoreHRSearchCoreHRJobChange method
func (r *Mock) MockCoreHRSearchCoreHRJobChange(f func(ctx context.Context, request *SearchCoreHRJobChangeReq, options ...MethodOptionFunc) (*SearchCoreHRJobChangeResp, *Response, error)) {
	r.mockCoreHRSearchCoreHRJobChange = f
}

// UnMockCoreHRSearchCoreHRJobChange un-mock CoreHRSearchCoreHRJobChange method
func (r *Mock) UnMockCoreHRSearchCoreHRJobChange() {
	r.mockCoreHRSearchCoreHRJobChange = nil
}

// SearchCoreHRJobChangeReq ...
type SearchCoreHRJobChangeReq struct {
	PageSize           int64             `query:"page_size" json:"-"`            // 分页大小, 最大 100, 示例值: 100, 取值范围: `1` ～ `100`
	PageToken          *string           `query:"page_token" json:"-"`           // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 6891251722631890445
	UserIDType         *IDType           `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType   *DepartmentIDType `query:"department_id_type" json:"-"`   // 此次调用中使用的部门 ID 类型, 示例值: open_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `open_department_id`
	EmploymentIDs      []string          `json:"employment_ids,omitempty"`       // 雇员 ID 列表, 示例值: ["ou_a294793e8fa21529f2a60e3e9de45520"], 最大长度: `30`
	JobChangeIDs       []string          `json:"job_change_ids,omitempty"`       // 异动记录 ID 列表, 示例值: ["7044427347159746085"], 最大长度: `10`
	Statuses           []string          `json:"statuses,omitempty"`             // 异动状态, 多个状态之间为「或」的关系, 示例值: ["Approving"], 可选值有: Approving: Approving  审批中, Approved: Approved  审批通过, Transformed: Transformed  已异动, Rejected: Rejected  已拒绝, Cancelled: Cancelled  已撤销, NoNeedApproval: NoNeedApproval  无需审批, 最大长度: `10`
	EffectiveDateStart *string           `json:"effective_date_start,omitempty"` // 异动生效日期-搜索范围开始, 需要与搜索范围结束一同使用, 示例值: "2022-01-01"
	EffectiveDateEnd   *string           `json:"effective_date_end,omitempty"`   // 异动生效日期 - 搜索范围结束, 示例值: "2022-01-01"
}

// SearchCoreHRJobChangeResp ...
type SearchCoreHRJobChangeResp struct {
	Items     []*SearchCoreHRJobChangeRespItem `json:"items,omitempty"`      // 员工异动列表
	HasMore   bool                             `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                           `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// SearchCoreHRJobChangeRespItem ...
type SearchCoreHRJobChangeRespItem struct {
	JobChangeID                    string                                     `json:"job_change_id,omitempty"`                     // 异动记录 id
	EmploymentID                   string                                     `json:"employment_id,omitempty"`                     // 雇员 id
	Status                         string                                     `json:"status,omitempty"`                            // 异动状态, 可选值有: Approving: 审批中, Approved: 审批通过, Transformed: 已异动, Rejected: 已拒绝, Cancelled: 已撤销, NoNeedApproval: 无需审批
	TransferTypeUniqueIdentifier   string                                     `json:"transfer_type_unique_identifier,omitempty"`   // 异动类型
	TransferReasonUniqueIdentifier string                                     `json:"transfer_reason_unique_identifier,omitempty"` // 异动原因
	ProcessID                      string                                     `json:"process_id,omitempty"`                        // 异动流程 id
	EffectiveDate                  string                                     `json:"effective_date,omitempty"`                    // 生效时间
	CreatedTime                    string                                     `json:"created_time,omitempty"`                      // 创建时间
	UpdatedTime                    string                                     `json:"updated_time,omitempty"`                      // 更新时间
	TransferInfo                   *SearchCoreHRJobChangeRespItemTransferInfo `json:"transfer_info,omitempty"`                     // 异动详细信息
}

// SearchCoreHRJobChangeRespItemTransferInfo ...
type SearchCoreHRJobChangeRespItemTransferInfo struct {
	Remark                     string                                                             `json:"remark,omitempty"`                        // 备注, 字段权限要求: 获取异动流程备注信息
	OfferInfo                  string                                                             `json:"offer_info,omitempty"`                    // offer信息
	TargetDottedManagerClean   bool                                                               `json:"target_dotted_manager_clean,omitempty"`   // 是否撤销虚线上级
	ProbationExist             bool                                                               `json:"probation_exist,omitempty"`               // 是否有试用期
	OriginalDepartment         string                                                             `json:"original_department,omitempty"`           // 原部门
	TargetDepartment           string                                                             `json:"target_department,omitempty"`             // 新部门
	OriginalWorkLocation       string                                                             `json:"original_work_location,omitempty"`        // 原工作地点
	TargetWorkLocation         string                                                             `json:"target_work_location,omitempty"`          // 新工作地点
	OriginalDirectManager      string                                                             `json:"original_direct_manager,omitempty"`       // 原直属上级
	TargetDirectManager        string                                                             `json:"target_direct_manager,omitempty"`         // 新直属上级
	OriginalDottedManager      string                                                             `json:"original_dotted_manager,omitempty"`       // 原虚线上级
	TargetDottedManager        string                                                             `json:"target_dotted_manager,omitempty"`         // 新虚线上级
	OriginalJob                string                                                             `json:"original_job,omitempty"`                  // 原职务, 字段权限要求（满足任一）: 获取职务级别信息, 读写员工的职务级别信息
	TargetJob                  string                                                             `json:"target_job,omitempty"`                    // 新职务, 字段权限要求（满足任一）: 获取职务级别信息, 读写员工的职务级别信息
	OriginalJobFamily          string                                                             `json:"original_job_family,omitempty"`           // 原序列
	TargetJobFamily            string                                                             `json:"target_job_family,omitempty"`             // 新序列
	OriginalJobLevel           string                                                             `json:"original_job_level,omitempty"`            // 原级别, 字段权限要求（满足任一）: 获取职务级别信息, 读写员工的职务级别信息
	TargetJobLevel             string                                                             `json:"target_job_level,omitempty"`              // 新级别, 字段权限要求（满足任一）: 获取职务级别信息, 读写员工的职务级别信息
	OriginalWorkforceType      string                                                             `json:"original_workforce_type,omitempty"`       // 原人员类型
	TargetWorkforceType        string                                                             `json:"target_workforce_type,omitempty"`         // 新人员类型
	OriginalCompany            string                                                             `json:"original_company,omitempty"`              // 原公司, 字段权限要求（满足任一）: 获取合同主体信息, 读写合同主体信息
	TargetCompany              string                                                             `json:"target_company,omitempty"`                // 新公司, 字段权限要求（满足任一）: 获取合同主体信息, 读写合同主体信息
	OriginalContractNumber     string                                                             `json:"original_contract_number,omitempty"`      // 原合同编号
	TargetContractNumber       string                                                             `json:"target_contract_number,omitempty"`        // 新合同编号
	OriginalContractType       string                                                             `json:"original_contract_type,omitempty"`        // 原合同类型
	TargetContractType         string                                                             `json:"target_contract_type,omitempty"`          // 新合同类型
	OriginalDurationType       string                                                             `json:"original_duration_type,omitempty"`        // 原期限类型
	TargetDurationType         string                                                             `json:"target_duration_type,omitempty"`          // 新期限类型
	OriginalSigningType        string                                                             `json:"original_signing_type,omitempty"`         // 原签订类型
	TargetSigningType          string                                                             `json:"target_signing_type,omitempty"`           // 新签订类型
	OriginalContractStartDate  string                                                             `json:"original_contract_start_date,omitempty"`  // 原合同开始日期, 字段权限要求（满足任一）: 获取合同期限信息, 读写合同期限信息
	TargetContractStartDate    string                                                             `json:"target_contract_start_date,omitempty"`    // 新合同开始日期, 字段权限要求（满足任一）: 获取合同期限信息, 读写合同期限信息
	OriginalContractEndDate    string                                                             `json:"original_contract_end_date,omitempty"`    // 原合同结束日期, 字段权限要求（满足任一）: 获取合同期限信息, 读写合同期限信息
	TargetContractEndDate      string                                                             `json:"target_contract_end_date,omitempty"`      // 新合同结束日期, 字段权限要求（满足任一）: 获取合同期限信息, 读写合同期限信息
	OriginalWorkingHoursType   string                                                             `json:"original_working_hours_type,omitempty"`   // 原工时制度
	TargetWorkingHoursType     string                                                             `json:"target_working_hours_type,omitempty"`     // 新工时制度
	OriginalWorkingCalendar    string                                                             `json:"original_working_calendar,omitempty"`     // 原工作日历
	TargetWorkingCalendar      string                                                             `json:"target_working_calendar,omitempty"`       // 新工作日历
	OriginalProbationEndDate   string                                                             `json:"original_probation_end_date,omitempty"`   // 原试用期预计结束日期
	TargetProbationEndDate     string                                                             `json:"target_probation_end_date,omitempty"`     // 新试用期预计结束日期
	OriginalWeeklyWorkingHours string                                                             `json:"original_weekly_working_hours,omitempty"` // 原周工作时长
	TargetWeeklyWorkingHours   string                                                             `json:"target_weekly_working_hours,omitempty"`   // 新周工作时长
	OriginalWorkShift          string                                                             `json:"original_work_shift,omitempty"`           // 原排班
	TargetWorkShift            string                                                             `json:"target_work_shift,omitempty"`             // 新排班
	OriginalCostCenterRate     []*SearchCoreHRJobChangeRespItemTransferInfoOriginalCostCenterRate `json:"original_cost_center_rate,omitempty"`     // 原成本中心分摊方式
	TargetCostCenterRate       []*SearchCoreHRJobChangeRespItemTransferInfoTargetCostCenterRate   `json:"target_cost_center_rate,omitempty"`       // 新成本中心分摊方式
	OriginalEmploymentChange   *SearchCoreHRJobChangeRespItemTransferInfoOriginalEmploymentChange `json:"original_employment_change,omitempty"`    // 原工作信息
	TargetEmploymentChange     *SearchCoreHRJobChangeRespItemTransferInfoTargetEmploymentChange   `json:"target_employment_change,omitempty"`      // 新工作信息
	OriginalJobGrade           string                                                             `json:"original_job_grade,omitempty"`            // 原职等, 字段权限要求（满足任一）: 获取职等信息, 读写职等信息
	TargetJobGrade             string                                                             `json:"target_job_grade,omitempty"`              // 新职等, 字段权限要求（满足任一）: 获取职等信息, 读写职等信息
}

// SearchCoreHRJobChangeRespItemTransferInfoOriginalCostCenterRate ...
type SearchCoreHRJobChangeRespItemTransferInfoOriginalCostCenterRate struct {
	CostCenterID string `json:"cost_center_id,omitempty"` // 成本中心 ID, 可以通过【查询单个成本中心信息】接口获取对应的成本中心信息
	Rate         int64  `json:"rate,omitempty"`           // 分摊比例
}

// SearchCoreHRJobChangeRespItemTransferInfoOriginalEmploymentChange ...
type SearchCoreHRJobChangeRespItemTransferInfoOriginalEmploymentChange struct {
	RegularEmployeeStartDate string                                                                          `json:"regular_employee_start_date,omitempty"` // 转正式员工日期
	SeniorityDate            string                                                                          `json:"seniority_date,omitempty"`              // 司龄起算日期
	EmployeeNumber           string                                                                          `json:"employee_number,omitempty"`             // 员工编号
	CustomFields             []*SearchCoreHRJobChangeRespItemTransferInfoOriginalEmploymentChangeCustomField `json:"custom_fields,omitempty"`               // 自定义字段, 字段权限要求: 获取异动工作信息自定义字段信息
}

// SearchCoreHRJobChangeRespItemTransferInfoOriginalEmploymentChangeCustomField ...
type SearchCoreHRJobChangeRespItemTransferInfoOriginalEmploymentChangeCustomField struct {
	CustomApiName string                                                                            `json:"custom_api_name,omitempty"` // 自定义字段 apiname, 即自定义字段的唯一标识
	Name          *SearchCoreHRJobChangeRespItemTransferInfoOriginalEmploymentChangeCustomFieldName `json:"name,omitempty"`            // 自定义字段名称
	Type          int64                                                                             `json:"type,omitempty"`            // 自定义字段类型
	Value         string                                                                            `json:"value,omitempty"`           // 字段值, 是 json 转义后的字符串, 根据元数据定义不同, 字段格式不同（如 123, 123.23, "true", ["id1", "id2"], "2006-01-02 15:04:05"）
}

// SearchCoreHRJobChangeRespItemTransferInfoOriginalEmploymentChangeCustomFieldName ...
type SearchCoreHRJobChangeRespItemTransferInfoOriginalEmploymentChangeCustomFieldName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// SearchCoreHRJobChangeRespItemTransferInfoTargetCostCenterRate ...
type SearchCoreHRJobChangeRespItemTransferInfoTargetCostCenterRate struct {
	CostCenterID string `json:"cost_center_id,omitempty"` // 成本中心 ID, 可以通过【查询单个成本中心信息】接口获取对应的成本中心信息
	Rate         int64  `json:"rate,omitempty"`           // 分摊比例
}

// SearchCoreHRJobChangeRespItemTransferInfoTargetEmploymentChange ...
type SearchCoreHRJobChangeRespItemTransferInfoTargetEmploymentChange struct {
	RegularEmployeeStartDate string                                                                        `json:"regular_employee_start_date,omitempty"` // 转正式员工日期
	SeniorityDate            string                                                                        `json:"seniority_date,omitempty"`              // 司龄起算日期
	EmployeeNumber           string                                                                        `json:"employee_number,omitempty"`             // 员工编号
	CustomFields             []*SearchCoreHRJobChangeRespItemTransferInfoTargetEmploymentChangeCustomField `json:"custom_fields,omitempty"`               // 自定义字段, 字段权限要求: 获取异动工作信息自定义字段信息
}

// SearchCoreHRJobChangeRespItemTransferInfoTargetEmploymentChangeCustomField ...
type SearchCoreHRJobChangeRespItemTransferInfoTargetEmploymentChangeCustomField struct {
	CustomApiName string                                                                          `json:"custom_api_name,omitempty"` // 自定义字段 apiname, 即自定义字段的唯一标识
	Name          *SearchCoreHRJobChangeRespItemTransferInfoTargetEmploymentChangeCustomFieldName `json:"name,omitempty"`            // 自定义字段名称
	Type          int64                                                                           `json:"type,omitempty"`            // 自定义字段类型
	Value         string                                                                          `json:"value,omitempty"`           // 字段值, 是 json 转义后的字符串, 根据元数据定义不同, 字段格式不同（如 123, 123.23, "true", ["id1", "id2"], "2006-01-02 15:04:05"）
}

// SearchCoreHRJobChangeRespItemTransferInfoTargetEmploymentChangeCustomFieldName ...
type SearchCoreHRJobChangeRespItemTransferInfoTargetEmploymentChangeCustomFieldName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// searchCoreHRJobChangeResp ...
type searchCoreHRJobChangeResp struct {
	Code  int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 错误描述
	Data  *SearchCoreHRJobChangeResp `json:"data,omitempty"`
	Error *ErrorDetail               `json:"error,omitempty"`
}
