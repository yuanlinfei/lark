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

// SearchCoreHRCurrency 根据货币 ID、状态查询货币信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-currency/search
func (r *CoreHRService) SearchCoreHRCurrency(ctx context.Context, request *SearchCoreHRCurrencyReq, options ...MethodOptionFunc) (*SearchCoreHRCurrencyResp, *Response, error) {
	if r.cli.mock.mockCoreHRSearchCoreHRCurrency != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#SearchCoreHRCurrency mock enable")
		return r.cli.mock.mockCoreHRSearchCoreHRCurrency(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "SearchCoreHRCurrency",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/basic_info/currencies/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchCoreHRCurrencyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRSearchCoreHRCurrency mock CoreHRSearchCoreHRCurrency method
func (r *Mock) MockCoreHRSearchCoreHRCurrency(f func(ctx context.Context, request *SearchCoreHRCurrencyReq, options ...MethodOptionFunc) (*SearchCoreHRCurrencyResp, *Response, error)) {
	r.mockCoreHRSearchCoreHRCurrency = f
}

// UnMockCoreHRSearchCoreHRCurrency un-mock CoreHRSearchCoreHRCurrency method
func (r *Mock) UnMockCoreHRSearchCoreHRCurrency() {
	r.mockCoreHRSearchCoreHRCurrency = nil
}

// SearchCoreHRCurrencyReq ...
type SearchCoreHRCurrencyReq struct {
	PageSize       int64    `query:"page_size" json:"-"`        // 分页大小, 最大 100, 示例值: 100, 取值范围: `1` ～ `100`
	PageToken      *string  `query:"page_token" json:"-"`       // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 6891251722631890445
	CurrencyIDList []string `json:"currency_id_list,omitempty"` // 货币 ID 列表, 示例值: ["6891251722611891445"], 最大长度: `100`
	StatusList     []int64  `json:"status_list,omitempty"`      // 状态列表, 示例值: [1], 可选值有: 1: 生效, 0: 失效, 默认值: `[1]`, 最大长度: `2`
}

// SearchCoreHRCurrencyResp ...
type SearchCoreHRCurrencyResp struct {
	Items     []*SearchCoreHRCurrencyRespItem `json:"items,omitempty"`      // 查询的货币信息
	PageToken string                          `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                            `json:"has_more,omitempty"`   // 是否还有更多项
}

// SearchCoreHRCurrencyRespItem ...
type SearchCoreHRCurrencyRespItem struct {
	CurrencyID         string                                      `json:"currency_id,omitempty"`           // 货币 ID
	CountryRegionID    string                                      `json:"country_region_id,omitempty"`     // 货币所属国家/地区 ID, 详细信息可通过[查询国家 / 地区信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-country_region/search)接口查询获得
	CurrencyName       []*SearchCoreHRCurrencyRespItemCurrencyName `json:"currency_name,omitempty"`         // 货币名称
	NumericCode        int64                                       `json:"numeric_code,omitempty"`          // 数字代码
	CurrencyAlpha3Code string                                      `json:"currency_alpha_3_code,omitempty"` // 三位字母代码
	Status             int64                                       `json:"status,omitempty"`                // 状态, 可选值有: 1: 生效, 0: 失效
}

// SearchCoreHRCurrencyRespItemCurrencyName ...
type SearchCoreHRCurrencyRespItemCurrencyName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// searchCoreHRCurrencyResp ...
type searchCoreHRCurrencyResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *SearchCoreHRCurrencyResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
