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

// UpdateContactGroup 使用该接口更新用户组信息, 请注意更新用户组时应用的通讯录权限范围需为“全部员工”, 否则会更新失败。[点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/patch
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/group/patch
func (r *ContactService) UpdateContactGroup(ctx context.Context, request *UpdateContactGroupReq, options ...MethodOptionFunc) (*UpdateContactGroupResp, *Response, error) {
	if r.cli.mock.mockContactUpdateContactGroup != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#UpdateContactGroup mock enable")
		return r.cli.mock.mockContactUpdateContactGroup(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "UpdateContactGroup",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/group/:group_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateContactGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactUpdateContactGroup mock ContactUpdateContactGroup method
func (r *Mock) MockContactUpdateContactGroup(f func(ctx context.Context, request *UpdateContactGroupReq, options ...MethodOptionFunc) (*UpdateContactGroupResp, *Response, error)) {
	r.mockContactUpdateContactGroup = f
}

// UnMockContactUpdateContactGroup un-mock ContactUpdateContactGroup method
func (r *Mock) UnMockContactUpdateContactGroup() {
	r.mockContactUpdateContactGroup = nil
}

// UpdateContactGroupReq ...
type UpdateContactGroupReq struct {
	GroupID          string            `path:"group_id" json:"-"`            // 用户组ID, 示例值: "g187131"
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值: open_department_id, 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门
	Name             *string           `json:"name,omitempty"`               // 用户组的名字, 企业内唯一, 最大长度: 100 字符, 示例值: "外包 IT 用户组"
	Description      *string           `json:"description,omitempty"`        // 用户组描述信息, 最大长度: 500 字, 示例值: "IT 外包用户组, 需要进行细粒度权限管控"
}

// UpdateContactGroupResp ...
type UpdateContactGroupResp struct {
}

// updateContactGroupResp ...
type updateContactGroupResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *UpdateContactGroupResp `json:"data,omitempty"`
}
