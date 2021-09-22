// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteSheetDataValidationDropdown
//
// 该接口根据 spreadsheetToken 、range 移除选定数据范围单元格的下拉列表设置，但保留选项文本。单个删除范围不超过5000单元格。单次请求range最大数量100个。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/delete-datavalidation
func (r *DriveService) DeleteSheetDataValidationDropdown(ctx context.Context, request *DeleteSheetDataValidationDropdownReq, options ...MethodOptionFunc) (*DeleteSheetDataValidationDropdownResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteSheetDataValidationDropdown != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DeleteSheetDataValidationDropdown mock enable")
		return r.cli.mock.mockDriveDeleteSheetDataValidationDropdown(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteSheetDataValidationDropdown",
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteSheetDataValidationDropdownResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveDeleteSheetDataValidationDropdown(f func(ctx context.Context, request *DeleteSheetDataValidationDropdownReq, options ...MethodOptionFunc) (*DeleteSheetDataValidationDropdownResp, *Response, error)) {
	r.mockDriveDeleteSheetDataValidationDropdown = f
}

func (r *Mock) UnMockDriveDeleteSheetDataValidationDropdown() {
	r.mockDriveDeleteSheetDataValidationDropdown = nil
}

type DeleteSheetDataValidationDropdownReq struct {
	SpreadSheetToken     string                                                     `path:"spreadsheetToken" json:"-"`      // spreadsheet 的 token，获取方式见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	DataValidationRanges []*DeleteSheetDataValidationDropdownReqDataValidationRange `json:"dataValidationRanges,omitempty"` // 范围数组，每个range 最大单元格数量5000，每个range独立执行，一个range的失败不影响其他range的执行。返回结果会返回每个range的执行结果
}

type DeleteSheetDataValidationDropdownReqDataValidationRange struct {
	Range             string  `json:"range,omitempty"`             // 查询范围，包含 sheetId 与单元格范围两部分，目前支持四种索引方式，详见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	DataValidationIDs []int64 `json:"dataValidationIds,omitempty"` // 指定需要删除的dataValidationIds
}

type deleteSheetDataValidationDropdownResp struct {
	Code int64                                  `json:"code,omitempty"` // 状态码，0代表成功
	Msg  *string                                `json:"msg,omitempty"`  // 状态信息
	Data *DeleteSheetDataValidationDropdownResp `json:"data,omitempty"`
}

type DeleteSheetDataValidationDropdownResp struct {
	RangeResults []*DeleteSheetDataValidationDropdownRespRangeResult `json:"rangeResults,omitempty"`
}

type DeleteSheetDataValidationDropdownRespRangeResult struct {
	Range        string  `json:"range,omitempty"`        // 执行的range,与请求入参中的range 对应
	Msg          *string `json:"msg,omitempty"`          // 结果信息
	Success      bool    `json:"success,omitempty"`      // 执行结果
	UpdatedCells int64   `json:"updatedCells,omitempty"` // 影响的单元格数量
}
