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

// GetVCAlertList 获取特定条件下租户的设备告警记录。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/alert/list
func (r *VCService) GetVCAlertList(ctx context.Context, request *GetVCAlertListReq, options ...MethodOptionFunc) (*GetVCAlertListResp, *Response, error) {
	if r.cli.mock.mockVCGetVCAlertList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#GetVCAlertList mock enable")
		return r.cli.mock.mockVCGetVCAlertList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCAlertList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/alerts",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getVCAlertListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCGetVCAlertList mock VCGetVCAlertList method
func (r *Mock) MockVCGetVCAlertList(f func(ctx context.Context, request *GetVCAlertListReq, options ...MethodOptionFunc) (*GetVCAlertListResp, *Response, error)) {
	r.mockVCGetVCAlertList = f
}

// UnMockVCGetVCAlertList un-mock VCGetVCAlertList method
func (r *Mock) UnMockVCGetVCAlertList() {
	r.mockVCGetVCAlertList = nil
}

// GetVCAlertListReq ...
type GetVCAlertListReq struct {
	StartTime  string  `query:"start_time" json:"-"`  // 开始时间（unix时间, 单位秒）, 示例值: "1608888867"
	EndTime    string  `query:"end_time" json:"-"`    // 结束时间（unix时间, 单位秒）, 示例值: "1608888867"
	QueryType  *int64  `query:"query_type" json:"-"`  // 查询对象类型, 不填返回所有, 示例值: 1, 可选值有: 1: 会议室, 2: 企业会议室连接器
	QueryValue *string `query:"query_value" json:"-"` // 查询对象ID, 会议室ID或企业会议室连接器ID, 示例值: "6911188411932033028"
	PageSize   *int64  `query:"page_size" json:"-"`   // 分页大小, 示例值: 100, 最大值: `200`
	PageToken  *string `query:"page_token" json:"-"`  // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "100"
}

// GetVCAlertListResp ...
type GetVCAlertListResp struct {
	HasMore   bool                      `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                    `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetVCAlertListRespItem `json:"items,omitempty"`      // 告警记录
}

// GetVCAlertListRespItem ...
type GetVCAlertListRespItem struct {
	AlertID       string                           `json:"alert_id,omitempty"`       // 告警ID
	ResourceScope string                           `json:"resource_scope,omitempty"` // 触发告警规则的会议室/服务器具体的名称
	MonitorTarget int64                            `json:"monitor_target,omitempty"` // 触发告警规则的监控对象, 可选值有: 1: 飞书会议室, 2: 飞书会议室签到板, 3: 飞书投屏盒子, 4: 飞书投屏, 5: sip会议室系统, 6: erc节点, 7: 飞书传感器
	AlertStrategy string                           `json:"alert_strategy,omitempty"` // 告警规则的规则描述
	AlertTime     string                           `json:"alert_time,omitempty"`     // 告警通知发生时间（unix时间, 单位秒）
	AlertLevel    int64                            `json:"alert_level,omitempty"`    // 告警等级: 严重/警告/提醒, 可选值有: 0: 提醒, 1: 警告, 2: 严重
	Contacts      []*GetVCAlertListRespItemContact `json:"contacts,omitempty"`       // 告警联系人
	NotifyMethods []int64                          `json:"notifyMethods,omitempty"`  // 通知方式, 可选值有: 0: 飞书机器人, 1: 邮件
	AlertRule     string                           `json:"alertRule,omitempty"`      // 规则名称
}

// GetVCAlertListRespItemContact ...
type GetVCAlertListRespItemContact struct {
	ContactType int64  `json:"contact_type,omitempty"` // 联系人类型, 可选值有: 1: 用户, 2: 用户组, 3: 部门
	ContactName string `json:"contact_name,omitempty"` // 联系人名
}

// getVCAlertListResp ...
type getVCAlertListResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetVCAlertListResp `json:"data,omitempty"`
}
