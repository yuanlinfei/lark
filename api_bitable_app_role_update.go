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

// UpdateBitableAppRole 更新自定义权限
//
// 更新自定义权限是全量更新, 会完全覆盖旧的自定义权限设置
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role/update
func (r *BitableService) UpdateBitableAppRole(ctx context.Context, request *UpdateBitableAppRoleReq, options ...MethodOptionFunc) (*UpdateBitableAppRoleResp, *Response, error) {
	if r.cli.mock.mockBitableUpdateBitableAppRole != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#UpdateBitableAppRole mock enable")
		return r.cli.mock.mockBitableUpdateBitableAppRole(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "UpdateBitableAppRole",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/roles/:role_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateBitableAppRoleResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableUpdateBitableAppRole mock BitableUpdateBitableAppRole method
func (r *Mock) MockBitableUpdateBitableAppRole(f func(ctx context.Context, request *UpdateBitableAppRoleReq, options ...MethodOptionFunc) (*UpdateBitableAppRoleResp, *Response, error)) {
	r.mockBitableUpdateBitableAppRole = f
}

// UnMockBitableUpdateBitableAppRole un-mock BitableUpdateBitableAppRole method
func (r *Mock) UnMockBitableUpdateBitableAppRole() {
	r.mockBitableUpdateBitableAppRole = nil
}

// UpdateBitableAppRoleReq ...
type UpdateBitableAppRoleReq struct {
	AppToken   string                              `path:"app_token" json:"-"`    // bitable app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	RoleID     string                              `path:"role_id" json:"-"`      // 自定义权限的id, 示例值: "roljRpwIUt"
	RoleName   string                              `json:"role_name,omitempty"`   // 自定义权限的名字, 示例值: "自定义权限1"
	TableRoles []*UpdateBitableAppRoleReqTableRole `json:"table_roles,omitempty"` // 数据表权限, 最大长度: `100`
}

// UpdateBitableAppRoleReqTableRole ...
type UpdateBitableAppRoleReqTableRole struct {
	TableName string                                   `json:"table_name,omitempty"` // 数据表名, 示例值: "数据表1"
	TablePerm int64                                    `json:"table_perm,omitempty"` // 数据表权限, `协作者可编辑自己的记录`和`可编辑指定字段`是`可编辑记录`的特殊情况, 可通过指定`rec_rule`或`field_perm`参数实现相同的效果, 示例值: 0, 可选值有: 0: 无权限, 1: 可阅读, 2: 可编辑记录, 4: 可编辑字段和记录, 默认值: `0`
	RecRule   *UpdateBitableAppRoleReqTableRoleRecRule `json:"rec_rule,omitempty"`   // 记录筛选条件, 在table_perm为1或2时有意义, 用于指定可编辑或可阅读某些记录
	FieldPerm map[string]int64                         `json:"field_perm,omitempty"` // 字段权限, 仅在table_perm为2时有意义, 设置字段可编辑或可阅读。类型为 map, key 是字段名, value 是字段权限, value 枚举值有: `1`: 可阅读, `2`: 可编辑
}

// UpdateBitableAppRoleReqTableRoleRecRule ...
type UpdateBitableAppRoleReqTableRoleRecRule struct {
	Conditions  []*UpdateBitableAppRoleReqTableRoleRecRuleCondition `json:"conditions,omitempty"`  // 记录筛选条件, 最大长度: `100`
	Conjunction *string                                             `json:"conjunction,omitempty"` // 多个筛选条件的关系, 示例值: "and", 可选值有: and: 与, or: 或, 默认值: `and`
	OtherPerm   *int64                                              `json:"other_perm,omitempty"`  // 其他记录权限, 仅在table_perm为2时有意义, 示例值: 0, 可选值有: 0: 禁止查看, 1: 仅可阅读, 默认值: `0`
}

// UpdateBitableAppRoleReqTableRoleRecRuleCondition ...
type UpdateBitableAppRoleReqTableRoleRecRuleCondition struct {
	FieldName string   `json:"field_name,omitempty"` // 字段名, 记录筛选条件是`创建人包含访问者本人`时, 此参数值为"", 示例值: "单选"
	Operator  *string  `json:"operator,omitempty"`   // 运算符, 示例值: "is", 可选值有: is: 等于, isNot: 不等于, contains: 包含, doesNotContain: 不包含, isEmpty: 为空, isNotEmpty: 不为空, 默认值: `is`
	Value     []string `json:"value,omitempty"`      // 单选或多选字段的选项id, 示例值: ["optbdVHf4q", "optrpd3eIJ"]
}

// UpdateBitableAppRoleResp ...
type UpdateBitableAppRoleResp struct {
	Role *UpdateBitableAppRoleRespRole `json:"role,omitempty"` // 自定义权限
}

// UpdateBitableAppRoleRespRole ...
type UpdateBitableAppRoleRespRole struct {
	RoleName   string                                   `json:"role_name,omitempty"`   // 自定义权限的名字
	RoleID     string                                   `json:"role_id,omitempty"`     // 自定义权限的id
	TableRoles []*UpdateBitableAppRoleRespRoleTableRole `json:"table_roles,omitempty"` // 数据表权限
}

// UpdateBitableAppRoleRespRoleTableRole ...
type UpdateBitableAppRoleRespRoleTableRole struct {
	TableName string                                        `json:"table_name,omitempty"` // 数据表名
	TablePerm int64                                         `json:"table_perm,omitempty"` // 数据表权限, `协作者可编辑自己的记录`和`可编辑指定字段`是`可编辑记录`的特殊情况, 可通过指定`rec_rule`或`field_perm`参数实现相同的效果, 可选值有: 0: 无权限, 1: 可阅读, 2: 可编辑记录, 4: 可编辑字段和记录
	RecRule   *UpdateBitableAppRoleRespRoleTableRoleRecRule `json:"rec_rule,omitempty"`   // 记录筛选条件, 在table_perm为1或2时有意义, 用于指定可编辑或可阅读某些记录
	FieldPerm map[string]int64                              `json:"field_perm,omitempty"` // 字段权限, 仅在table_perm为2时有意义, 设置字段可编辑或可阅读。类型为 map, key 是字段名, value 是字段权限, value 枚举值有: `1`: 可阅读, `2`: 可编辑
}

// UpdateBitableAppRoleRespRoleTableRoleRecRule ...
type UpdateBitableAppRoleRespRoleTableRoleRecRule struct {
	Conditions  []*UpdateBitableAppRoleRespRoleTableRoleRecRuleCondition `json:"conditions,omitempty"`  // 记录筛选条件
	Conjunction string                                                   `json:"conjunction,omitempty"` // 多个筛选条件的关系, 可选值有: and: 与, or: 或
	OtherPerm   int64                                                    `json:"other_perm,omitempty"`  // 其他记录权限, 仅在table_perm为2时有意义, 可选值有: 0: 禁止查看, 1: 仅可阅读
}

// UpdateBitableAppRoleRespRoleTableRoleRecRuleCondition ...
type UpdateBitableAppRoleRespRoleTableRoleRecRuleCondition struct {
	FieldName string   `json:"field_name,omitempty"` // 字段名, 记录筛选条件是`创建人包含访问者本人`时, 此参数值为""
	Operator  string   `json:"operator,omitempty"`   // 运算符, 可选值有: is: 等于, isNot: 不等于, contains: 包含, doesNotContain: 不包含, isEmpty: 为空, isNotEmpty: 不为空
	Value     []string `json:"value,omitempty"`      // 单选或多选字段的选项id
	FieldType int64    `json:"field_type,omitempty"` // 字段类型
}

// updateBitableAppRoleResp ...
type updateBitableAppRoleResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *UpdateBitableAppRoleResp `json:"data,omitempty"`
}
