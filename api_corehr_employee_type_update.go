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

// UpdateCoreHREmployeeType 更新人员类型。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/employee_type/patch
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/employee_type/patch
func (r *CoreHRService) UpdateCoreHREmployeeType(ctx context.Context, request *UpdateCoreHREmployeeTypeReq, options ...MethodOptionFunc) (*UpdateCoreHREmployeeTypeResp, *Response, error) {
	if r.cli.mock.mockCoreHRUpdateCoreHREmployeeType != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#UpdateCoreHREmployeeType mock enable")
		return r.cli.mock.mockCoreHRUpdateCoreHREmployeeType(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "UpdateCoreHREmployeeType",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/employee_types/:employee_type_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateCoreHREmployeeTypeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRUpdateCoreHREmployeeType mock CoreHRUpdateCoreHREmployeeType method
func (r *Mock) MockCoreHRUpdateCoreHREmployeeType(f func(ctx context.Context, request *UpdateCoreHREmployeeTypeReq, options ...MethodOptionFunc) (*UpdateCoreHREmployeeTypeResp, *Response, error)) {
	r.mockCoreHRUpdateCoreHREmployeeType = f
}

// UnMockCoreHRUpdateCoreHREmployeeType un-mock CoreHRUpdateCoreHREmployeeType method
func (r *Mock) UnMockCoreHRUpdateCoreHREmployeeType() {
	r.mockCoreHRUpdateCoreHREmployeeType = nil
}

// UpdateCoreHREmployeeTypeReq ...
type UpdateCoreHREmployeeTypeReq struct {
	EmployeeTypeID      string                                    `path:"employee_type_id" json:"-"`       // 雇员类型 ID, 示例值: "6969828847931885087"
	ClientToken         *string                                   `query:"client_token" json:"-"`          // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	Name                []*UpdateCoreHREmployeeTypeReqName        `json:"name,omitempty"`                  // 名称
	DefaultEmployeeType *bool                                     `json:"default_employee_type,omitempty"` // 是否为默认人员类型, 每个租户只能定义一个默认人员类型, 示例值: true
	Active              *bool                                     `json:"active,omitempty"`                // 启用, 示例值: true
	Code                *string                                   `json:"code,omitempty"`                  // 编码, 示例值: "1245"
	CustomFields        []*UpdateCoreHREmployeeTypeReqCustomField `json:"custom_fields,omitempty"`         // 自定义字段
}

// UpdateCoreHREmployeeTypeReqCustomField ...
type UpdateCoreHREmployeeTypeReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// UpdateCoreHREmployeeTypeReqName ...
type UpdateCoreHREmployeeTypeReqName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// UpdateCoreHREmployeeTypeResp ...
type UpdateCoreHREmployeeTypeResp struct {
	EmployeeType *UpdateCoreHREmployeeTypeRespEmployeeType `json:"employee_type,omitempty"` // 人员类型
}

// UpdateCoreHREmployeeTypeRespEmployeeType ...
type UpdateCoreHREmployeeTypeRespEmployeeType struct {
	ID                  string                                                 `json:"id,omitempty"`                    // 雇员类型ID
	Name                []*UpdateCoreHREmployeeTypeRespEmployeeTypeName        `json:"name,omitempty"`                  // 名称
	DefaultEmployeeType bool                                                   `json:"default_employee_type,omitempty"` // 是否为默认人员类型, 每个租户只能定义一个默认人员类型
	Active              bool                                                   `json:"active,omitempty"`                // 启用
	Code                string                                                 `json:"code,omitempty"`                  // 编码
	CustomFields        []*UpdateCoreHREmployeeTypeRespEmployeeTypeCustomField `json:"custom_fields,omitempty"`         // 自定义字段
}

// UpdateCoreHREmployeeTypeRespEmployeeTypeCustomField ...
type UpdateCoreHREmployeeTypeRespEmployeeTypeCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// UpdateCoreHREmployeeTypeRespEmployeeTypeName ...
type UpdateCoreHREmployeeTypeRespEmployeeTypeName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// updateCoreHREmployeeTypeResp ...
type updateCoreHREmployeeTypeResp struct {
	Code  int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                        `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateCoreHREmployeeTypeResp `json:"data,omitempty"`
	Error *ErrorDetail                  `json:"error,omitempty"`
}
