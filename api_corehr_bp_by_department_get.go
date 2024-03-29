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

// GetCoreHrbpByDepartment 查询部门的 HRBP 信息, 包括来自上级部门的 HRBP。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/bp/get_by_department
func (r *CoreHRService) GetCoreHrbpByDepartment(ctx context.Context, request *GetCoreHrbpByDepartmentReq, options ...MethodOptionFunc) (*GetCoreHrbpByDepartmentResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHrbpByDepartment != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHrbpByDepartment mock enable")
		return r.cli.mock.mockCoreHRGetCoreHrbpByDepartment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHrbpByDepartment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/bps/get_by_department",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHrbpByDepartmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHrbpByDepartment mock CoreHRGetCoreHrbpByDepartment method
func (r *Mock) MockCoreHRGetCoreHrbpByDepartment(f func(ctx context.Context, request *GetCoreHrbpByDepartmentReq, options ...MethodOptionFunc) (*GetCoreHrbpByDepartmentResp, *Response, error)) {
	r.mockCoreHRGetCoreHrbpByDepartment = f
}

// UnMockCoreHRGetCoreHrbpByDepartment un-mock CoreHRGetCoreHrbpByDepartment method
func (r *Mock) UnMockCoreHRGetCoreHrbpByDepartment() {
	r.mockCoreHRGetCoreHrbpByDepartment = nil
}

// GetCoreHrbpByDepartmentReq ...
type GetCoreHrbpByDepartmentReq struct {
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门 ID 类型, 示例值: open_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `open_department_id`
	DepartmentID     string            `json:"department_id,omitempty"`      // 部门 ID, 示例值: "6893014062142064111"
}

// GetCoreHrbpByDepartmentResp ...
type GetCoreHrbpByDepartmentResp struct {
	Items []*GetCoreHrbpByDepartmentRespItem `json:"items,omitempty"` // 部门 HRBP 信息, 依次为部门及各层级上级部门
}

// GetCoreHrbpByDepartmentRespItem ...
type GetCoreHrbpByDepartmentRespItem struct {
	DepartmentID string   `json:"department_id,omitempty"` // 部门 ID
	HrbpIDs      []string `json:"hrbp_ids,omitempty"`      // 部门 HRBP 雇佣 ID
}

// getCoreHrbpByDepartmentResp ...
type getCoreHrbpByDepartmentResp struct {
	Code  int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                       `json:"msg,omitempty"`  // 错误描述
	Data  *GetCoreHrbpByDepartmentResp `json:"data,omitempty"`
	Error *ErrorDetail                 `json:"error,omitempty"`
}
