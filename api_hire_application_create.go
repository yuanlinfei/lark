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

// CreateHireApplication 根据人才 ID 和职位 ID 创建投递
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/create
func (r *HireService) CreateHireApplication(ctx context.Context, request *CreateHireApplicationReq, options ...MethodOptionFunc) (*CreateHireApplicationResp, *Response, error) {
	if r.cli.mock.mockHireCreateHireApplication != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#CreateHireApplication mock enable")
		return r.cli.mock.mockHireCreateHireApplication(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "CreateHireApplication",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/applications",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createHireApplicationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireCreateHireApplication mock HireCreateHireApplication method
func (r *Mock) MockHireCreateHireApplication(f func(ctx context.Context, request *CreateHireApplicationReq, options ...MethodOptionFunc) (*CreateHireApplicationResp, *Response, error)) {
	r.mockHireCreateHireApplication = f
}

// UnMockHireCreateHireApplication un-mock HireCreateHireApplication method
func (r *Mock) UnMockHireCreateHireApplication() {
	r.mockHireCreateHireApplication = nil
}

// CreateHireApplicationReq ...
type CreateHireApplicationReq struct {
	TalentID                         string   `json:"talent_id,omitempty"`                            // 人才ID, 示例值: "12312312312"
	JobID                            string   `json:"job_id,omitempty"`                               // 职位ID, 示例值: "12312312312"
	ResumeSourceID                   *string  `json:"resume_source_id,omitempty"`                     // 简历来源 ID, 可通过「获取简历来源」接口查询。若简历来源类型属于「员工转岗」或「实习生转正」, 人才需处于已入职状态, 示例值: "7115289562569591070"
	ApplicationPreferredCityCodeList []string `json:"application_preferred_city_code_list,omitempty"` // 意向投递城市列表, 可从「获取职位信息」返回的工作地点列表获取, 示例值: ["CT_1"]
}

// CreateHireApplicationResp ...
type CreateHireApplicationResp struct {
	ID string `json:"id,omitempty"` // 投递ID
}

// createHireApplicationResp ...
type createHireApplicationResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateHireApplicationResp `json:"data,omitempty"`
}
