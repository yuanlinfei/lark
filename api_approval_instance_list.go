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

// GetApprovalInstanceList
//
// 根据 approval_code 批量获取审批实例的 instance_code, 用于拉取租户下某个审批定义的全部审批实例。
// 默认以审批创建时间排序。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQDOyUjL0gjM14CN4ITN
func (r *ApprovalService) GetApprovalInstanceList(ctx context.Context, request *GetApprovalInstanceListReq, options ...MethodOptionFunc) (*GetApprovalInstanceListResp, *Response, error) {
	if r.cli.mock.mockApprovalGetApprovalInstanceList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#GetApprovalInstanceList mock enable")
		return r.cli.mock.mockApprovalGetApprovalInstanceList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "GetApprovalInstanceList",
		Method:                "POST",
		URL:                   r.cli.wwwBaseURL + "/approval/openapi/v2/instance/list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApprovalInstanceListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalGetApprovalInstanceList mock ApprovalGetApprovalInstanceList method
func (r *Mock) MockApprovalGetApprovalInstanceList(f func(ctx context.Context, request *GetApprovalInstanceListReq, options ...MethodOptionFunc) (*GetApprovalInstanceListResp, *Response, error)) {
	r.mockApprovalGetApprovalInstanceList = f
}

// UnMockApprovalGetApprovalInstanceList un-mock ApprovalGetApprovalInstanceList method
func (r *Mock) UnMockApprovalGetApprovalInstanceList() {
	r.mockApprovalGetApprovalInstanceList = nil
}

// GetApprovalInstanceListReq ...
type GetApprovalInstanceListReq struct {
	ApprovalCode string `json:"approval_code,omitempty"` // 审批定义唯一标识
	StartTime    int64  `json:"start_time,omitempty"`    // 审批实例创建时间区间（毫秒）
	EndTime      int64  `json:"end_time,omitempty"`      // 审批实例创建时间区间（毫秒）
	Offset       int64  `json:"offset,omitempty"`        // 查询偏移量
	Limit        int64  `json:"limit,omitempty"`         // 查询限制量 注:不得大于100
}

// GetApprovalInstanceListResp ...
type GetApprovalInstanceListResp struct {
	InstanceCodeList []string `json:"instance_code_list,omitempty"` // 审批实例 Code
}

// getApprovalInstanceListResp ...
type getApprovalInstanceListResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非0表示失败
	Msg  string                       `json:"msg,omitempty"`  // 返回码的描述
	Data *GetApprovalInstanceListResp `json:"data,omitempty"` // 返回业务信息
}