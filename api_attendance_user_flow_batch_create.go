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

// BatchCreateAttendanceUserFlow 导入授权内员工的打卡流水记录。导入后, 会根据员工所在的考勤组班次规则, 计算最终的打卡状态与结果。
//
// 适用于考勤机数据导入等场景。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_flow/batch_create
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/user_task/batch_create
func (r *AttendanceService) BatchCreateAttendanceUserFlow(ctx context.Context, request *BatchCreateAttendanceUserFlowReq, options ...MethodOptionFunc) (*BatchCreateAttendanceUserFlowResp, *Response, error) {
	if r.cli.mock.mockAttendanceBatchCreateAttendanceUserFlow != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#BatchCreateAttendanceUserFlow mock enable")
		return r.cli.mock.mockAttendanceBatchCreateAttendanceUserFlow(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "BatchCreateAttendanceUserFlow",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_flows/batch_create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchCreateAttendanceUserFlowResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceBatchCreateAttendanceUserFlow mock AttendanceBatchCreateAttendanceUserFlow method
func (r *Mock) MockAttendanceBatchCreateAttendanceUserFlow(f func(ctx context.Context, request *BatchCreateAttendanceUserFlowReq, options ...MethodOptionFunc) (*BatchCreateAttendanceUserFlowResp, *Response, error)) {
	r.mockAttendanceBatchCreateAttendanceUserFlow = f
}

// UnMockAttendanceBatchCreateAttendanceUserFlow un-mock AttendanceBatchCreateAttendanceUserFlow method
func (r *Mock) UnMockAttendanceBatchCreateAttendanceUserFlow() {
	r.mockAttendanceBatchCreateAttendanceUserFlow = nil
}

// BatchCreateAttendanceUserFlowReq ...
type BatchCreateAttendanceUserFlowReq struct {
	EmployeeType EmployeeType                                  `query:"employee_type" json:"-"` // 请求体和响应体中的 user_id 和 creator_id 的员工工号类型, 示例值: employee_id, 可选值有: employee_id: 员工 employee ID, 即[飞书管理后台](https://bytedance.feishu.cn/admin/contacts/departmentanduser) > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即[飞书管理后台](https://bytedance.feishu.cn/admin/contacts/departmentanduser) > 组织架构 > 成员与部门 > 成员详情中的工号
	FlowRecords  []*BatchCreateAttendanceUserFlowReqFlowRecord `json:"flow_records,omitempty"`  // 打卡流水记录列表(数量限制50)
}

// BatchCreateAttendanceUserFlowReqFlowRecord ...
type BatchCreateAttendanceUserFlowReqFlowRecord struct {
	UserID       string   `json:"user_id,omitempty"`       // 用户 ID, 示例值: "abd754f7"
	CreatorID    string   `json:"creator_id,omitempty"`    // 记录创建者 ID, 示例值: "abd754f7"
	LocationName string   `json:"location_name,omitempty"` // 打卡位置名称信息, 示例值: "西溪八方城"
	CheckTime    string   `json:"check_time,omitempty"`    // 打卡时间, 精确到秒的时间戳, 示例值: "1611476284"
	Comment      string   `json:"comment,omitempty"`       // 打卡备注, 示例值: "上班打卡"
	RecordID     *string  `json:"record_id,omitempty"`     // 打卡记录 ID, 示例值: "6709359313699356941"
	Ssid         *string  `json:"ssid,omitempty"`          // 打卡 Wi-Fi 的 SSID, 示例值: "b0:b8:67:5c:1d:72"
	Bssid        *string  `json:"bssid,omitempty"`         // 打卡 Wi-Fi 的 MAC 地址, 示例值: "b0:b8:67:5c:1d:72"
	IsField      *bool    `json:"is_field,omitempty"`      // 是否为外勤打卡, 示例值: true
	IsWifi       *bool    `json:"is_wifi,omitempty"`       // 是否为 Wi-Fi 打卡, 示例值: true
	Type         *int64   `json:"type,omitempty"`          // 记录生成方式, 示例值: 在开放平台调用时, 此参数无效, 内部值始终是7, 可选值有: 0: 用户打卡, 1: 管理员修改, 2: 用户补卡, 3: 系统自动生成, 4: 下班免打卡, 5: 考勤机, 6: 极速打卡, 7: 考勤开放平台导入
	PhotoURLs    []string `json:"photo_urls,omitempty"`    // 打卡照片列表（该字段目前不支持）, 示例值: ["https://time.clockin.biz/manage/download/6840389754748502021"]
	CheckResult  *string  `json:"check_result,omitempty"`  // 打卡结果, 示例值: "Invalid", 可选值有: NoNeedCheck: 无需打卡, SystemCheck: 系统打卡, Normal: 正常, Early: 早退, Late: 迟到, SeriousLate: 严重迟到, Lack: 缺卡, Invalid: 无效, None: 无状态, Todo: 尚未打卡
}

// BatchCreateAttendanceUserFlowResp ...
type BatchCreateAttendanceUserFlowResp struct {
	FlowRecords []*BatchCreateAttendanceUserFlowRespFlowRecord `json:"flow_records,omitempty"` // 打卡流水记录列表
}

// BatchCreateAttendanceUserFlowRespFlowRecord ...
type BatchCreateAttendanceUserFlowRespFlowRecord struct {
	UserID       string   `json:"user_id,omitempty"`       // 用户 ID
	CreatorID    string   `json:"creator_id,omitempty"`    // 记录创建者 ID
	LocationName string   `json:"location_name,omitempty"` // 打卡位置名称信息
	CheckTime    string   `json:"check_time,omitempty"`    // 打卡时间, 精确到秒的时间戳
	Comment      string   `json:"comment,omitempty"`       // 打卡备注
	RecordID     string   `json:"record_id,omitempty"`     // 打卡记录 ID
	Ssid         string   `json:"ssid,omitempty"`          // 打卡 Wi-Fi 的 SSID
	Bssid        string   `json:"bssid,omitempty"`         // 打卡 Wi-Fi 的 MAC 地址
	IsField      bool     `json:"is_field,omitempty"`      // 是否为外勤打卡
	IsWifi       bool     `json:"is_wifi,omitempty"`       // 是否为 Wi-Fi 打卡
	Type         int64    `json:"type,omitempty"`          // 记录生成方式, 可选值有: 0: 用户打卡, 1: 管理员修改, 2: 用户补卡, 3: 系统自动生成, 4: 下班免打卡, 5: 考勤机, 6: 极速打卡, 7: 考勤开放平台导入
	PhotoURLs    []string `json:"photo_urls,omitempty"`    // 打卡照片列表
	CheckResult  string   `json:"check_result,omitempty"`  // 打卡结果, 可选值有: NoNeedCheck: 无需打卡, SystemCheck: 系统打卡, Normal: 正常, Early: 早退, Late: 迟到, SeriousLate: 严重迟到, Lack: 缺卡, Invalid: 无效, None: 无状态, Todo: 尚未打卡
}

// batchCreateAttendanceUserFlowResp ...
type batchCreateAttendanceUserFlowResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *BatchCreateAttendanceUserFlowResp `json:"data,omitempty"`
}
