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

// GetCoreHRCountryRegionList 批量查询国家/地区信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/country_region/list
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/location_data/list
//
// Deprecated
func (r *CoreHRService) GetCoreHRCountryRegionList(ctx context.Context, request *GetCoreHRCountryRegionListReq, options ...MethodOptionFunc) (*GetCoreHRCountryRegionListResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRCountryRegionList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRCountryRegionList mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRCountryRegionList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRCountryRegionList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/country_regions",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRCountryRegionListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRCountryRegionList mock CoreHRGetCoreHRCountryRegionList method
func (r *Mock) MockCoreHRGetCoreHRCountryRegionList(f func(ctx context.Context, request *GetCoreHRCountryRegionListReq, options ...MethodOptionFunc) (*GetCoreHRCountryRegionListResp, *Response, error)) {
	r.mockCoreHRGetCoreHRCountryRegionList = f
}

// UnMockCoreHRGetCoreHRCountryRegionList un-mock CoreHRGetCoreHRCountryRegionList method
func (r *Mock) UnMockCoreHRGetCoreHRCountryRegionList() {
	r.mockCoreHRGetCoreHRCountryRegionList = nil
}

// GetCoreHRCountryRegionListReq ...
type GetCoreHRCountryRegionListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 1231231987
	PageSize  int64   `query:"page_size" json:"-"`  // 分页大小, 示例值: 100
}

// GetCoreHRCountryRegionListResp ...
type GetCoreHRCountryRegionListResp struct {
	Items     []*GetCoreHRCountryRegionListRespItem `json:"items,omitempty"`      // 国家/地区信息
	HasMore   bool                                  `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetCoreHRCountryRegionListRespItem ...
type GetCoreHRCountryRegionListRespItem struct {
	ID         string                                    `json:"id,omitempty"`           // 国家/地区id
	Name       []*GetCoreHRCountryRegionListRespItemName `json:"name,omitempty"`         // 国家/地区名称
	Alpha3Code string                                    `json:"alpha_3_code,omitempty"` // 国家地区三字码
	Alpha2Code string                                    `json:"alpha_2_code,omitempty"` // 国家地区二字码
}

// GetCoreHRCountryRegionListRespItemName ...
type GetCoreHRCountryRegionListRespItemName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHRCountryRegionListResp ...
type getCoreHRCountryRegionListResp struct {
	Code  int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                          `json:"msg,omitempty"`  // 错误描述
	Data  *GetCoreHRCountryRegionListResp `json:"data,omitempty"`
	Error *ErrorDetail                    `json:"error,omitempty"`
}
