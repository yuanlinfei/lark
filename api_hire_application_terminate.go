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

// TerminateHireApplication 根据投递 ID 修改投递状态为「已终止」
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/terminate
func (r *HireService) TerminateHireApplication(ctx context.Context, request *TerminateHireApplicationReq, options ...MethodOptionFunc) (*TerminateHireApplicationResp, *Response, error) {
	if r.cli.mock.mockHireTerminateHireApplication != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#TerminateHireApplication mock enable")
		return r.cli.mock.mockHireTerminateHireApplication(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "TerminateHireApplication",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/applications/:application_id/terminate",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(terminateHireApplicationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireTerminateHireApplication mock HireTerminateHireApplication method
func (r *Mock) MockHireTerminateHireApplication(f func(ctx context.Context, request *TerminateHireApplicationReq, options ...MethodOptionFunc) (*TerminateHireApplicationResp, *Response, error)) {
	r.mockHireTerminateHireApplication = f
}

// UnMockHireTerminateHireApplication un-mock HireTerminateHireApplication method
func (r *Mock) UnMockHireTerminateHireApplication() {
	r.mockHireTerminateHireApplication = nil
}

// TerminateHireApplicationReq ...
type TerminateHireApplicationReq struct {
	ApplicationID         string   `path:"application_id" json:"-"`           // 投递ID, 示例值: "12312312312"
	TerminationType       int64    `json:"termination_type,omitempty"`        // 终止原因的类型, 示例值: 1, 可选值有: 1: 我们拒绝了候选人, 22: 候选人拒绝了我们, 27: 其他
	TerminationReasonList []string `json:"termination_reason_list,omitempty"` // 终止的具体原因的id列表, 示例值: ["6891560630172518670"]
	TerminationReasonNote *string  `json:"termination_reason_note,omitempty"` // 终止备注, 示例值: "不符合期望"
}

// TerminateHireApplicationResp ...
type TerminateHireApplicationResp struct {
}

// terminateHireApplicationResp ...
type terminateHireApplicationResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *TerminateHireApplicationResp `json:"data,omitempty"`
}
