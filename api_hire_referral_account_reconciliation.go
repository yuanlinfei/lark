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

// ReconcileHireReferralAccount 定时将时间段内的账户充值信息同步到招聘, 与招聘实际提取金额做对比, 保证系统异常或其他意外情况发生时, 双方系统可及时监控到充值异常等错误
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/referral_account/reconciliation
func (r *HireService) ReconcileHireReferralAccount(ctx context.Context, request *ReconcileHireReferralAccountReq, options ...MethodOptionFunc) (*ReconcileHireReferralAccountResp, *Response, error) {
	if r.cli.mock.mockHireReconcileHireReferralAccount != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#ReconcileHireReferralAccount mock enable")
		return r.cli.mock.mockHireReconcileHireReferralAccount(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "ReconcileHireReferralAccount",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/referral_account/reconciliation",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(reconcileHireReferralAccountResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireReconcileHireReferralAccount mock HireReconcileHireReferralAccount method
func (r *Mock) MockHireReconcileHireReferralAccount(f func(ctx context.Context, request *ReconcileHireReferralAccountReq, options ...MethodOptionFunc) (*ReconcileHireReferralAccountResp, *Response, error)) {
	r.mockHireReconcileHireReferralAccount = f
}

// UnMockHireReconcileHireReferralAccount un-mock HireReconcileHireReferralAccount method
func (r *Mock) UnMockHireReconcileHireReferralAccount() {
	r.mockHireReconcileHireReferralAccount = nil
}

// ReconcileHireReferralAccountReq ...
type ReconcileHireReferralAccountReq struct {
	StartTransTime *string                                       `json:"start_trans_time,omitempty"` // 按时间范围进行对账时 时间段的起始交易时间, 示例值: "1685416831621"
	EndTransTime   *string                                       `json:"end_trans_time,omitempty"`   // 按时间范围进行对账时 时间段的截止交易时间, 示例值: "1685416831622"
	TradeDetails   []*ReconcileHireReferralAccountReqTradeDetail `json:"trade_details,omitempty"`    // 交易信息
}

// ReconcileHireReferralAccountReqTradeDetail ...
type ReconcileHireReferralAccountReqTradeDetail struct {
	AccountID               string `json:"account_id,omitempty"`                 // 账户ID, 示例值: "6930815272790114324"
	TotalRechargeRewardInfo *int64 `json:"total_recharge_reward_info,omitempty"` // 时间段内该账户在积分商城的实际充值金额
}

// ReconcileHireReferralAccountResp ...
type ReconcileHireReferralAccountResp struct {
	CheckFailedList []*ReconcileHireReferralAccountRespCheckFailed `json:"check_failed_list,omitempty"` // 核对失败的信息
}

// ReconcileHireReferralAccountRespCheckFailed ...
type ReconcileHireReferralAccountRespCheckFailed struct {
	AccountID               string `json:"account_id,omitempty"`                 // 账户ID
	TotalWithdrawRewardInfo int64  `json:"total_withdraw_reward_info,omitempty"` // 招聘系统内的提取金额
	TotalRechargeRewardInfo int64  `json:"total_recharge_reward_info,omitempty"` // 商城实际充值金额
}

// reconcileHireReferralAccountResp ...
type reconcileHireReferralAccountResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *ReconcileHireReferralAccountResp `json:"data,omitempty"`
}
