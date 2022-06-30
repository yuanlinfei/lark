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

// CreateTaskFollower 该接口用于创建任务关注人。可以一次性添加多位关注人。关注人ID要使用表示用户的ID。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/create
func (r *TaskService) CreateTaskFollower(ctx context.Context, request *CreateTaskFollowerReq, options ...MethodOptionFunc) (*CreateTaskFollowerResp, *Response, error) {
	if r.cli.mock.mockTaskCreateTaskFollower != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#CreateTaskFollower mock enable")
		return r.cli.mock.mockTaskCreateTaskFollower(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CreateTaskFollower",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/followers",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createTaskFollowerResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskCreateTaskFollower mock TaskCreateTaskFollower method
func (r *Mock) MockTaskCreateTaskFollower(f func(ctx context.Context, request *CreateTaskFollowerReq, options ...MethodOptionFunc) (*CreateTaskFollowerResp, *Response, error)) {
	r.mockTaskCreateTaskFollower = f
}

// UnMockTaskCreateTaskFollower un-mock TaskCreateTaskFollower method
func (r *Mock) UnMockTaskCreateTaskFollower() {
	r.mockTaskCreateTaskFollower = nil
}

// CreateTaskFollowerReq ...
type CreateTaskFollowerReq struct {
	TaskID     string   `path:"task_id" json:"-"`       // 任务 ID, 示例值: "83912691-2e43-47fc-94a4-d512e03984fa"
	UserIDType *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: <md-enum>, <md-enum-item key="open_id" >用户的 open id</md-enum-item>, <md-enum-item key="union_id" >用户的 union id</md-enum-item>, <md-enum-item key="user_id" >用户的 user id</md-enum-item>, </md-enum>, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ID         *string  `json:"id,omitempty"`           // 任务关注人 ID, 示例值: "ou_99e1a581b36ecc4862cbfbce473f3123"
	IDList     []string `json:"id_list,omitempty"`      // 要添加的关注人ID列表, 示例值: [, "ou_550cc75233d8b7b9fcbdad65f34433f4", "ou_d1e9d27cf3235b40ca9a67c67ef088b0", ]
}

// CreateTaskFollowerResp ...
type CreateTaskFollowerResp struct {
	Follower *CreateTaskFollowerRespFollower `json:"follower,omitempty"` // 创建后的任务关注者
}

// CreateTaskFollowerRespFollower ...
type CreateTaskFollowerRespFollower struct {
	ID     string   `json:"id,omitempty"`      // 任务关注人 ID
	IDList []string `json:"id_list,omitempty"` // 要添加的关注人ID列表
}

// createTaskFollowerResp ...
type createTaskFollowerResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *CreateTaskFollowerResp `json:"data,omitempty"`
}