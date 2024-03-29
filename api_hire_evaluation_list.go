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

// GetHireEvaluationList 获取简历评估信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/evaluation/list
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/candidate-management/delivery-process-management/evaluation/list
func (r *HireService) GetHireEvaluationList(ctx context.Context, request *GetHireEvaluationListReq, options ...MethodOptionFunc) (*GetHireEvaluationListResp, *Response, error) {
	if r.cli.mock.mockHireGetHireEvaluationList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#GetHireEvaluationList mock enable")
		return r.cli.mock.mockHireGetHireEvaluationList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireEvaluationList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/evaluations",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireEvaluationListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireEvaluationList mock HireGetHireEvaluationList method
func (r *Mock) MockHireGetHireEvaluationList(f func(ctx context.Context, request *GetHireEvaluationListReq, options ...MethodOptionFunc) (*GetHireEvaluationListResp, *Response, error)) {
	r.mockHireGetHireEvaluationList = f
}

// UnMockHireGetHireEvaluationList un-mock HireGetHireEvaluationList method
func (r *Mock) UnMockHireGetHireEvaluationList() {
	r.mockHireGetHireEvaluationList = nil
}

// GetHireEvaluationListReq ...
type GetHireEvaluationListReq struct {
	PageToken       *string `query:"page_token" json:"-"`        // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: eyJvZmZzZXQiOjEsInRpbWVzdGFtcCI6MTY0MDc2NTYzMjA4OCwiaWQiOm51bGx9
	PageSize        *int64  `query:"page_size" json:"-"`         // 每页获取记录数量, 最大100, 示例值: 100, 默认值: `10`, 最大值: `100`
	ApplicationID   *string `query:"application_id" json:"-"`    // 投递 ID, 示例值: 6875569957036738823
	UpdateStartTime *string `query:"update_start_time" json:"-"` // 最早更新时间, 毫秒级时间戳, 示例值: 1600843767338
	UpdateEndTime   *string `query:"update_end_time" json:"-"`   // 最晚更新时间, 毫秒级时间戳, 示例值: 1600843938726
	UserIDType      *IDType `query:"user_id_type" json:"-"`      // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以people_admin_id来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetHireEvaluationListResp ...
type GetHireEvaluationListResp struct {
	HasMore   bool                             `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                           `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetHireEvaluationListRespItem `json:"items,omitempty"`      // 简历评估信息列表
}

// GetHireEvaluationListRespItem ...
type GetHireEvaluationListRespItem struct {
	ID            string `json:"id,omitempty"`             // 评估 ID
	ApplicationID string `json:"application_id,omitempty"` // 投递 ID
	StageID       string `json:"stage_id,omitempty"`       // 投递阶段
	CreatorID     string `json:"creator_id,omitempty"`     // 创建人user_id
	EvaluatorID   string `json:"evaluator_id,omitempty"`   // 评估人user_id
	CommitStatus  int64  `json:"commit_status,omitempty"`  // 提交状态, 可选值有: 1: 已提交, 2: 未提交
	Conclusion    int64  `json:"conclusion,omitempty"`     // 评估结论, 可选值有: 1: 通过, 2: 未通过
	Content       string `json:"content,omitempty"`        // 评估详情
	CreateTime    string `json:"create_time,omitempty"`    // 创建时间
	UpdateTime    string `json:"update_time,omitempty"`    // 最近更新时间
}

// getHireEvaluationListResp ...
type getHireEvaluationListResp struct {
	Code  int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 错误描述
	Data  *GetHireEvaluationListResp `json:"data,omitempty"`
	Error *ErrorDetail               `json:"error,omitempty"`
}
