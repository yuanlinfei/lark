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

// UpdateDriveFileSubscription 根据订阅ID更新订阅状态
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-subscription/patch
func (r *DriveService) UpdateDriveFileSubscription(ctx context.Context, request *UpdateDriveFileSubscriptionReq, options ...MethodOptionFunc) (*UpdateDriveFileSubscriptionResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateDriveFileSubscription != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateDriveFileSubscription mock enable")
		return r.cli.mock.mockDriveUpdateDriveFileSubscription(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Drive",
		API:                 "UpdateDriveFileSubscription",
		Method:              "PATCH",
		URL:                 r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/subscriptions/:subscription_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(updateDriveFileSubscriptionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateDriveFileSubscription mock DriveUpdateDriveFileSubscription method
func (r *Mock) MockDriveUpdateDriveFileSubscription(f func(ctx context.Context, request *UpdateDriveFileSubscriptionReq, options ...MethodOptionFunc) (*UpdateDriveFileSubscriptionResp, *Response, error)) {
	r.mockDriveUpdateDriveFileSubscription = f
}

// UnMockDriveUpdateDriveFileSubscription un-mock DriveUpdateDriveFileSubscription method
func (r *Mock) UnMockDriveUpdateDriveFileSubscription() {
	r.mockDriveUpdateDriveFileSubscription = nil
}

// UpdateDriveFileSubscriptionReq ...
type UpdateDriveFileSubscriptionReq struct {
	FileToken      string   `path:"file_token" json:"-"`      // 文档token, 示例值: "doxcnxxxxxxxxxxxxxxxxxxxxxx"
	SubscriptionID string   `path:"subscription_id" json:"-"` // 订阅关系ID, 示例值: "1234567890987654321"
	IsSubscribe    bool     `json:"is_subscribe,omitempty"`   // 是否订阅, 示例值: true
	FileType       FileType `json:"file_type,omitempty"`      // 文档类型, 示例值: "doc", 可选值有: doc: 文档, docx: 新版文档, wiki: 知识库wiki
}

// UpdateDriveFileSubscriptionResp ...
type UpdateDriveFileSubscriptionResp struct {
	Subscription *UpdateDriveFileSubscriptionRespSubscription `json:"subscription,omitempty"` // 本次修改的文档订阅信息
}

// UpdateDriveFileSubscriptionRespSubscription ...
type UpdateDriveFileSubscriptionRespSubscription struct {
	SubscriptionID   string   `json:"subscription_id,omitempty"`   // 订阅关系ID
	SubscriptionType string   `json:"subscription_type,omitempty"` // 订阅类型, 可选值有: comment_update: 评论更新
	IsSubcribe       bool     `json:"is_subcribe,omitempty"`       // 是否订阅
	FileType         FileType `json:"file_type,omitempty"`         // 文档类型, 可选值有: doc: 文档, docx: 新版文档, wiki: 知识库wiki
}

// updateDriveFileSubscriptionResp ...
type updateDriveFileSubscriptionResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *UpdateDriveFileSubscriptionResp `json:"data,omitempty"`
}
