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

// CreateTaskComment 该接口用于创建和回复任务的评论。当parent_id字段为0时, 为创建评论；当parent_id不为0时, 为回复某条评论
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/create
func (r *TaskService) CreateTaskComment(ctx context.Context, request *CreateTaskCommentReq, options ...MethodOptionFunc) (*CreateTaskCommentResp, *Response, error) {
	if r.cli.mock.mockTaskCreateTaskComment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#CreateTaskComment mock enable")
		return r.cli.mock.mockTaskCreateTaskComment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CreateTaskComment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/comments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createTaskCommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskCreateTaskComment mock TaskCreateTaskComment method
func (r *Mock) MockTaskCreateTaskComment(f func(ctx context.Context, request *CreateTaskCommentReq, options ...MethodOptionFunc) (*CreateTaskCommentResp, *Response, error)) {
	r.mockTaskCreateTaskComment = f
}

// UnMockTaskCreateTaskComment un-mock TaskCreateTaskComment method
func (r *Mock) UnMockTaskCreateTaskComment() {
	r.mockTaskCreateTaskComment = nil
}

// CreateTaskCommentReq ...
type CreateTaskCommentReq struct {
	TaskID          string  `path:"task_id" json:"-"`            // 任务 ID, 示例值: "83912691-2e43-47fc-94a4-d512e03984fa"
	Content         *string `json:"content,omitempty"`           // 评论内容, <md-alert>, 评论内容和富文本评论内容同时存在时只使用富文本评论内容, </md-alert>, 示例值: "举杯邀明月, 对影成三人", 长度范围: `0` ～ `65536` 字符
	ParentID        *string `json:"parent_id,omitempty"`         // 评论的父ID, 创建评论时若不为空则为某条评论的回复, 若为空则不是回复, 示例值: "6937231762296684564"
	ID              *string `json:"id,omitempty"`                // 评论ID, 由飞书服务器发号, 示例值: "6937231762296684564"
	CreateMilliTime *string `json:"create_milli_time,omitempty"` // 评论创建的时间戳, 单位为毫秒, 用于展示, 创建时不用填写, 示例值: "1657075055135"
	RichContent     *string `json:"rich_content,omitempty"`      // 富文本评论内容。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module), 示例值: "举杯邀明月, 对影成三人<at id=7058204817822318612></at>", 长度范围: `0` ～ `65536` 字符
}

// CreateTaskCommentResp ...
type CreateTaskCommentResp struct {
	Comment *CreateTaskCommentRespComment `json:"comment,omitempty"` // 返回创建好的任务评论
}

// CreateTaskCommentRespComment ...
type CreateTaskCommentRespComment struct {
	Content         string `json:"content,omitempty"`           // 评论内容, <md-alert>, 评论内容和富文本评论内容同时存在时只使用富文本评论内容, </md-alert>
	ParentID        string `json:"parent_id,omitempty"`         // 评论的父ID, 创建评论时若不为空则为某条评论的回复, 若为空则不是回复
	ID              string `json:"id,omitempty"`                // 评论ID, 由飞书服务器发号
	CreateMilliTime string `json:"create_milli_time,omitempty"` // 评论创建的时间戳, 单位为毫秒, 用于展示, 创建时不用填写
	RichContent     string `json:"rich_content,omitempty"`      // 富文本评论内容。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module)
}

// createTaskCommentResp ...
type createTaskCommentResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *CreateTaskCommentResp `json:"data,omitempty"`
}
