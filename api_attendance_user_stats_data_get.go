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

// GetAttendanceUserStatsData 查询日度统计或月度统计的统计数据。
//
// * 调用统计开放接口api目前不返回休假加班新增字段类型
// * 当天不在考勤组时没有相关统计数据
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_data/query
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/user_stats_data/query-3
func (r *AttendanceService) GetAttendanceUserStatsData(ctx context.Context, request *GetAttendanceUserStatsDataReq, options ...MethodOptionFunc) (*GetAttendanceUserStatsDataResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceUserStatsData != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceUserStatsData mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceUserStatsData(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceUserStatsData",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_stats_datas/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceUserStatsDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceGetAttendanceUserStatsData mock AttendanceGetAttendanceUserStatsData method
func (r *Mock) MockAttendanceGetAttendanceUserStatsData(f func(ctx context.Context, request *GetAttendanceUserStatsDataReq, options ...MethodOptionFunc) (*GetAttendanceUserStatsDataResp, *Response, error)) {
	r.mockAttendanceGetAttendanceUserStatsData = f
}

// UnMockAttendanceGetAttendanceUserStatsData un-mock AttendanceGetAttendanceUserStatsData method
func (r *Mock) UnMockAttendanceGetAttendanceUserStatsData() {
	r.mockAttendanceGetAttendanceUserStatsData = nil
}

// GetAttendanceUserStatsDataReq ...
type GetAttendanceUserStatsDataReq struct {
	EmployeeType     EmployeeType `query:"employee_type" json:"-"`      // 请求体中的 user_ids 和响应体中的 user_id 的员工工号类型, 示例值: employee_id, 可选值有: employee_id: 员工 employee ID, 即[飞书管理后台](https://bytedance.feishu.cn/admin/contacts/departmentanduser) > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即[飞书管理后台](https://bytedance.feishu.cn/admin/contacts/departmentanduser) > 组织架构 > 成员与部门 > 成员详情中的工号
	Locale           string       `json:"locale,omitempty"`             // 语言类型, 示例值: "zh", 可选值有: en: 英语, ja: 日语, zh: 中文
	StatsType        string       `json:"stats_type,omitempty"`         // 统计类型, 示例值: "month", 可选值有: daily: 日度统计, month: 月度统计
	StartDate        int64        `json:"start_date,omitempty"`         // 开始时间, 示例值: 20210316
	EndDate          int64        `json:"end_date,omitempty"`           // 结束时间, （时间间隔不超过 31 天）, 示例值: 20210323
	UserIDs          []string     `json:"user_ids,omitempty"`           // 查询的用户 ID 列表, （用户数量不超过 200）, * 必填字段(已全部升级到新系统, 新系统要求必填), 示例值: ["ec8ddg56", "af5ddg73"]
	NeedHistory      *bool        `json:"need_history,omitempty"`       // 是否需要历史数据, 示例值: true
	CurrentGroupOnly *bool        `json:"current_group_only,omitempty"` // 只展示当前考勤组, 示例值: true
	UserID           *string      `json:"user_id,omitempty"`            // 操作者的 user_id, * 不同的操作者（管理员）的每个报表可能有不同的字段设置, 系统将根据 user_id 查询指定报表的统计数据, * 必填字段（已全部升级到新系统, 新系统要求该字段必填）, 示例值: "ec8ddg56"
}

// GetAttendanceUserStatsDataResp ...
type GetAttendanceUserStatsDataResp struct {
	UserDatas       []*GetAttendanceUserStatsDataRespUserData `json:"user_datas,omitempty"`        // 用户统计数据（限制1000条, 超过1000条会截断）
	InvalidUserList []string                                  `json:"invalid_user_list,omitempty"` // 无权限获取的用户列表
}

// GetAttendanceUserStatsDataRespUserData ...
type GetAttendanceUserStatsDataRespUserData struct {
	Name   string                                        `json:"name,omitempty"`    // 用户姓名
	UserID string                                        `json:"user_id,omitempty"` // 用户 ID
	Datas  []*GetAttendanceUserStatsDataRespUserDataData `json:"datas,omitempty"`   // 用户的统计数据
}

// GetAttendanceUserStatsDataRespUserDataData ...
type GetAttendanceUserStatsDataRespUserDataData struct {
	Code     string                                               `json:"code,omitempty"`     // 字段编号
	Value    string                                               `json:"value,omitempty"`    // 数据值
	Features []*GetAttendanceUserStatsDataRespUserDataDataFeature `json:"features,omitempty"` // 数据属性
	Title    string                                               `json:"title,omitempty"`    // 字段标题
}

// GetAttendanceUserStatsDataRespUserDataDataFeature ...
type GetAttendanceUserStatsDataRespUserDataDataFeature struct {
	Key   string `json:"key,omitempty"`   // 统计数据列附加属性的名称
	Value string `json:"value,omitempty"` // 统计数据列附加属性的值, * 先展示上下班的打卡结果, 再展示假勤申请时间(如果有)
}

// getAttendanceUserStatsDataResp ...
type getAttendanceUserStatsDataResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceUserStatsDataResp `json:"data,omitempty"`
}
