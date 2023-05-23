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

// GetParentDepartment 该接口用来递归获取部门父部门的信息, 并按照由子到父的顺序返回有权限的父部门信息列表。
//
// 使用`tenant_access_token`时, 该接口只返回可见性范围内的父部门信息。
// 例如: A >>B>>C>>D四级部门, 通讯录权限只到B, 那么查询D部门的parent, 会返回B和C两级部门。
// 使用user_access_token时, 该接口只返回对于用户可见的父部门信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/parent
func (r *ContactService) GetParentDepartment(ctx context.Context, request *GetParentDepartmentReq, options ...MethodOptionFunc) (*GetParentDepartmentResp, *Response, error) {
	if r.cli.mock.mockContactGetParentDepartment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetParentDepartment mock enable")
		return r.cli.mock.mockContactGetParentDepartment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetParentDepartment",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/departments/parent",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getParentDepartmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetParentDepartment mock ContactGetParentDepartment method
func (r *Mock) MockContactGetParentDepartment(f func(ctx context.Context, request *GetParentDepartmentReq, options ...MethodOptionFunc) (*GetParentDepartmentResp, *Response, error)) {
	r.mockContactGetParentDepartment = f
}

// UnMockContactGetParentDepartment un-mock ContactGetParentDepartment method
func (r *Mock) UnMockContactGetParentDepartment() {
	r.mockContactGetParentDepartment = nil
}

// GetParentDepartmentReq ...
type GetParentDepartmentReq struct {
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值: "open_department_id", 可选值有: department_id: 用来标识租户内一个唯一的部门, open_department_id: 用来在具体某个应用中标识一个部门, 同一个部门 在不同应用中的 open_department_id 不相同。
	DepartmentID     string            `query:"department_id" json:"-"`      // 部门ID, 示例值: "od-4e6ac4d14bcd5071a37a39de902c7141"
	PageToken        *string           `query:"page_token" json:"-"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "AQD9/Rn9eij9Pm39ED40/RD/cIFmu77WxpxPB/2oHfQLZ%2BG8JG6tK7%2BZnHiT7COhD2hMSICh/eBl7cpzU6JEC3J7COKNe4jrQ8ExwBCR"
	PageSize         *int64            `query:"page_size" json:"-"`          // 分页大小, 示例值: 10, 默认值: `20`, 最大值: `50`
}

// GetParentDepartmentResp ...
type GetParentDepartmentResp struct {
	HasMore   bool                           `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                         `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetParentDepartmentRespItem `json:"items,omitempty"`
}

// GetParentDepartmentRespItem ...
type GetParentDepartmentRespItem struct {
	Name                   string                               `json:"name,omitempty"`                      // 部门名称, 注意: 不可包含斜杠, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	I18nName               *GetParentDepartmentRespItemI18nName `json:"i18n_name,omitempty"`                 // 国际化的部门名称, 注意: 不可包含斜杠, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	ParentDepartmentID     string                               `json:"parent_department_id,omitempty"`      // 父部门的ID, * 在根部门下创建新部门, 该参数值为 “0”, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	DepartmentID           string                               `json:"department_id,omitempty"`             // 本部门的自定义部门ID, 注意: 除需要满足正则规则外, 同时不能以`od-`开头, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	OpenDepartmentID       string                               `json:"open_department_id,omitempty"`        // 部门的open_id, 类型与通过请求的查询参数传入的department_id_type相同
	LeaderUserID           string                               `json:"leader_user_id,omitempty"`            // 部门主管用户ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	ChatID                 string                               `json:"chat_id,omitempty"`                   // 部门群ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	Order                  string                               `json:"order,omitempty"`                     // 部门的排序, 即部门在其同级部门的展示顺序, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	UnitIDs                []string                             `json:"unit_ids,omitempty"`                  // 部门单位自定义ID列表, 当前只支持一个, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	MemberCount            int64                                `json:"member_count,omitempty"`              // 部门下用户的个数, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
	Status                 *GetParentDepartmentRespItemStatus   `json:"status,omitempty"`                    // 部门状态, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取部门基础信息, 以应用身份访问通讯录, 读取通讯录
	Leaders                []*GetParentDepartmentRespItemLeader `json:"leaders,omitempty"`                   // 部门负责人
	GroupChatEmployeeTypes []int64                              `json:"group_chat_employee_types,omitempty"` // 部门群雇员类型限制。[]空列表时, 表示为无任何雇员类型。类型字段可包含以下值, 支持多个类型值；若有多个, 用英文', '分隔: 1、正式员工, 2、实习生, 3、外包, 4、劳务, 5、顾问, 6、其他自定义类型字段, 可通过下方接口获取到该租户的自定义员工类型的名称, 参见[获取人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list)。
	DepartmentHrbps        []string                             `json:"department_hrbps,omitempty"`          // 部门HRBP, 字段权限要求: 查询部门 HRBP 信息
	PrimaryMemberCount     int64                                `json:"primary_member_count,omitempty"`      // 当前部门主属成员（即成员的主部门为当前部门）的数量, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
}

// GetParentDepartmentRespItemI18nName ...
type GetParentDepartmentRespItemI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 部门的中文名
	JaJp string `json:"ja_jp,omitempty"` // 部门的日文名
	EnUs string `json:"en_us,omitempty"` // 部门的英文名
}

// GetParentDepartmentRespItemLeader ...
type GetParentDepartmentRespItemLeader struct {
	LeaderType int64  `json:"leaderType,omitempty"` // 负责人类型, 可选值有: 1: 主负责人, 2: 副负责人
	LeaderID   string `json:"leaderID,omitempty"`   // 负责人ID, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取通讯录部门组织架构信息, 以应用身份访问通讯录, 读取通讯录
}

// GetParentDepartmentRespItemStatus ...
type GetParentDepartmentRespItemStatus struct {
	IsDeleted bool `json:"is_deleted,omitempty"` // 是否被删除
}

// getParentDepartmentResp ...
type getParentDepartmentResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetParentDepartmentResp `json:"data,omitempty"`
}
