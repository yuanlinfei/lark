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

// SearchApplicationWorkplaceAccessData 获取工作台访问数据（包含默认工作台与定制工作台）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/workplace-v1/workplace_access_data/search
func (r *ApplicationService) SearchApplicationWorkplaceAccessData(ctx context.Context, request *SearchApplicationWorkplaceAccessDataReq, options ...MethodOptionFunc) (*SearchApplicationWorkplaceAccessDataResp, *Response, error) {
	if r.cli.mock.mockApplicationSearchApplicationWorkplaceAccessData != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Application#SearchApplicationWorkplaceAccessData mock enable")
		return r.cli.mock.mockApplicationSearchApplicationWorkplaceAccessData(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "SearchApplicationWorkplaceAccessData",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/workplace/v1/workplace_access_data/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchApplicationWorkplaceAccessDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationSearchApplicationWorkplaceAccessData mock ApplicationSearchApplicationWorkplaceAccessData method
func (r *Mock) MockApplicationSearchApplicationWorkplaceAccessData(f func(ctx context.Context, request *SearchApplicationWorkplaceAccessDataReq, options ...MethodOptionFunc) (*SearchApplicationWorkplaceAccessDataResp, *Response, error)) {
	r.mockApplicationSearchApplicationWorkplaceAccessData = f
}

// UnMockApplicationSearchApplicationWorkplaceAccessData un-mock ApplicationSearchApplicationWorkplaceAccessData method
func (r *Mock) UnMockApplicationSearchApplicationWorkplaceAccessData() {
	r.mockApplicationSearchApplicationWorkplaceAccessData = nil
}

// SearchApplicationWorkplaceAccessDataReq ...
type SearchApplicationWorkplaceAccessDataReq struct {
	FromDate  string  `query:"from_date" json:"-"`  // 数据检索开始时间, 精确到日。格式yyyy-MM-dd, 示例值: 2023-03-01
	ToDate    string  `query:"to_date" json:"-"`    // 数据检索结束时间, 精确到日。格式yyyy-MM-dd, 示例值: 2023-03-22
	PageSize  int64   `query:"page_size" json:"-"`  // 分页大小, 最小为 1, 最大为 200, 默认为 20, 示例值: 20, 默认值: `20`, 取值范围: `1` ～ `200`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: ddowkdkl9w2d
}

// SearchApplicationWorkplaceAccessDataResp ...
type SearchApplicationWorkplaceAccessDataResp struct {
	Items     []*SearchApplicationWorkplaceAccessDataRespItem `json:"items,omitempty"`      // 工作台访问数据
	HasMore   bool                                            `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                          `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// SearchApplicationWorkplaceAccessDataRespItem ...
type SearchApplicationWorkplaceAccessDataRespItem struct {
	Date             string                                                        `json:"date,omitempty"`              // 时间, 精确到天, 格式yyyy-MM-dd
	AllWorkplace     *SearchApplicationWorkplaceAccessDataRespItemAllWorkplace     `json:"all_workplace,omitempty"`     // 全部工作台的访问数据。包含默认工作台和定制工作台。由于历史原因, 部分情况下这这两个数据的和加起来不等于全部工作台的访问数据。如有疑问, 可联系飞书技术支持。
	DefaultWorkplace *SearchApplicationWorkplaceAccessDataRespItemDefaultWorkplace `json:"default_workplace,omitempty"` // 默认工作台的访问数据
}

// SearchApplicationWorkplaceAccessDataRespItemAllWorkplace ...
type SearchApplicationWorkplaceAccessDataRespItemAllWorkplace struct {
	Pv int64 `json:"pv,omitempty"` // 访问次数
	Uv int64 `json:"uv,omitempty"` // 访问用户数(去重)
}

// SearchApplicationWorkplaceAccessDataRespItemDefaultWorkplace ...
type SearchApplicationWorkplaceAccessDataRespItemDefaultWorkplace struct {
	Pv int64 `json:"pv,omitempty"` // 访问次数
	Uv int64 `json:"uv,omitempty"` // 访问用户数(去重)
}

// searchApplicationWorkplaceAccessDataResp ...
type searchApplicationWorkplaceAccessDataResp struct {
	Code  int64                                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                                    `json:"msg,omitempty"`  // 错误描述
	Data  *SearchApplicationWorkplaceAccessDataResp `json:"data,omitempty"`
	Error *ErrorDetail                              `json:"error,omitempty"`
}
