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

// DeleteVCRoomLevel 该接口可以用来删除某个会议室层级。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_level/del
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/room_level/del
func (r *VCService) DeleteVCRoomLevel(ctx context.Context, request *DeleteVCRoomLevelReq, options ...MethodOptionFunc) (*DeleteVCRoomLevelResp, *Response, error) {
	if r.cli.mock.mockVCDeleteVCRoomLevel != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] VC#DeleteVCRoomLevel mock enable")
		return r.cli.mock.mockVCDeleteVCRoomLevel(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "DeleteVCRoomLevel",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/room_levels/del",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteVCRoomLevelResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCDeleteVCRoomLevel mock VCDeleteVCRoomLevel method
func (r *Mock) MockVCDeleteVCRoomLevel(f func(ctx context.Context, request *DeleteVCRoomLevelReq, options ...MethodOptionFunc) (*DeleteVCRoomLevelResp, *Response, error)) {
	r.mockVCDeleteVCRoomLevel = f
}

// UnMockVCDeleteVCRoomLevel un-mock VCDeleteVCRoomLevel method
func (r *Mock) UnMockVCDeleteVCRoomLevel() {
	r.mockVCDeleteVCRoomLevel = nil
}

// DeleteVCRoomLevelReq ...
type DeleteVCRoomLevelReq struct {
	RoomLevelID string `json:"room_level_id,omitempty"` // 层级ID, 示例值: "omb_4ad1a2c7a2fbc5fc9570f38456931293", 长度范围: `1` ～ `100` 字符
	DeleteChild *bool  `json:"delete_child,omitempty"`  // 是否删除所有子层级, 示例值: false
}

// DeleteVCRoomLevelResp ...
type DeleteVCRoomLevelResp struct {
}

// deleteVCRoomLevelResp ...
type deleteVCRoomLevelResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *DeleteVCRoomLevelResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
