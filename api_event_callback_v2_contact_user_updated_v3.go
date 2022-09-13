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

// EventV2ContactUserUpdatedV3 通过该事件订阅员工变更。old_object中只展示更新的字段的原始值。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=user&event=updated)
//
// 只有当应用拥有被改动字段的数据权限时, 才会接收到事件。具体的数据权限与字段的关系请参考[应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN), 或查看事件体参数列表的字段描述。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/events/updated
func (r *EventCallbackService) HandlerEventV2ContactUserUpdatedV3(f EventV2ContactUserUpdatedV3Handler) {
	r.cli.eventHandler.eventV2ContactUserUpdatedV3Handler = f
}

// EventV2ContactUserUpdatedV3Handler event EventV2ContactUserUpdatedV3 handler
type EventV2ContactUserUpdatedV3Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ContactUserUpdatedV3) (string, error)

// EventV2ContactUserUpdatedV3 ...
type EventV2ContactUserUpdatedV3 struct {
	Object    *EventV2ContactUserUpdatedV3Object    `json:"object,omitempty"`     // 变更后信息
	OldObject *EventV2ContactUserUpdatedV3OldObject `json:"old_object,omitempty"` // 变更前信息, body中只包含有变更的字段
}

// EventV2ContactUserUpdatedV3Object ...
type EventV2ContactUserUpdatedV3Object struct {
	OpenID          string                                         `json:"open_id,omitempty"`          // 用户的open_id [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	UnionID         string                                         `json:"union_id,omitempty"`         // 用户的union_id [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	UserID          string                                         `json:"user_id,omitempty"`          // 租户内用户的唯一标识 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 字段权限要求: 获取用户 user ID
	Name            string                                         `json:"name,omitempty"`             // 用户名, 最小长度: `1` 字符, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	EnName          string                                         `json:"en_name,omitempty"`          // 英文名, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Nickname        string                                         `json:"nickname,omitempty"`         // 别名, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Email           string                                         `json:"email,omitempty"`            // 邮箱, 字段权限要求: 获取用户邮箱信息
	EnterpriseEmail string                                         `json:"enterprise_email,omitempty"` // 企业邮箱
	JobTitle        string                                         `json:"job_title,omitempty"`        // 职务
	Mobile          string                                         `json:"mobile,omitempty"`           // 手机号, 字段权限要求: 获取用户手机号
	Gender          int64                                          `json:"gender,omitempty"`           // 性别, 可选值有: 0: 未知, 1: 男, 2: 女, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户性别, 以应用身份访问通讯录, 读取通讯录
	Avatar          *EventV2ContactUserUpdatedV3ObjectAvatar       `json:"avatar,omitempty"`           // 用户头像信息, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Status          *EventV2ContactUserUpdatedV3ObjectStatus       `json:"status,omitempty"`           // 用户状态, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	DepartmentIDs   []string                                       `json:"department_ids,omitempty"`   // 用户所属部门的ID列表, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户组织架构信息, 以应用身份访问通讯录, 读取通讯录
	LeaderUserID    string                                         `json:"leader_user_id,omitempty"`   // 用户的直接主管的用户open_id [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	City            string                                         `json:"city,omitempty"`             // 城市, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	Country         string                                         `json:"country,omitempty"`          // 国家, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	WorkStation     string                                         `json:"work_station,omitempty"`     // 工位, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	JoinTime        int64                                          `json:"join_time,omitempty"`        // 入职时间, 取值范围: `1` ～ `2147483647`, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	EmployeeNo      string                                         `json:"employee_no,omitempty"`      // 工号, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	EmployeeType    int64                                          `json:"employee_type,omitempty"`    // 员工类型, 可选值有: 1: 正式员工, 2: 实习生, 3: 外包, 4: 劳务, 5: 顾问, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	Orders          []*EventV2ContactUserUpdatedV3ObjectOrder      `json:"orders,omitempty"`           // 用户排序信息, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户组织架构信息, 以应用身份访问通讯录, 读取通讯录
	CustomAttrs     []*EventV2ContactUserUpdatedV3ObjectCustomAttr `json:"custom_attrs,omitempty"`     // 自定义属性, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
}

// EventV2ContactUserUpdatedV3ObjectAvatar ...
type EventV2ContactUserUpdatedV3ObjectAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

// EventV2ContactUserUpdatedV3ObjectCustomAttr ...
type EventV2ContactUserUpdatedV3ObjectCustomAttr struct {
	Type  string                                            `json:"type,omitempty"`  // 自定义字段类型, `TEXT`: 文本, `HREF`: 网页, `ENUMERATION`: 枚举, `PICTURE_ENUM`: 图片, `GENERIC_USER`: 用户, 具体说明参见常见问题的[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525)
	ID    string                                            `json:"id,omitempty"`    // 自定义字段ID
	Value *EventV2ContactUserUpdatedV3ObjectCustomAttrValue `json:"value,omitempty"` // 自定义字段取值
}

// EventV2ContactUserUpdatedV3ObjectCustomAttrValue ...
type EventV2ContactUserUpdatedV3ObjectCustomAttrValue struct {
	Text        string                                                       `json:"text,omitempty"`         // 字段类型为`TEXT`时该参数定义字段值, 必填；字段类型为`HREF`时该参数定义网页标题, 必填
	URL         string                                                       `json:"url,omitempty"`          // 字段类型为 HREF 时, 该参数定义默认 URL, 例如手机端跳转小程序, PC端跳转网页
	PcURL       string                                                       `json:"pc_url,omitempty"`       // 字段类型为 HREF 时, 该参数定义PC端 URL
	OptionID    string                                                       `json:"option_id,omitempty"`    // 字段类型为 ENUMERATION 或 PICTURE_ENUM 时, 该参数定义选项值
	OptionValue string                                                       `json:"option_value,omitempty"` // 选项类型的值, 表示 成员详情/自定义字段 中选项选中的值
	Name        string                                                       `json:"name,omitempty"`         // 选项类型为图片时, 表示图片的名称
	PictureURL  string                                                       `json:"picture_url,omitempty"`  // 图片链接
	GenericUser *EventV2ContactUserUpdatedV3ObjectCustomAttrValueGenericUser `json:"generic_user,omitempty"` // 字段类型为 GENERIC_USER 时, 该参数定义引用人员
}

// EventV2ContactUserUpdatedV3ObjectCustomAttrValueGenericUser ...
type EventV2ContactUserUpdatedV3ObjectCustomAttrValueGenericUser struct {
	ID   string `json:"id,omitempty"`   // 用户的user_id, 具体参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Type int64  `json:"type,omitempty"` // 用户类型: 1: 用户, 目前固定为1, 表示用户类型
}

// EventV2ContactUserUpdatedV3ObjectOrder ...
type EventV2ContactUserUpdatedV3ObjectOrder struct {
	DepartmentID    string `json:"department_id,omitempty"`    // 排序信息对应的部门ID, ID值与查询参数中的department_id_type 对应, 表示用户所在的、且需要排序的部门, 不同 ID 的说明参见及获取方式参见 [部门ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview)
	UserOrder       int64  `json:"user_order,omitempty"`       // 用户在其直属部门内的排序, 数值越大, 排序越靠前
	DepartmentOrder int64  `json:"department_order,omitempty"` // 用户所属的多个部门间的排序, 数值越大, 排序越靠前
}

// EventV2ContactUserUpdatedV3ObjectStatus ...
type EventV2ContactUserUpdatedV3ObjectStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否暂停
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
	IsExited    bool `json:"is_exited,omitempty"`    // 是否主动退出, 主动退出一段时间后用户会自动转为已离职
	IsUnjoin    bool `json:"is_unjoin,omitempty"`    // 是否未加入, 需要用户自主确认才能加入团队
}

// EventV2ContactUserUpdatedV3OldObject ...
type EventV2ContactUserUpdatedV3OldObject struct {
	OpenID        string                                            `json:"open_id,omitempty"`        // 用户的open_id [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	UnionID       string                                            `json:"union_id,omitempty"`       // 用户的union_id [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	UserID        string                                            `json:"user_id,omitempty"`        // 租户内用户的唯一标识 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 字段权限要求: 获取用户 user ID
	Name          string                                            `json:"name,omitempty"`           // 用户名, 最小长度: `1` 字符, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	EnName        string                                            `json:"en_name,omitempty"`        // 英文名, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Nickname      string                                            `json:"nickname,omitempty"`       // 别名, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Email         string                                            `json:"email,omitempty"`          // 邮箱, 字段权限要求: 获取用户邮箱信息
	JobTitle      string                                            `json:"job_title,omitempty"`      // 职务
	Mobile        string                                            `json:"mobile,omitempty"`         // 手机号, 字段权限要求: 获取用户手机号
	Gender        int64                                             `json:"gender,omitempty"`         // 性别, 可选值有: 0: 未知, 1: 男, 2: 女, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户性别, 以应用身份访问通讯录, 读取通讯录
	Avatar        *EventV2ContactUserUpdatedV3OldObjectAvatar       `json:"avatar,omitempty"`         // 用户头像信息, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录
	Status        *EventV2ContactUserUpdatedV3OldObjectStatus       `json:"status,omitempty"`         // 用户状态, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	DepartmentIDs []string                                          `json:"department_ids,omitempty"` // 用户所属部门的ID列表, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户组织架构信息, 以应用身份访问通讯录, 读取通讯录
	LeaderUserID  string                                            `json:"leader_user_id,omitempty"` // 用户的直接主管的用户open_id [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	City          string                                            `json:"city,omitempty"`           // 城市, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	Country       string                                            `json:"country,omitempty"`        // 国家, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	WorkStation   string                                            `json:"work_station,omitempty"`   // 工位, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	JoinTime      int64                                             `json:"join_time,omitempty"`      // 入职时间, 取值范围: `1` ～ `2147483647`, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	EmployeeNo    string                                            `json:"employee_no,omitempty"`    // 工号, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	EmployeeType  int64                                             `json:"employee_type,omitempty"`  // 员工类型, 可选值有: 1: 正式员工, 2: 实习生, 3: 外包, 4: 劳务, 5: 顾问, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
	Orders        []*EventV2ContactUserUpdatedV3OldObjectOrder      `json:"orders,omitempty"`         // 用户排序信息, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户组织架构信息, 以应用身份访问通讯录, 读取通讯录
	CustomAttrs   []*EventV2ContactUserUpdatedV3OldObjectCustomAttr `json:"custom_attrs,omitempty"`   // 自定义属性, 字段权限要求（满足任一）: 以应用身份读取通讯录, 获取用户雇佣信息, 以应用身份访问通讯录, 读取通讯录
}

// EventV2ContactUserUpdatedV3OldObjectAvatar ...
type EventV2ContactUserUpdatedV3OldObjectAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

// EventV2ContactUserUpdatedV3OldObjectCustomAttr ...
type EventV2ContactUserUpdatedV3OldObjectCustomAttr struct {
	Type  string                                               `json:"type,omitempty"`  // 自定义字段类型, `TEXT`: 文本, `HREF`: 网页, `ENUMERATION`: 枚举, `PICTURE_ENUM`: 图片, `GENERIC_USER`: 用户, 具体说明参见常见问题的[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525)
	ID    string                                               `json:"id,omitempty"`    // 自定义字段ID
	Value *EventV2ContactUserUpdatedV3OldObjectCustomAttrValue `json:"value,omitempty"` // 自定义字段取值
}

// EventV2ContactUserUpdatedV3OldObjectCustomAttrValue ...
type EventV2ContactUserUpdatedV3OldObjectCustomAttrValue struct {
	Text        string                                                          `json:"text,omitempty"`         // 字段类型为`TEXT`时该参数定义字段值, 必填；字段类型为`HREF`时该参数定义网页标题, 必填
	URL         string                                                          `json:"url,omitempty"`          // 字段类型为 HREF 时, 该参数定义默认 URL, 例如手机端跳转小程序, PC端跳转网页
	PcURL       string                                                          `json:"pc_url,omitempty"`       // 字段类型为 HREF 时, 该参数定义PC端 URL
	OptionID    string                                                          `json:"option_id,omitempty"`    // 字段类型为 ENUMERATION 或 PICTURE_ENUM 时, 该参数定义选项值
	OptionValue string                                                          `json:"option_value,omitempty"` // 选项类型的值, 表示 成员详情/自定义字段 中选项选中的值
	Name        string                                                          `json:"name,omitempty"`         // 选项类型为图片时, 表示图片的名称
	PictureURL  string                                                          `json:"picture_url,omitempty"`  // 图片链接
	GenericUser *EventV2ContactUserUpdatedV3OldObjectCustomAttrValueGenericUser `json:"generic_user,omitempty"` // 字段类型为 GENERIC_USER 时, 该参数定义引用人员
}

// EventV2ContactUserUpdatedV3OldObjectCustomAttrValueGenericUser ...
type EventV2ContactUserUpdatedV3OldObjectCustomAttrValueGenericUser struct {
	ID   string `json:"id,omitempty"`   // 用户的user_id, 具体参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Type int64  `json:"type,omitempty"` // 用户类型: 1: 用户, 目前固定为1, 表示用户类型
}

// EventV2ContactUserUpdatedV3OldObjectOrder ...
type EventV2ContactUserUpdatedV3OldObjectOrder struct {
	DepartmentID    string `json:"department_id,omitempty"`    // 排序信息对应的部门ID, ID值与查询参数中的department_id_type 对应, 表示用户所在的、且需要排序的部门, 不同 ID 的说明参见及获取方式参见 [部门ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview)
	UserOrder       int64  `json:"user_order,omitempty"`       // 用户在其直属部门内的排序, 数值越大, 排序越靠前
	DepartmentOrder int64  `json:"department_order,omitempty"` // 用户所属的多个部门间的排序, 数值越大, 排序越靠前
}

// EventV2ContactUserUpdatedV3OldObjectStatus ...
type EventV2ContactUserUpdatedV3OldObjectStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否暂停
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
	IsExited    bool `json:"is_exited,omitempty"`    // 是否主动退出, 主动退出一段时间后用户会自动转为已离职
	IsUnjoin    bool `json:"is_unjoin,omitempty"`    // 是否未加入, 需要用户自主确认才能加入团队
}
