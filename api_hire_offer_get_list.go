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

// GetHireOfferList 根据人才 ID 获取 Offer 列表。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/offer/list
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/candidate-management/delivery-process-management/offer/list
func (r *HireService) GetHireOfferList(ctx context.Context, request *GetHireOfferListReq, options ...MethodOptionFunc) (*GetHireOfferListResp, *Response, error) {
	if r.cli.mock.mockHireGetHireOfferList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#GetHireOfferList mock enable")
		return r.cli.mock.mockHireGetHireOfferList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireOfferList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/offers",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireOfferListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireOfferList mock HireGetHireOfferList method
func (r *Mock) MockHireGetHireOfferList(f func(ctx context.Context, request *GetHireOfferListReq, options ...MethodOptionFunc) (*GetHireOfferListResp, *Response, error)) {
	r.mockHireGetHireOfferList = f
}

// UnMockHireGetHireOfferList un-mock HireGetHireOfferList method
func (r *Mock) UnMockHireGetHireOfferList() {
	r.mockHireGetHireOfferList = nil
}

// GetHireOfferListReq ...
type GetHireOfferListReq struct {
	PageToken          *string `query:"page_token" json:"-"`            // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 1231231987
	PageSize           *int64  `query:"page_size" json:"-"`             // 分页大小, 示例值: 100, 默认值: `1`
	TalentID           string  `query:"talent_id" json:"-"`             // 人才 ID, 示例值: 7096320678581242123
	UserIDType         *IDType `query:"user_id_type" json:"-"`          // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以people_admin_id来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	EmployeeTypeIDType *IDType `query:"employee_type_id_type" json:"-"` // 此次调用中使用的「人员类型 ID」的类型, 示例值: 1, 可选值有: people_admin_employee_type_id: 「人力系统管理后台」适用的人员类型 ID。人力系统管理后台逐步下线中, 建议不继续使用此 ID。, employee_type_enum_id: 「飞书管理后台」适用的人员类型 ID, 通过[「查询人员类型」](https://open.feishu.cn/document/server-docs/contact-v3/employee_type_enum/list)接口获取, 默认值: `people_admin_employee_type_id`
}

// GetHireOfferListResp ...
type GetHireOfferListResp struct {
	HasMore   bool                        `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                      `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetHireOfferListRespItem `json:"items,omitempty"`      // Offer 列表
}

// GetHireOfferListRespItem ...
type GetHireOfferListRespItem struct {
	ID            string                                `json:"id,omitempty"`             // Offer ID
	JobInfo       *GetHireOfferListRespItemJobInfo      `json:"job_info,omitempty"`       // Offer 职位
	CreateTime    string                                `json:"create_time,omitempty"`    // 创建时间
	OfferStatus   int64                                 `json:"offer_status,omitempty"`   // Offer 状态, 可选值有: 1: 未申请, 2: 审批中, 3: 审批已撤回, 4: 审批通过, 5: 审批不通过, 6: Offer 已发出, 7: 候选人已接受, 8: 候选人已拒绝, 9: Offer 已失效, 10: 未审批, 11: 实习待入职（仅实习 Offer 具有）, 12: 实习已入职（仅实习 Offer 具有）, 13: 实习已离职（仅实习 Offer 具有）
	OfferType     int64                                 `json:"offer_type,omitempty"`     // Offer 类型, 可选值有: 1: 正式 Offer, 2: 实习 Offer
	EmployeeType  *GetHireOfferListRespItemEmployeeType `json:"employee_type,omitempty"`  // Offer 人员类型
	ApplicationID string                                `json:"application_id,omitempty"` // Offer 投递 ID
}

// GetHireOfferListRespItemEmployeeType ...
type GetHireOfferListRespItemEmployeeType struct {
	ID     string `json:"id,omitempty"`      // ID
	ZhName string `json:"zh_name,omitempty"` // 中文名称
	EnName string `json:"en_name,omitempty"` // 英文名称
}

// GetHireOfferListRespItemJobInfo ...
type GetHireOfferListRespItemJobInfo struct {
	JobID   string `json:"job_id,omitempty"`   // Offer 职位 ID
	JobName string `json:"job_name,omitempty"` // Offer 职位名称
}

// getHireOfferListResp ...
type getHireOfferListResp struct {
	Code  int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                `json:"msg,omitempty"`  // 错误描述
	Data  *GetHireOfferListResp `json:"data,omitempty"`
	Error *ErrorDetail          `json:"error,omitempty"`
}
