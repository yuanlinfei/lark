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

// SearchCoreHRProbation 搜索试用期信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/probation/search
func (r *CoreHRService) SearchCoreHRProbation(ctx context.Context, request *SearchCoreHRProbationReq, options ...MethodOptionFunc) (*SearchCoreHRProbationResp, *Response, error) {
	if r.cli.mock.mockCoreHRSearchCoreHRProbation != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHR#SearchCoreHRProbation mock enable")
		return r.cli.mock.mockCoreHRSearchCoreHRProbation(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "SearchCoreHRProbation",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/probation/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchCoreHRProbationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRSearchCoreHRProbation mock CoreHRSearchCoreHRProbation method
func (r *Mock) MockCoreHRSearchCoreHRProbation(f func(ctx context.Context, request *SearchCoreHRProbationReq, options ...MethodOptionFunc) (*SearchCoreHRProbationResp, *Response, error)) {
	r.mockCoreHRSearchCoreHRProbation = f
}

// UnMockCoreHRSearchCoreHRProbation un-mock CoreHRSearchCoreHRProbation method
func (r *Mock) UnMockCoreHRSearchCoreHRProbation() {
	r.mockCoreHRSearchCoreHRProbation = nil
}

// SearchCoreHRProbationReq ...
type SearchCoreHRProbationReq struct {
	PageSize                      int64             `query:"page_size" json:"-"`                         // 分页大小, 最大 100, 示例值: 100, 取值范围: `1` ～ `100`
	PageToken                     *string           `query:"page_token" json:"-"`                        // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 6891251722631890445
	UserIDType                    *IDType           `query:"user_id_type" json:"-"`                      // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType              *DepartmentIDType `query:"department_id_type" json:"-"`                // 此次调用中使用的部门 ID 类型, 示例值: open_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `open_department_id`
	EmploymentIDs                 []string          `json:"employment_ids,omitempty"`                    // 雇佣 ID 列表, 示例值: ["7140964208476371111"]
	DepartmentIDs                 []string          `json:"department_ids,omitempty"`                    // 部门 ID 列表, 示例值: ["7140964208476371121"]
	ProbationStartDateStart       *string           `json:"probation_start_date_start,omitempty"`        // 试用期开始日期 - 搜索范围开始, 需要与搜索范围结束一同使用, 示例值: "2022-05-18"
	ProbationStartDateEnd         *string           `json:"probation_start_date_end,omitempty"`          // 试用期开始日期 - 搜索范围结束, 示例值: "2022-05-20"
	ProbationExpectedEndDateStart *string           `json:"probation_expected_end_date_start,omitempty"` // 试用期预计结束日期 - 搜索范围开始, 需要与搜索范围结束一同使用, 示例值: "2022-06-20"
	ProbationExpectedEndDateEnd   *string           `json:"probation_expected_end_date_end,omitempty"`   // 试用期预计结束日期 - 搜索范围结束, 示例值: "2022-07-20"
	ActualProbationEndDateStart   *string           `json:"actual_probation_end_date_start,omitempty"`   // 试用期实际结束日期 - 搜索范围开始, 需要与搜索范围结束一同使用, 示例值: "2022-08-20"
	ActualProbationEndDateEnd     *string           `json:"actual_probation_end_date_end,omitempty"`     // 试用期实际结束日期 - 搜索范围结束, 示例值: "2022-09-20"
	InitiatingTimeStart           *string           `json:"initiating_time_start,omitempty"`             // 转正发起日期 - 搜索范围开始, 需要与搜索范围结束一同使用, 示例值: "2022-10-20"
	InitiatingTimeEnd             *string           `json:"initiating_time_end,omitempty"`               // 转正发起日期 - 搜索范围结束, 示例值: "2022-11-20"
	ProbationStatus               *string           `json:"probation_status,omitempty"`                  // 试用期状态, 示例值: "approved", 可选值有: pending: 审批中, rejected: 已拒绝, waiting: 待发起转正, approved: 审批通过, converted: 已转正, offboarded: 已离职
	FinalAssessmentResult         *string           `json:"final_assessment_result,omitempty"`           // 试用期最终考核结果, 示例值: "approved", 可选值有: approved: 通过, rejected: 不通过
	FinalAssessmentGrade          *string           `json:"final_assessment_grade,omitempty"`            // 试用期最终考核等级, 枚举值 api_name 可通过[【获取字段详情】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/get_by_param)接口查询, 查询参数如下: object_api_name: probation_management, custom_api_name: final_assessment_grade, <b>字段权限要求: </b>, 按照试用期考核等级搜索 (corehr:probation.grade.search:read), 示例值: "grade_a"
}

// SearchCoreHRProbationResp ...
type SearchCoreHRProbationResp struct {
	Items     []*SearchCoreHRProbationRespItem `json:"items,omitempty"`      // 查询的试用期信息
	PageToken string                           `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                             `json:"has_more,omitempty"`   // 是否还有更多项
}

// SearchCoreHRProbationRespItem ...
type SearchCoreHRProbationRespItem struct {
	EmploymentID             string                                              `json:"employment_id,omitempty"`               // 雇佣 ID
	ProbationID              string                                              `json:"probation_id,omitempty"`                // 试用期信息 ID
	ProbationStartDate       string                                              `json:"probation_start_date,omitempty"`        // 试用期开始日期
	ProbationExpectedEndDate string                                              `json:"probation_expected_end_date,omitempty"` // 试用期预计结束日期
	ActualProbationEndDate   string                                              `json:"actual_probation_end_date,omitempty"`   // 试用期实际结束日期
	InitiatingTime           string                                              `json:"initiating_time,omitempty"`             // 转正发起日期
	SubmissionType           *SearchCoreHRProbationRespItemSubmissionType        `json:"submission_type,omitempty"`             // 发起方, hr_submission: HR 发起, self_submission: 员工发起, system : 系统发起
	InitiatorID              string                                              `json:"initiator_id,omitempty"`                // 转正发起人的雇佣 ID, 当系统发起转正时该字段为空
	ProbationStatus          *SearchCoreHRProbationRespItemProbationStatus       `json:"probation_status,omitempty"`            // 试用期状态: pending: 审批中, rejected: 已拒绝, waiting: 待发起转正, approved: 审批通过, converted: 已转正, offboarded: 已离职
	SelfReview               string                                              `json:"self_review,omitempty"`                 // 员工自评, 字段权限要求（满足任一）: 读取员工自评信息, 读写员工自评信息
	Notes                    string                                              `json:"notes,omitempty"`                       // 备注, 字段权限要求（满足任一）: 读取试用期备注信息, 读写试用期备注信息
	ProcessID                string                                              `json:"process_id,omitempty"`                  // 流程实例 ID
	ConvertedViaBpm          bool                                                `json:"converted_via_bpm,omitempty"`           // 是否通过 BPM 转正
	CustomFields             []*SearchCoreHRProbationRespItemCustomField         `json:"custom_fields,omitempty"`               // 自定义字段, 字段权限要求（满足任一）: 获取试用期自定义字段信息, 读写试用期自定义字段信息
	FinalAssessmentStatus    *SearchCoreHRProbationRespItemFinalAssessmentStatus `json:"final_assessment_status,omitempty"`     // 试用期考核最终状态: not_started: 未开始, in_process: 进行中, completed: 已完成, no_need: 无需考核
	FinalAssessmentResult    *SearchCoreHRProbationRespItemFinalAssessmentResult `json:"final_assessment_result,omitempty"`     // 试用期考核最终结果: approved: 通过, rejected: 不通过
	FinalAssessmentScore     float64                                             `json:"final_assessment_score,omitempty"`      // 试用期考核最终得分, 字段权限要求（满足任一）: 获取试用期考核信息, 读写试用期考核信息
	FinalAssessmentGrade     *SearchCoreHRProbationRespItemFinalAssessmentGrade  `json:"final_assessment_grade,omitempty"`      // 试用期考核最终等级, 枚举值 api_name 可通过[【获取字段详情】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/get_by_param)接口查询, 查询参数如下: object_api_name = "probation_management", custom_api_name = "final_assessment_grade", 字段权限要求（满足任一）: 获取试用期考核信息, 读写试用期考核信息
	FinalAssessmentComment   string                                              `json:"final_assessment_comment,omitempty"`    // 试用期考核最终评语, 字段权限要求（满足任一）: 获取试用期考核信息, 读写试用期考核信息
	FinalAssessmentDetail    string                                              `json:"final_assessment_detail,omitempty"`     // 最终考核结果页面超链接, 字段权限要求（满足任一）: 获取试用期考核信息, 读写试用期考核信息
	Assessments              []*SearchCoreHRProbationRespItemAssessment          `json:"assessments,omitempty"`                 // 试用期考核结果列表, 字段权限要求（满足任一）: 获取试用期考核信息, 读写试用期考核信息
}

// SearchCoreHRProbationRespItemAssessment ...
type SearchCoreHRProbationRespItemAssessment struct {
	AssessmentID       string                                                   `json:"assessment_id,omitempty"`        // 考核结果 ID
	AssessmentStatus   *SearchCoreHRProbationRespItemAssessmentAssessmentStatus `json:"assessment_status,omitempty"`    // 考核状态, not_started: 未开始, in_process: 进行中, completed: 已完成, no_need: 无需考核
	AssessmentResult   *SearchCoreHRProbationRespItemAssessmentAssessmentResult `json:"assessment_result,omitempty"`    // 试用期考核结果, approved: 通过, rejected: 不通过
	AssessmentScore    float64                                                  `json:"assessment_score,omitempty"`     // 考核得分
	AssessmentGrade    *SearchCoreHRProbationRespItemAssessmentAssessmentGrade  `json:"assessment_grade,omitempty"`     // 试用期考核等级, 枚举值 api_name 可通过[【获取字段详情】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/get_by_param)接口查询, 查询参数如下: object_api_name = "probation_management", custom_api_name = "final_assessment_grade"
	AssessmentComment  string                                                   `json:"assessment_comment,omitempty"`   // 考核评语
	AssessmentDetail   string                                                   `json:"assessment_detail,omitempty"`    // 考核结果页面超链接
	IsFinalAsssessment bool                                                     `json:"is_final_asssessment,omitempty"` // 是否为最终考核结果
}

// SearchCoreHRProbationRespItemAssessmentAssessmentGrade ...
type SearchCoreHRProbationRespItemAssessmentAssessmentGrade struct {
	EnumName string                                                           `json:"enum_name,omitempty"` // 枚举值
	Display  []*SearchCoreHRProbationRespItemAssessmentAssessmentGradeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// SearchCoreHRProbationRespItemAssessmentAssessmentGradeDisplay ...
type SearchCoreHRProbationRespItemAssessmentAssessmentGradeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// SearchCoreHRProbationRespItemAssessmentAssessmentResult ...
type SearchCoreHRProbationRespItemAssessmentAssessmentResult struct {
	EnumName string                                                            `json:"enum_name,omitempty"` // 枚举值
	Display  []*SearchCoreHRProbationRespItemAssessmentAssessmentResultDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// SearchCoreHRProbationRespItemAssessmentAssessmentResultDisplay ...
type SearchCoreHRProbationRespItemAssessmentAssessmentResultDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// SearchCoreHRProbationRespItemAssessmentAssessmentStatus ...
type SearchCoreHRProbationRespItemAssessmentAssessmentStatus struct {
	EnumName string                                                            `json:"enum_name,omitempty"` // 枚举值
	Display  []*SearchCoreHRProbationRespItemAssessmentAssessmentStatusDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// SearchCoreHRProbationRespItemAssessmentAssessmentStatusDisplay ...
type SearchCoreHRProbationRespItemAssessmentAssessmentStatusDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// SearchCoreHRProbationRespItemCustomField ...
type SearchCoreHRProbationRespItemCustomField struct {
	CustomApiName string                                        `json:"custom_api_name,omitempty"` // 自定义字段 apiname, 即自定义字段的唯一标识
	Name          *SearchCoreHRProbationRespItemCustomFieldName `json:"name,omitempty"`            // 自定义字段名称
	Type          int64                                         `json:"type,omitempty"`            // 自定义字段类型
	Value         string                                        `json:"value,omitempty"`           // 字段值, 是 json 转义后的字符串, 根据元数据定义不同, 字段格式不同（如 123, 123.23, "true", ["id1", "id2"], "2006-01-02 15:04:05"）
}

// SearchCoreHRProbationRespItemCustomFieldName ...
type SearchCoreHRProbationRespItemCustomFieldName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// SearchCoreHRProbationRespItemFinalAssessmentGrade ...
type SearchCoreHRProbationRespItemFinalAssessmentGrade struct {
	EnumName string                                                      `json:"enum_name,omitempty"` // 枚举值
	Display  []*SearchCoreHRProbationRespItemFinalAssessmentGradeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// SearchCoreHRProbationRespItemFinalAssessmentGradeDisplay ...
type SearchCoreHRProbationRespItemFinalAssessmentGradeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// SearchCoreHRProbationRespItemFinalAssessmentResult ...
type SearchCoreHRProbationRespItemFinalAssessmentResult struct {
	EnumName string                                                       `json:"enum_name,omitempty"` // 枚举值
	Display  []*SearchCoreHRProbationRespItemFinalAssessmentResultDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// SearchCoreHRProbationRespItemFinalAssessmentResultDisplay ...
type SearchCoreHRProbationRespItemFinalAssessmentResultDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// SearchCoreHRProbationRespItemFinalAssessmentStatus ...
type SearchCoreHRProbationRespItemFinalAssessmentStatus struct {
	EnumName string                                                       `json:"enum_name,omitempty"` // 枚举值
	Display  []*SearchCoreHRProbationRespItemFinalAssessmentStatusDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// SearchCoreHRProbationRespItemFinalAssessmentStatusDisplay ...
type SearchCoreHRProbationRespItemFinalAssessmentStatusDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// SearchCoreHRProbationRespItemProbationStatus ...
type SearchCoreHRProbationRespItemProbationStatus struct {
	EnumName string                                                 `json:"enum_name,omitempty"` // 枚举值
	Display  []*SearchCoreHRProbationRespItemProbationStatusDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// SearchCoreHRProbationRespItemProbationStatusDisplay ...
type SearchCoreHRProbationRespItemProbationStatusDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// SearchCoreHRProbationRespItemSubmissionType ...
type SearchCoreHRProbationRespItemSubmissionType struct {
	EnumName string                                                `json:"enum_name,omitempty"` // 枚举值
	Display  []*SearchCoreHRProbationRespItemSubmissionTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// SearchCoreHRProbationRespItemSubmissionTypeDisplay ...
type SearchCoreHRProbationRespItemSubmissionTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// searchCoreHRProbationResp ...
type searchCoreHRProbationResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *SearchCoreHRProbationResp `json:"data,omitempty"`
}
