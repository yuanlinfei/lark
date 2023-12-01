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

// UpdateHelpdeskAgentSkill 该接口用于更新客服技能。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/patch
// new doc: https://open.feishu.cn/document/server-docs/helpdesk-v1/agent-function/agent_skill/patch
func (r *HelpdeskService) UpdateHelpdeskAgentSkill(ctx context.Context, request *UpdateHelpdeskAgentSkillReq, options ...MethodOptionFunc) (*UpdateHelpdeskAgentSkillResp, *Response, error) {
	if r.cli.mock.mockHelpdeskUpdateHelpdeskAgentSkill != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#UpdateHelpdeskAgentSkill mock enable")
		return r.cli.mock.mockHelpdeskUpdateHelpdeskAgentSkill(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "UpdateHelpdeskAgentSkill",
		Method:              "PATCH",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/agent_skills/:agent_skill_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(updateHelpdeskAgentSkillResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskUpdateHelpdeskAgentSkill mock HelpdeskUpdateHelpdeskAgentSkill method
func (r *Mock) MockHelpdeskUpdateHelpdeskAgentSkill(f func(ctx context.Context, request *UpdateHelpdeskAgentSkillReq, options ...MethodOptionFunc) (*UpdateHelpdeskAgentSkillResp, *Response, error)) {
	r.mockHelpdeskUpdateHelpdeskAgentSkill = f
}

// UnMockHelpdeskUpdateHelpdeskAgentSkill un-mock HelpdeskUpdateHelpdeskAgentSkill method
func (r *Mock) UnMockHelpdeskUpdateHelpdeskAgentSkill() {
	r.mockHelpdeskUpdateHelpdeskAgentSkill = nil
}

// UpdateHelpdeskAgentSkillReq ...
type UpdateHelpdeskAgentSkillReq struct {
	AgentSkillID string                                 `path:"agent_skill_id" json:"-"` // agent skill id, 示例值: "test-skill-id"
	AgentSkill   *UpdateHelpdeskAgentSkillReqAgentSkill `json:"agent_skill,omitempty"`   // 更新技能
}

// UpdateHelpdeskAgentSkillReqAgentSkill ...
type UpdateHelpdeskAgentSkillReqAgentSkill struct {
	Name     *string                                      `json:"name,omitempty"`      // 技能名, 示例值: "skill-name"
	Rules    []*UpdateHelpdeskAgentSkillReqAgentSkillRule `json:"rules,omitempty"`     // 技能rules
	AgentIDs []string                                     `json:"agent_ids,omitempty"` // 具有此技能的客服ids, 示例值: ["ou_ea21d7f018e1155d960e40d33191f966"]
}

// UpdateHelpdeskAgentSkillReqAgentSkillRule ...
type UpdateHelpdeskAgentSkillReqAgentSkillRule struct {
	ID               *string `json:"id,omitempty"`                // rule id, 参考[获取客服技能rules](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill_rule/list) 用于获取rules options, 示例值: "test-skill-id"
	SelectedOperator *int64  `json:"selected_operator,omitempty"` // 运算符比较, 参考[客服技能运算符选项](https://open.feishu.cn/document/ukTMukTMukTM/ucDOyYjL3gjM24yN4IjN/operator-options), 示例值: 8
	OperatorOptions  []int64 `json:"operator_options,omitempty"`  // rule操作数value, [客服技能及运算符](https://open.feishu.cn/document/ukTMukTMukTM/ucDOyYjL3gjM24yN4IjN/operator-options), 示例值: [7, 8]
	Operand          *string `json:"operand,omitempty"`           // rule 操作数的值, 示例值: "{\"selected_departments\":[{\"id\":\"部门ID\", \"name\":\"IT\"}]}"
}

// UpdateHelpdeskAgentSkillResp ...
type UpdateHelpdeskAgentSkillResp struct {
}

// updateHelpdeskAgentSkillResp ...
type updateHelpdeskAgentSkillResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *UpdateHelpdeskAgentSkillResp `json:"data,omitempty"`
}
