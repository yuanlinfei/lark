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

// GetSpreadsheet 该接口用于获取电子表格的基础信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet/get
func (r *DriveService) GetSpreadsheet(ctx context.Context, request *GetSpreadsheetReq, options ...MethodOptionFunc) (*GetSpreadsheetResp, *Response, error) {
	if r.cli.mock.mockDriveGetSpreadsheet != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetSpreadsheet mock enable")
		return r.cli.mock.mockDriveGetSpreadsheet(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetSpreadsheet",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getSpreadsheetResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetSpreadsheet mock DriveGetSpreadsheet method
func (r *Mock) MockDriveGetSpreadsheet(f func(ctx context.Context, request *GetSpreadsheetReq, options ...MethodOptionFunc) (*GetSpreadsheetResp, *Response, error)) {
	r.mockDriveGetSpreadsheet = f
}

// UnMockDriveGetSpreadsheet un-mock DriveGetSpreadsheet method
func (r *Mock) UnMockDriveGetSpreadsheet() {
	r.mockDriveGetSpreadsheet = nil
}

// GetSpreadsheetReq ...
type GetSpreadsheetReq struct {
	SpreadSheetToken string  `path:"spreadsheet_token" json:"-"` // 表格的token, 示例值: "shtxxxxxxxxxxxxxxx"
	UserIDType       *IDType `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetSpreadsheetResp ...
type GetSpreadsheetResp struct {
	Spreadsheet *GetSpreadsheetRespSpreadsheet `json:"spreadsheet,omitempty"` // 电子表格属性
}

// GetSpreadsheetRespSpreadsheet ...
type GetSpreadsheetRespSpreadsheet struct {
	Title   string `json:"title,omitempty"`    // 电子表格标题
	OwnerID string `json:"owner_id,omitempty"` // 电子表格owner
	Token   string `json:"token,omitempty"`    // 电子表格token
	URL     string `json:"url,omitempty"`      // 电子表格url
}

// getSpreadsheetResp ...
type getSpreadsheetResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetSpreadsheetResp `json:"data,omitempty"`
}
