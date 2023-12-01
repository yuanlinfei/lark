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

// GetTaskList 以分页的方式获取任务列表。当使用user_access_token时, 获取与该用户身份相关的所有任务。当使用tenant_access_token时, 获取以该应用身份通过“创建任务“接口创建的所有任务（并非获取该应用所在租户下所有用户创建的任务）。
//
// 本接口支持通过任务创建时间以及任务的完成状态对任务进行过滤。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/list
// new doc: https://open.feishu.cn/document/server-docs/task-v1/task/list
func (r *TaskService) GetTaskList(ctx context.Context, request *GetTaskListReq, options ...MethodOptionFunc) (*GetTaskListResp, *Response, error) {
	if r.cli.mock.mockTaskGetTaskList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#GetTaskList mock enable")
		return r.cli.mock.mockTaskGetTaskList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "GetTaskList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getTaskListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskGetTaskList mock TaskGetTaskList method
func (r *Mock) MockTaskGetTaskList(f func(ctx context.Context, request *GetTaskListReq, options ...MethodOptionFunc) (*GetTaskListResp, *Response, error)) {
	r.mockTaskGetTaskList = f
}

// UnMockTaskGetTaskList un-mock TaskGetTaskList method
func (r *Mock) UnMockTaskGetTaskList() {
	r.mockTaskGetTaskList = nil
}

// GetTaskListReq ...
type GetTaskListReq struct {
	PageSize        *int64  `query:"page_size" json:"-"`         // 分页大小, 示例值: 10, 最大值: `100`
	PageToken       *string `query:"page_token" json:"-"`        // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: MTYzMTg3ODUxNQ[
	StartCreateTime *string `query:"start_create_time" json:"-"` // 范围查询任务时, 查询的起始时间。不填时默认起始时间为第一个任务的创建时间, 示例值: 1652323331
	EndCreateTime   *string `query:"end_create_time" json:"-"`   // 范围查询任务时, 查询的结束时间。不填时默认结束时间为最后一个任务的创建时间, 示例值: 1652323335
	TaskCompleted   *bool   `query:"task_completed" json:"-"`    // 可用于查询时过滤任务完成状态。true表示只返回已完成的任务, false表示只返回未完成的任务。不填时表示同时返回两种完成状态的任务, 示例值: false
	UserIDType      *IDType `query:"user_id_type" json:"-"`      // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetTaskListResp ...
type GetTaskListResp struct {
	Items     []*GetTaskListRespItem `json:"items,omitempty"`      // 任务列表
	PageToken string                 `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                   `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetTaskListRespItem ...
type GetTaskListRespItem struct {
	ID              string                             `json:"id,omitempty"`               // 任务的唯一ID, 例如"83912691-2e43-47fc-94a4-d512e03984fa"
	Summary         string                             `json:"summary,omitempty"`          // 任务的标题, 类型为文本字符串, 如果要在任务标题中插入 URL 或者 @某个用户, 请使用rich_summary字段, 创建任务时, 任务标题(summary字段)和任务富文本标题(rich_summary字段)不能同时为空, 需要至少填充其中一个字段, 任务标题和任务富文本标题同时存在时只使用富文本标题。
	Description     string                             `json:"description,omitempty"`      // 任务的描述, 类型为文本字符串, 如果要在任务描述中插入 URL 或者 @某个用户, 请使用rich_description字段, 任务备注和任务富文本备注同时存在时只使用富文本备注。
	CompleteTime    string                             `json:"complete_time,omitempty"`    // 任务的完成时间戳（单位为秒）, 完成时间为0则表示任务尚未完成, 不支持部分完成, 只有整个任务完成, 该字段才会有非0值。
	CreatorID       string                             `json:"creator_id,omitempty"`       // 任务的创建者 ID, 其中查询字段 user_id_type 是用于控制返回体中 creator_id 的类型, 不传时默认返回 open_id, 特别的, 使用tenant_access_token 调用接口时, 如果是user_id_type是openid, 则返回代表该应用身份的openid；当user_id_type为user_id时, 不返回creator_id。原因是user_id代表一个真实飞书用户的id, 应用身份没有user_id。使用user_access_token调用接口正常返回创建者。
	Extra           string                             `json:"extra,omitempty"`            // 附属信息, 接入方可以传入base64 编码后的自定义的数据。用户如果需要对当前任务备注信息, 但对外不显示, 可使用该字段进行存储, 该数据会在获取任务详情时, 原样返回给用户。
	CreateTime      string                             `json:"create_time,omitempty"`      // 任务的创建时间的Unix时间戳（单位为秒）
	UpdateTime      string                             `json:"update_time,omitempty"`      // 任务的更新时间的Unix时间戳（单位为秒）, 创建任务时update_time与create_time相同
	Due             *GetTaskListRespItemDue            `json:"due,omitempty"`              // 任务的截止时间设置
	Origin          *GetTaskListRespItemOrigin         `json:"origin,omitempty"`           // 任务关联的第三方平台来源信息
	CanEdit         bool                               `json:"can_edit,omitempty"`         // 此字段用于控制该任务在飞书任务中心是否可编辑, 默认为false, 已经废弃, 向前兼容故仍然保留, 但不推荐使用
	Custom          string                             `json:"custom,omitempty"`           // 自定义完成配置, 此字段用于设置完成任务时的页面跳转, 或展示提示语。详细参见: [任务字段补充说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/Supplementary-directions-of-task-fields)
	Source          int64                              `json:"source,omitempty"`           // 任务创建的来源, 可选值有: 0: 未知类型, 1: 来源任务中心创建, 2: 来源消息转任务, 3: 来源云文档, 4: 来源文档单品, 5: 来源PANO, 6: 来源tenant_access_token创建的任务, 7: 来源user_access_token创建的任务, 8: 来源新版云文档
	Followers       []*GetTaskListRespItemFollower     `json:"followers,omitempty"`        // 任务的关注者
	Collaborators   []*GetTaskListRespItemCollaborator `json:"collaborators,omitempty"`    // 任务的执行者
	CollaboratorIDs []string                           `json:"collaborator_ids,omitempty"` // 创建任务时添加的执行者用户id列表, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档: [如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get)。
	FollowerIDs     []string                           `json:"follower_ids,omitempty"`     // 创建任务时添加的关注者用户id列表, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档: [如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get)。
	RepeatRule      string                             `json:"repeat_rule,omitempty"`      // 重复任务的规则表达式, 语法格式参见[RRule语法规范](https://www.ietf.org/rfc/rfc2445.txt) 4.3.10小节
	RichSummary     string                             `json:"rich_summary,omitempty"`     // 富文本任务标题。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module), 。创建任务时, 任务标题(summary字段)和任务富文本标题(rich_summary字段)不能同时为空, 需要至少填充其中一个字段。
	RichDescription string                             `json:"rich_description,omitempty"` // 富文本任务备注。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module)
}

// GetTaskListRespItemCollaborator ...
type GetTaskListRespItemCollaborator struct {
	ID     string   `json:"id,omitempty"`      // 任务执行者的 ID, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档[如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 已经废弃, 为了向前兼容早期只支持单次添加一个人的情况而保留, 但不再推荐使用, 建议使用id_list字段
	IDList []string `json:"id_list,omitempty"` // 执行者的用户ID列表, 传入的值为 user_id 或 open_id, 由user_id_type 决定。user_id和open_id的获取可见文档[如何获取相关id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get)。
}

// GetTaskListRespItemDue ...
type GetTaskListRespItemDue struct {
	Time     string `json:"time,omitempty"`       // 表示截止时间的Unix时间戳（单位为秒）。
	Timezone string `json:"timezone,omitempty"`   // 截止时间对应的时区, 传入值需要符合IANA Time Zone Database标准, 规范见[Time Zone Database](https://www.iana.org/time-zones)。
	IsAllDay bool   `json:"is_all_day,omitempty"` // 标记任务是否为全天任务, 包括如下取值: true: 表示是全天任务, 全天任务的截止时间为当天 UTC 时间的 0 点, false: 表示不是全天任务。
}

// GetTaskListRespItemFollower ...
type GetTaskListRespItemFollower struct {
	ID     string   `json:"id,omitempty"`      // 任务关注人 ID
	IDList []string `json:"id_list,omitempty"` // 要删除的关注人ID列表
}

// GetTaskListRespItemOrigin ...
type GetTaskListRespItemOrigin struct {
	PlatformI18nName string                         `json:"platform_i18n_name,omitempty"` // 任务来源的名称, 用于在任务中心详情页展示。需要提供一个字典, 支持多种语言名称映射。应用在使用不同语言时, 导入来源也将展示对应的内容。详细参见: [任务字段补充说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/Supplementary-directions-of-task-fields)
	Href             *GetTaskListRespItemOriginHref `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

// GetTaskListRespItemOriginHref ...
type GetTaskListRespItemOriginHref struct {
	URL   string `json:"url,omitempty"`   // 具体链接地址, URL仅支持解析http、https。详细参见: [任务字段补充说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/Supplementary-directions-of-task-fields)
	Title string `json:"title,omitempty"` // 链接对应的标题
}

// getTaskListResp ...
type getTaskListResp struct {
	Code int64            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *GetTaskListResp `json:"data,omitempty"`
}
