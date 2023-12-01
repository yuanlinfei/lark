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

// SearchVCRoomLevel 该接口可以用来搜索会议室层级, 支持使用自定义会议室层级 ID 进行查询。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_level/search
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/room_level/search
func (r *VCService) SearchVCRoomLevel(ctx context.Context, request *SearchVCRoomLevelReq, options ...MethodOptionFunc) (*SearchVCRoomLevelResp, *Response, error) {
	if r.cli.mock.mockVCSearchVCRoomLevel != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#SearchVCRoomLevel mock enable")
		return r.cli.mock.mockVCSearchVCRoomLevel(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "SearchVCRoomLevel",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/room_levels/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(searchVCRoomLevelResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCSearchVCRoomLevel mock VCSearchVCRoomLevel method
func (r *Mock) MockVCSearchVCRoomLevel(f func(ctx context.Context, request *SearchVCRoomLevelReq, options ...MethodOptionFunc) (*SearchVCRoomLevelResp, *Response, error)) {
	r.mockVCSearchVCRoomLevel = f
}

// UnMockVCSearchVCRoomLevel un-mock VCSearchVCRoomLevel method
func (r *Mock) UnMockVCSearchVCRoomLevel() {
	r.mockVCSearchVCRoomLevel = nil
}

// SearchVCRoomLevelReq ...
type SearchVCRoomLevelReq struct {
	CustomLevelIDs []string `query:"custom_level_ids" json:"-"` // 用于查询指定会议室层级的自定义会议室层级ID, 示例值: 1000, 1001
}

// SearchVCRoomLevelResp ...
type SearchVCRoomLevelResp struct {
	LevelIDs []string `json:"level_ids,omitempty"` // 层级ID列表
}

// searchVCRoomLevelResp ...
type searchVCRoomLevelResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *SearchVCRoomLevelResp `json:"data,omitempty"`
}
