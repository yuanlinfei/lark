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

// GetCoreHRCustomFieldList 获取「飞书人事」具体对象下的自定义字段列表。注: 在「人员档案信息配置」-「个人信息」功能中添加的分组, 实际上是一个自定义对象, 可以通过该接口查询自定义对象的所有字段。使用方式可参考[【操作手册】如何通过 OpenAPI 维护自定义字段](https://feishu.feishu.cn/docx/QlUudBfCtosWMbxx3vxcOFDknn7)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/query
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/custom_field/query
func (r *CoreHRService) GetCoreHRCustomFieldList(ctx context.Context, request *GetCoreHRCustomFieldListReq, options ...MethodOptionFunc) (*GetCoreHRCustomFieldListResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRCustomFieldList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRCustomFieldList mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRCustomFieldList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRCustomFieldList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/custom_fields/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRCustomFieldListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRCustomFieldList mock CoreHRGetCoreHRCustomFieldList method
func (r *Mock) MockCoreHRGetCoreHRCustomFieldList(f func(ctx context.Context, request *GetCoreHRCustomFieldListReq, options ...MethodOptionFunc) (*GetCoreHRCustomFieldListResp, *Response, error)) {
	r.mockCoreHRGetCoreHRCustomFieldList = f
}

// UnMockCoreHRGetCoreHRCustomFieldList un-mock CoreHRGetCoreHRCustomFieldList method
func (r *Mock) UnMockCoreHRGetCoreHRCustomFieldList() {
	r.mockCoreHRGetCoreHRCustomFieldList = nil
}

// GetCoreHRCustomFieldListReq ...
type GetCoreHRCustomFieldListReq struct {
	ObjectApiNameList []string `query:"object_api_name_list" json:"-"` // 所属对象 apiname, 支持一个或多个, 当前数量限制为 20 个, 示例值: person
}

// GetCoreHRCustomFieldListResp ...
type GetCoreHRCustomFieldListResp struct {
	Items []*GetCoreHRCustomFieldListRespItem `json:"items,omitempty"` // 自定义字段列表
}

// GetCoreHRCustomFieldListRespItem ...
type GetCoreHRCustomFieldListRespItem struct {
	CustomApiName string                                       `json:"custom_api_name,omitempty"` // 自定义字段 apiname, 即自定义字段的唯一标识
	Name          *GetCoreHRCustomFieldListRespItemName        `json:"name,omitempty"`            // 自定义字段名称
	Description   *GetCoreHRCustomFieldListRespItemDescription `json:"description,omitempty"`     // 描述
	IsOpen        bool                                         `json:"is_open,omitempty"`         // 是否启用
	IsRequired    bool                                         `json:"is_required,omitempty"`     // 是否必填
	IsUnique      bool                                         `json:"is_unique,omitempty"`       // 是否唯一
	ObjectApiName string                                       `json:"object_api_name,omitempty"` // 所属对象 apiname
	Type          int64                                        `json:"type,omitempty"`            // 自定义字段类型, 可选值有: 1: 文本 Text, “文本”和“超链接”属于该类型, 2: 布尔 Boolean, 3: 数字 Number, 4: 枚举 Option, “单选”和“多选”为该类型, 5: 查找 Lookup, “人员（单选）”、“人员（多选）”和个人信息中的自定义分组为该类型, 7: 日期时间 Date time, 8: 附件 Attachment, “附件单选”和“附件多选”为该类型
	CreateTime    string                                       `json:"create_time,omitempty"`     // 创建时间, 秒级时间戳
	UpdateTime    string                                       `json:"update_time,omitempty"`     // 更新时间, 秒级时间戳
}

// GetCoreHRCustomFieldListRespItemDescription ...
type GetCoreHRCustomFieldListRespItemDescription struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetCoreHRCustomFieldListRespItemName ...
type GetCoreHRCustomFieldListRespItemName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// getCoreHRCustomFieldListResp ...
type getCoreHRCustomFieldListResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHRCustomFieldListResp `json:"data,omitempty"`
}
