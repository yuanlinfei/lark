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

// CreateHelpdeskCategory 该接口用于创建知识库分类。
//
// 注意事项:
// user_access_token 访问, 需要操作者是当前服务台的客服、管理员或所有者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/create
// new doc: https://open.feishu.cn/document/server-docs/helpdesk-v1/faq-management/category/create
func (r *HelpdeskService) CreateHelpdeskCategory(ctx context.Context, request *CreateHelpdeskCategoryReq, options ...MethodOptionFunc) (*CreateHelpdeskCategoryResp, *Response, error) {
	if r.cli.mock.mockHelpdeskCreateHelpdeskCategory != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Helpdesk#CreateHelpdeskCategory mock enable")
		return r.cli.mock.mockHelpdeskCreateHelpdeskCategory(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "CreateHelpdeskCategory",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/categories",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(createHelpdeskCategoryResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskCreateHelpdeskCategory mock HelpdeskCreateHelpdeskCategory method
func (r *Mock) MockHelpdeskCreateHelpdeskCategory(f func(ctx context.Context, request *CreateHelpdeskCategoryReq, options ...MethodOptionFunc) (*CreateHelpdeskCategoryResp, *Response, error)) {
	r.mockHelpdeskCreateHelpdeskCategory = f
}

// UnMockHelpdeskCreateHelpdeskCategory un-mock HelpdeskCreateHelpdeskCategory method
func (r *Mock) UnMockHelpdeskCreateHelpdeskCategory() {
	r.mockHelpdeskCreateHelpdeskCategory = nil
}

// CreateHelpdeskCategoryReq ...
type CreateHelpdeskCategoryReq struct {
	Name     string  `json:"name,omitempty"`      // 名称, 示例值: "创建团队和邀请成员"
	ParentID string  `json:"parent_id,omitempty"` // 父知识库分类ID, 示例值: "0"
	Language *string `json:"language,omitempty"`  // 语言, 示例值: "zh_cn"
}

// CreateHelpdeskCategoryResp ...
type CreateHelpdeskCategoryResp struct {
	Category *HelpdeskCategory `json:"category,omitempty"` // 知识库分类
}

// createHelpdeskCategoryResp ...
type createHelpdeskCategoryResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *CreateHelpdeskCategoryResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}
