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

// DeleteLingoEntity 通过 entity_id 删除已有的词条, 无需经过词典管理员审核。因此, 调用该接口时应当慎重操作。
//
// 也支持通过 provider 和 outer_id 删除对应的词条。此时路径中的 entity_id 为固定的 enterprise_0
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/lingo-v1/entity/delete
func (r *LingoService) DeleteLingoEntity(ctx context.Context, request *DeleteLingoEntityReq, options ...MethodOptionFunc) (*DeleteLingoEntityResp, *Response, error) {
	if r.cli.mock.mockLingoDeleteLingoEntity != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Lingo#DeleteLingoEntity mock enable")
		return r.cli.mock.mockLingoDeleteLingoEntity(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Lingo",
		API:                   "DeleteLingoEntity",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/lingo/v1/entities/:entity_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteLingoEntityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockLingoDeleteLingoEntity mock LingoDeleteLingoEntity method
func (r *Mock) MockLingoDeleteLingoEntity(f func(ctx context.Context, request *DeleteLingoEntityReq, options ...MethodOptionFunc) (*DeleteLingoEntityResp, *Response, error)) {
	r.mockLingoDeleteLingoEntity = f
}

// UnMockLingoDeleteLingoEntity un-mock LingoDeleteLingoEntity method
func (r *Mock) UnMockLingoDeleteLingoEntity() {
	r.mockLingoDeleteLingoEntity = nil
}

// DeleteLingoEntityReq ...
type DeleteLingoEntityReq struct {
	EntityID string  `path:"entity_id" json:"-"` // 词条 ID, 示例值: "enterprise_43742132363"
	Provider *string `query:"provider" json:"-"` // 外部系统（使用时需要将路径中的词条 ID 固定为: enterprise_0, 且提供 provider 和 outer_id）, 示例值: 星云, 长度范围: `2` ～ `32` 字符
	OuterID  *string `query:"outer_id" json:"-"` // 词条在外部系统中对应的唯一 ID（使用时需要将路径中的词条 ID 固定为: enterprise_0, 且提供 provider 和 outer_id）, 示例值: 123aaa, 长度范围: `1` ～ `64` 字符
}

// DeleteLingoEntityResp ...
type DeleteLingoEntityResp struct {
}

// deleteLingoEntityResp ...
type deleteLingoEntityResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *DeleteLingoEntityResp `json:"data,omitempty"`
}
