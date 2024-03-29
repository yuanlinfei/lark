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

// SearchCoreHRCostCenter 查询成本中心信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/cost_center/search
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/cost_center/search
func (r *CoreHRService) SearchCoreHRCostCenter(ctx context.Context, request *SearchCoreHRCostCenterReq, options ...MethodOptionFunc) (*SearchCoreHRCostCenterResp, *Response, error) {
	if r.cli.mock.mockCoreHRSearchCoreHRCostCenter != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#SearchCoreHRCostCenter mock enable")
		return r.cli.mock.mockCoreHRSearchCoreHRCostCenter(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "SearchCoreHRCostCenter",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/cost_centers/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchCoreHRCostCenterResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRSearchCoreHRCostCenter mock CoreHRSearchCoreHRCostCenter method
func (r *Mock) MockCoreHRSearchCoreHRCostCenter(f func(ctx context.Context, request *SearchCoreHRCostCenterReq, options ...MethodOptionFunc) (*SearchCoreHRCostCenterResp, *Response, error)) {
	r.mockCoreHRSearchCoreHRCostCenter = f
}

// UnMockCoreHRSearchCoreHRCostCenter un-mock CoreHRSearchCoreHRCostCenter method
func (r *Mock) UnMockCoreHRSearchCoreHRCostCenter() {
	r.mockCoreHRSearchCoreHRCostCenter = nil
}

// SearchCoreHRCostCenterReq ...
type SearchCoreHRCostCenterReq struct {
	PageSize           int64    `query:"page_size" json:"-"`             // 分页大小, 最大 100, 示例值: 100, 取值范围: `1` ～ `100`
	PageToken          *string  `query:"page_token" json:"-"`            // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 6891251722631890445
	UserIDType         *IDType  `query:"user_id_type" json:"-"`          // 用户 ID 类型, 示例值: people_corehr_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `people_corehr_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	CostCenterIDList   []string `json:"cost_center_id_list,omitempty"`   // 成本中心ID 列表, 示例值: ["7140964208476371111"]
	NameList           []string `json:"name_list,omitempty"`             // 成长中心名称列表, 精确匹配, 示例值: ["技术部成本中心"]
	Code               *string  `json:"code,omitempty"`                  // 成本中心编码, 示例值: "MDPD00000023"
	ParentCostCenterID *string  `json:"parent_cost_center_id,omitempty"` // 上级成本中心ID, 可用于查询直接下级成本中心, 示例值: "6862995757234914824"
	GetAllVersion      *bool    `json:"get_all_version,omitempty"`       // 是否获取所有成本中心版本, true 为获取成本中心所有版本记录, false 为仅获取当前生效的成本中心记录, 默认为 false, 当填写 true 并输入其他查询条件时, 返回的是所有符合查询条件的版本信息, 示例值: true, 默认值: `false`
}

// SearchCoreHRCostCenterResp ...
type SearchCoreHRCostCenterResp struct {
	Items     []*SearchCoreHRCostCenterRespItem `json:"items,omitempty"`      // 成本中心信息
	PageToken string                            `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                              `json:"has_more,omitempty"`   // 是否还有更多项
}

// SearchCoreHRCostCenterRespItem ...
type SearchCoreHRCostCenterRespItem struct {
	CostCenterID       string                                       `json:"cost_center_id,omitempty"`        // 成本中心ID
	VersionID          string                                       `json:"version_id,omitempty"`            // 成本中心版本ID
	Name               []*SearchCoreHRCostCenterRespItemName        `json:"name,omitempty"`                  // 成本中心名称
	Code               string                                       `json:"code,omitempty"`                  // 编码
	ParentCostCenterID string                                       `json:"parent_cost_center_id,omitempty"` // 上级成本中心ID
	Managers           []string                                     `json:"managers,omitempty"`              // 成本中心负责人ID 列表
	Description        []*SearchCoreHRCostCenterRespItemDescription `json:"description,omitempty"`           // 成本中心描述
	EffectiveTime      string                                       `json:"effective_time,omitempty"`        // 生效时间
	ExpirationTime     string                                       `json:"expiration_time,omitempty"`       // 过期时间
	Active             bool                                         `json:"active,omitempty"`                // 当前实体是否启用
}

// SearchCoreHRCostCenterRespItemDescription ...
type SearchCoreHRCostCenterRespItemDescription struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// SearchCoreHRCostCenterRespItemName ...
type SearchCoreHRCostCenterRespItemName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// searchCoreHRCostCenterResp ...
type searchCoreHRCostCenterResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *SearchCoreHRCostCenterResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}
