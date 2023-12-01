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

// SearchApprovalTask 该接口通过不同条件查询审批系统中符合条件的审批任务列表。
//
// 如需了解审批任务各状态的说明, 以及审批任务状态变更事件, 可参见[审批任务状态变更](https://open.feishu.cn/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/common-event/approval-task-event)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/search
// new doc: https://open.feishu.cn/document/server-docs/approval-v4/approval-search/search
func (r *ApprovalService) SearchApprovalTask(ctx context.Context, request *SearchApprovalTaskReq, options ...MethodOptionFunc) (*SearchApprovalTaskResp, *Response, error) {
	if r.cli.mock.mockApprovalSearchApprovalTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#SearchApprovalTask mock enable")
		return r.cli.mock.mockApprovalSearchApprovalTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "SearchApprovalTask",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/tasks/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchApprovalTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalSearchApprovalTask mock ApprovalSearchApprovalTask method
func (r *Mock) MockApprovalSearchApprovalTask(f func(ctx context.Context, request *SearchApprovalTaskReq, options ...MethodOptionFunc) (*SearchApprovalTaskResp, *Response, error)) {
	r.mockApprovalSearchApprovalTask = f
}

// UnMockApprovalSearchApprovalTask un-mock ApprovalSearchApprovalTask method
func (r *Mock) UnMockApprovalSearchApprovalTask() {
	r.mockApprovalSearchApprovalTask = nil
}

// SearchApprovalTaskReq ...
type SearchApprovalTaskReq struct {
	PageSize           *int64   `query:"page_size" json:"-"`            // 分页大小, 示例值: 10, 默认值: `10`, 取值范围: `5` ～ `200`
	PageToken          *string  `query:"page_token" json:"-"`           // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: nF1ZXJ5VGhlbkZldGNoCgAAAAAA6PZwFmUzSldvTC1yU
	UserIDType         *IDType  `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	UserID             *string  `json:"user_id,omitempty"`              // 根据x_user_type填写审批人id, 示例值: "lwiu098wj"
	ApprovalCode       *string  `json:"approval_code,omitempty"`        // 审批定义 code, 示例值: "EB828003-9FFE-4B3F-AA50-2E199E2ED942"
	InstanceCode       *string  `json:"instance_code,omitempty"`        // 审批实例 code, 示例值: "EB828003-9FFE-4B3F-AA50-2E199E2ED943"
	InstanceExternalID *string  `json:"instance_external_id,omitempty"` // 审批实例第三方 id 注: 和 approval_code 取并集, 示例值: "EB828003-9FFE-4B3F-AA50-2E199E2ED976"
	GroupExternalID    *string  `json:"group_external_id,omitempty"`    // 审批定义分组第三方 id 注: 和 instance_code 取并集, 示例值: "1234567"
	TaskTitle          *string  `json:"task_title,omitempty"`           // 审批任务标题（只有第三方审批有）, 示例值: "test"
	TaskStatus         *string  `json:"task_status,omitempty"`          // 审批任务状态, 注: 若不设置, 查询全部状态 若不在集合中, 报错, 示例值: "PENDING", 可选值有: PENDING: 审批中, REJECTED: 拒绝, APPROVED: 通过, TRANSFERRED: 转交, DONE: 已完成, RM_REPEAT: 去重, PROCESSED: 已处理, ALL: 所有状态
	TaskStartTimeFrom  *string  `json:"task_start_time_from,omitempty"` // 任务查询开始时间（unix毫秒时间戳）, 示例值: "1547654251506"
	TaskStartTimeTo    *string  `json:"task_start_time_to,omitempty"`   // 任务查询结束时间 (unix毫秒时间戳), 示例值: "1547654251506"
	Locale             *string  `json:"locale,omitempty"`               // 地区, 示例值: "zh-CN", 可选值有: zh-CN: 中文, en-US: 英文, ja-JP: 日文
	TaskStatusList     []string `json:"task_status_list,omitempty"`     // 可选择task_status中的多个状态, 当填写此参数时, task_status失效, 示例值: ["PENDING"]
	Order              *int64   `json:"order,omitempty"`                // 按任务时间排序, 示例值: 2, 可选值有: 0: 按update_time倒排, 1: 按update_time正排, 2: 按start_time倒排, 3: 按start_time正排, 默认值: `2`
}

// SearchApprovalTaskResp ...
type SearchApprovalTaskResp struct {
	Count     int64                         `json:"count,omitempty"`      // 查询返回条数
	TaskList  []*SearchApprovalTaskRespTask `json:"task_list,omitempty"`  // 审批任务列表
	PageToken string                        `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                          `json:"has_more,omitempty"`   // 是否还有更多项
}

// SearchApprovalTaskRespTask ...
type SearchApprovalTaskRespTask struct {
	Approval *SearchApprovalTaskRespTaskApproval `json:"approval,omitempty"` // 审批定义
	Group    *SearchApprovalTaskRespTaskGroup    `json:"group,omitempty"`    // 审批定义分组
	Instance *SearchApprovalTaskRespTaskInstance `json:"instance,omitempty"` // 审批实例信息
	Task     *SearchApprovalTaskRespTaskTask     `json:"task,omitempty"`     // 审批任务
}

// SearchApprovalTaskRespTaskApproval ...
type SearchApprovalTaskRespTaskApproval struct {
	Code       string                                      `json:"code,omitempty"`        // 审批定义 code
	Name       string                                      `json:"name,omitempty"`        // 审批定义名称
	IsExternal bool                                        `json:"is_external,omitempty"` // 是否为第三方审批
	External   *SearchApprovalTaskRespTaskApprovalExternal `json:"external,omitempty"`    // 第三方审批信息
	ApprovalID string                                      `json:"approval_id,omitempty"` // 审批定义Id
	Icon       string                                      `json:"icon,omitempty"`        // 审批定义图标信息
}

// SearchApprovalTaskRespTaskApprovalExternal ...
type SearchApprovalTaskRespTaskApprovalExternal struct {
	BatchCcRead bool `json:"batch_cc_read,omitempty"` // 是否支持批量读
}

// SearchApprovalTaskRespTaskGroup ...
type SearchApprovalTaskRespTaskGroup struct {
	ExternalID string `json:"external_id,omitempty"` // 审批定义分组外部 id
	Name       string `json:"name,omitempty"`        // 审批定义分组名称
}

// SearchApprovalTaskRespTaskInstance ...
type SearchApprovalTaskRespTaskInstance struct {
	Code       string                                  `json:"code,omitempty"`        // 审批实例 code
	ExternalID string                                  `json:"external_id,omitempty"` // 审批实例外部 id
	UserID     string                                  `json:"user_id,omitempty"`     // 审批实例发起人 id
	StartTime  string                                  `json:"start_time,omitempty"`  // 审批实例开始时间
	EndTime    string                                  `json:"end_time,omitempty"`    // 审批实例结束时间
	Status     string                                  `json:"status,omitempty"`      // 审批实例状态, 可选值有: reject: 拒绝, pending: 审批中, recall: 撤回, deleted: 已删除, approved: 通过
	Title      string                                  `json:"title,omitempty"`       // 审批实例名称（只有第三方审批有）
	Extra      string                                  `json:"extra,omitempty"`       // 审批实例扩展字段, string型json
	SerialID   string                                  `json:"serial_id,omitempty"`   // 审批流水号
	Link       *SearchApprovalTaskRespTaskInstanceLink `json:"link,omitempty"`        // 审批实例链接（只有第三方审批有）
}

// SearchApprovalTaskRespTaskInstanceLink ...
type SearchApprovalTaskRespTaskInstanceLink struct {
	PcLink     string `json:"pc_link,omitempty"`     // 审批实例 pc 端链接
	MobileLink string `json:"mobile_link,omitempty"` // 审批实例移动端链接
}

// SearchApprovalTaskRespTaskTask ...
type SearchApprovalTaskRespTaskTask struct {
	UserID         string                              `json:"user_id,omitempty"`          // 审批任务审批人 id
	StartTime      string                              `json:"start_time,omitempty"`       // 审批任务开始时间
	EndTime        string                              `json:"end_time,omitempty"`         // 审批任务结束时间
	Status         string                              `json:"status,omitempty"`           // 审批任务状态, 可选值有: rejected: 拒绝, pending: 审批中, approved: 通过, transferred: 转交, done: 已完成, rm_repeat: 去重, processed: 已处理, hidden: 隐藏
	Title          string                              `json:"title,omitempty"`            // 审批任务名称（只有第三方审批有）
	Extra          string                              `json:"extra,omitempty"`            // 审批任务扩展字段, string型json
	Link           *SearchApprovalTaskRespTaskTaskLink `json:"link,omitempty"`             // 审批任务链接（只有第三方审批有）
	TaskID         string                              `json:"task_id,omitempty"`          // 任务id
	UpdateTime     string                              `json:"update_time,omitempty"`      // 审批任务更新时间
	TaskExternalID string                              `json:"task_external_id,omitempty"` // 三方审批扩展 ID
}

// SearchApprovalTaskRespTaskTaskLink ...
type SearchApprovalTaskRespTaskTaskLink struct {
	PcLink     string `json:"pc_link,omitempty"`     // 审批实例 pc 端链接
	MobileLink string `json:"mobile_link,omitempty"` // 审批实例移动端链接
}

// searchApprovalTaskResp ...
type searchApprovalTaskResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *SearchApprovalTaskResp `json:"data,omitempty"`
}
