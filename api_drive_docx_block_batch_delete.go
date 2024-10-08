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

// BatchDeleteDocxBlock 指定需要操作的块, 删除其指定范围的子块。如果操作成功, 接口将返回应用删除操作后的文档版本号。
//
// 应用频率限制: 单个应用调用频率上限为每秒 3 次, 超过该频率限制, 接口将返回 HTTP 状态码 400 及错误码 99991400；
// 文档频率限制: 单篇文档并发编辑上限为每秒 3 次, 超过该频率限制, 接口将返回 HTTP 状态码 429, 编辑操作包括:
// - 创建块
// - 删除块
// - 更新块
// - 批量更新块
// 当请求被限频, 应用需要处理限频状态码, 并使用指数退避算法或其它一些频控策略降低对 API 的调用速率。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block-children/batch_delete
// new doc: https://open.feishu.cn/document/server-docs/docs/docs/docx-v1/document-block/batch_delete
func (r *DriveService) BatchDeleteDocxBlock(ctx context.Context, request *BatchDeleteDocxBlockReq, options ...MethodOptionFunc) (*BatchDeleteDocxBlockResp, *Response, error) {
	if r.cli.mock.mockDriveBatchDeleteDocxBlock != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#BatchDeleteDocxBlock mock enable")
		return r.cli.mock.mockDriveBatchDeleteDocxBlock(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "BatchDeleteDocxBlock",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children/batch_delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchDeleteDocxBlockResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveBatchDeleteDocxBlock mock DriveBatchDeleteDocxBlock method
func (r *Mock) MockDriveBatchDeleteDocxBlock(f func(ctx context.Context, request *BatchDeleteDocxBlockReq, options ...MethodOptionFunc) (*BatchDeleteDocxBlockResp, *Response, error)) {
	r.mockDriveBatchDeleteDocxBlock = f
}

// UnMockDriveBatchDeleteDocxBlock un-mock DriveBatchDeleteDocxBlock method
func (r *Mock) UnMockDriveBatchDeleteDocxBlock() {
	r.mockDriveBatchDeleteDocxBlock = nil
}

// BatchDeleteDocxBlockReq ...
type BatchDeleteDocxBlockReq struct {
	DocumentID         string  `path:"document_id" json:"-"`           // 文档的唯一标识。点击[这里](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-overview)了解如何获取文档的 `document_id`, 示例值: "doxcnePuYufKa49ISjhD8Iabcef"
	BlockID            string  `path:"block_id" json:"-"`              // 父 Block 的唯一标识。你可通过调用[获取文档所有块](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block/list)获取块的 block_id, 注意: 此接口不支持删除表格（Table）和分栏（Grid）块的子块。你需通过[更新块的内容](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block/patch)的对应请求实现, 此接口不支持删除高亮（Callout）块的全部子块, 示例值: "doxcnO6UW6wAw2qIcYf4hZabcef"
	DocumentRevisionID *int64  `query:"document_revision_id" json:"-"` // 要操作的文档版本。-1 表示文档最新版本。文档创建后, 版本为 1。你需确保你已拥有文档的编辑权限。你可通过调用[获取文档基本信息](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document/get)获取文档的最新 revision_id, 示例值:1, 默认值: `-1`, 最小值: `-1`
	ClientToken        *string `query:"client_token" json:"-"`         // 操作的唯一标识, 与接口返回值的 client_token 相对应, 用于幂等的进行更新操作。此值为空表示将发起一次新的请求, 此值非空表示幂等的进行更新操作, 示例值: "fe599b60-450f-46ff-b2ef-9f6675625b97"
	StartIndex         int64   `json:"start_index"`          // 删除的起始索引（操作区间左闭右开）, 示例值: 0, 最小值: `0`
	EndIndex           int64   `json:"end_index"`            // 删除的末尾索引（操作区间左闭右开）, 示例值: 1, 最小值: `1`
}

// BatchDeleteDocxBlockResp ...
type BatchDeleteDocxBlockResp struct {
	DocumentRevisionID int64  `json:"document_revision_id,omitempty"` // 当前删除操作成功后文档的版本号
	ClientToken        string `json:"client_token,omitempty"`         // 操作的唯一标识, 更新请求中使用此值表示幂等的进行此次更新
}

// batchDeleteDocxBlockResp ...
type batchDeleteDocxBlockResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *BatchDeleteDocxBlockResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
