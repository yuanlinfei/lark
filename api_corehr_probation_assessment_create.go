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

// CreateCoreHRProbationAssessment 新增员工试用期考核结果
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/probation-assessment/create
func (r *CoreHRService) CreateCoreHRProbationAssessment(ctx context.Context, request *CreateCoreHRProbationAssessmentReq, options ...MethodOptionFunc) (*CreateCoreHRProbationAssessmentResp, *Response, error) {
	if r.cli.mock.mockCoreHRCreateCoreHRProbationAssessment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHR#CreateCoreHRProbationAssessment mock enable")
		return r.cli.mock.mockCoreHRCreateCoreHRProbationAssessment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "CreateCoreHRProbationAssessment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/probation/assessments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHRProbationAssessmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRCreateCoreHRProbationAssessment mock CoreHRCreateCoreHRProbationAssessment method
func (r *Mock) MockCoreHRCreateCoreHRProbationAssessment(f func(ctx context.Context, request *CreateCoreHRProbationAssessmentReq, options ...MethodOptionFunc) (*CreateCoreHRProbationAssessmentResp, *Response, error)) {
	r.mockCoreHRCreateCoreHRProbationAssessment = f
}

// UnMockCoreHRCreateCoreHRProbationAssessment un-mock CoreHRCreateCoreHRProbationAssessment method
func (r *Mock) UnMockCoreHRCreateCoreHRProbationAssessment() {
	r.mockCoreHRCreateCoreHRProbationAssessment = nil
}

// CreateCoreHRProbationAssessmentReq ...
type CreateCoreHRProbationAssessmentReq struct {
	ClientToken  *string                                         `query:"client_token" json:"-"`  // 根据 client_token 是否一致来判断是否为同一请求, 示例值: 6822122262122064111
	UserIDType   *IDType                                         `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	EmploymentID string                                          `json:"employment_id,omitempty"` // 试用期人员的雇佣 ID, 示例值: "7140964208476371111"
	Assessments  []*CreateCoreHRProbationAssessmentReqAssessment `json:"assessments,omitempty"`   // 试用期考核结果列表
}

// CreateCoreHRProbationAssessmentReqAssessment ...
type CreateCoreHRProbationAssessmentReqAssessment struct {
	AssessmentStatus   string   `json:"assessment_status,omitempty"`    // 考核状态, 示例值: "completed", 可选值有: not_started: 未开始, in_process: 进行中, completed: 已完成, no_need: 无需考核
	AssessmentResult   *string  `json:"assessment_result,omitempty"`    // 试用期考核结果, 示例值: "approved", 可选值有: approved: 通过, rejected: 不通过
	AssessmentScore    *float64 `json:"assessment_score,omitempty"`     // 考核得分, 示例值: 99.9
	AssessmentGrade    *string  `json:"assessment_grade,omitempty"`     // 试用期考核等级, 枚举值 api_name 可通过[【获取字段详情】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/get_by_param)接口查询, 查询参数如下: object_api_name = "probation_management", custom_api_name = "final_assessment_grade", 示例值: "grade_a"
	AssessmentComment  *string  `json:"assessment_comment,omitempty"`   // 考核评语, 示例值: "超出预期"
	AssessmentDetail   *string  `json:"assessment_detail,omitempty"`    // 考核结果页面超链接, 示例值: "暂无示例"
	IsFinalAsssessment bool     `json:"is_final_asssessment,omitempty"` // 是否为最终考核结果, 示例值: false
}

// CreateCoreHRProbationAssessmentResp ...
type CreateCoreHRProbationAssessmentResp struct {
	AssessmentIDs []string `json:"assessment_ids,omitempty"` // 创建的试用期考核记录 ID 列表, 有序返回
}

// createCoreHRProbationAssessmentResp ...
type createCoreHRProbationAssessmentResp struct {
	Code int64                                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                               `json:"msg,omitempty"`  // 错误描述
	Data *CreateCoreHRProbationAssessmentResp `json:"data,omitempty"`
}
