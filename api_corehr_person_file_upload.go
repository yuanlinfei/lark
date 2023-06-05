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

// UploadCoreHrPersonFile 上传文件。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/person/upload
func (r *CoreHrService) UploadCoreHrPersonFile(ctx context.Context, request *UploadCoreHrPersonFileReq, options ...MethodOptionFunc) (*UploadCoreHrPersonFileResp, *Response, error) {
	if r.cli.mock.mockCoreHrUploadCoreHrPersonFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#UploadCoreHrPersonFile mock enable")
		return r.cli.mock.mockCoreHrUploadCoreHrPersonFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "UploadCoreHrPersonFile",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/persons/upload",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(uploadCoreHrPersonFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrUploadCoreHrPersonFile mock CoreHrUploadCoreHrPersonFile method
func (r *Mock) MockCoreHrUploadCoreHrPersonFile(f func(ctx context.Context, request *UploadCoreHrPersonFileReq, options ...MethodOptionFunc) (*UploadCoreHrPersonFileResp, *Response, error)) {
	r.mockCoreHrUploadCoreHrPersonFile = f
}

// UnMockCoreHrUploadCoreHrPersonFile un-mock CoreHrUploadCoreHrPersonFile method
func (r *Mock) UnMockCoreHrUploadCoreHrPersonFile() {
	r.mockCoreHrUploadCoreHrPersonFile = nil
}

// UploadCoreHrPersonFileReq ...
type UploadCoreHrPersonFileReq struct {
	FileContent io.Reader `json:"file_content,omitempty"` // 文件二进制内容, 示例值: file binary
	FileName    string    `json:"file_name,omitempty"`    // 文件名称, 示例值: "个人信息"
}

// UploadCoreHrPersonFileResp ...
type UploadCoreHrPersonFileResp struct {
	ID string `json:"id,omitempty"` // 上传文件ID
}

// uploadCoreHrPersonFileResp ...
type uploadCoreHrPersonFileResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *UploadCoreHrPersonFileResp `json:"data,omitempty"`
}