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

// InsertSheetDimensionRange 该接口用于根据 spreadsheetToken 和维度信息 插入空行/列。
//
// 如 startIndex=3, endIndex=7, 则从第 4 行开始开始插入行列, 一直到第 7 行, 共插入 4 行；单次操作不超过5000行或列。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQjMzUjL0IzM14CNyMTN
// new doc: https://open.feishu.cn/document/server-docs/docs/sheets-v3/sheet-rowcol/insert-rows-or-columns
func (r *DriveService) InsertSheetDimensionRange(ctx context.Context, request *InsertSheetDimensionRangeReq, options ...MethodOptionFunc) (*InsertSheetDimensionRangeResp, *Response, error) {
	if r.cli.mock.mockDriveInsertSheetDimensionRange != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#InsertSheetDimensionRange mock enable")
		return r.cli.mock.mockDriveInsertSheetDimensionRange(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "InsertSheetDimensionRange",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/insert_dimension_range",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(insertSheetDimensionRangeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveInsertSheetDimensionRange mock DriveInsertSheetDimensionRange method
func (r *Mock) MockDriveInsertSheetDimensionRange(f func(ctx context.Context, request *InsertSheetDimensionRangeReq, options ...MethodOptionFunc) (*InsertSheetDimensionRangeResp, *Response, error)) {
	r.mockDriveInsertSheetDimensionRange = f
}

// UnMockDriveInsertSheetDimensionRange un-mock DriveInsertSheetDimensionRange method
func (r *Mock) UnMockDriveInsertSheetDimensionRange() {
	r.mockDriveInsertSheetDimensionRange = nil
}

// InsertSheetDimensionRangeReq ...
type InsertSheetDimensionRangeReq struct {
	SpreadSheetToken string                                 `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token, 获取方式见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	Dimension        *InsertSheetDimensionRangeReqDimension `json:"dimension,omitempty"`       // 需要插入行列的维度信息
	InheritStyle     *string                                `json:"inheritStyle,omitempty"`    // BEFORE 或 AFTER, 不填为不继承 style
}

// InsertSheetDimensionRangeReqDimension ...
type InsertSheetDimensionRangeReqDimension struct {
	SheetID        string  `json:"sheetId,omitempty"`        // sheet 的 Id
	MajorDimension *string `json:"majorDimension,omitempty"` // 默认 ROWS, 可选 ROWS、COLUMNS
	StartIndex     int64   `json:"startIndex"`               // 开始的位置
	EndIndex       int64   `json:"endIndex,omitempty"`       // 结束的位置
}

// InsertSheetDimensionRangeResp ...
type InsertSheetDimensionRangeResp struct {
}

// insertSheetDimensionRangeResp ...
type insertSheetDimensionRangeResp struct {
	Code  int64                          `json:"code,omitempty"`
	Msg   string                         `json:"msg,omitempty"`
	Data  *InsertSheetDimensionRangeResp `json:"data,omitempty"`
	Error *ErrorDetail                   `json:"error,omitempty"`
}
