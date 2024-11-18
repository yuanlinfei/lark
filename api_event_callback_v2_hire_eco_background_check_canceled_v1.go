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

// EventV2HireEcoBackgroundCheckCanceledV1 用户在招聘系统终止背调后, 系统会推送事件给对应的应用开发者。开发者可根据事件获取背调 ID, 完成在三方服务商处的订单取消等后续操作。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=hire&version=v1&resource=eco_background_check&event=canceled)
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/eco_background_check/events/canceled
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/ecological-docking/eco_background_check/events/canceled
func (r *EventCallbackService) HandlerEventV2HireEcoBackgroundCheckCanceledV1(f EventV2HireEcoBackgroundCheckCanceledV1Handler) {
	r.cli.eventHandler.eventV2HireEcoBackgroundCheckCanceledV1Handler = f
}

// EventV2HireEcoBackgroundCheckCanceledV1Handler event EventV2HireEcoBackgroundCheckCanceledV1 handler
type EventV2HireEcoBackgroundCheckCanceledV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2HireEcoBackgroundCheckCanceledV1) (string, error)

// EventV2HireEcoBackgroundCheckCanceledV1 ...
type EventV2HireEcoBackgroundCheckCanceledV1 struct {
	BackgroundCheckID string `json:"background_check_id,omitempty"` // 背调 ID, 招聘系统内唯一
	TerminationReason string `json:"termination_reason,omitempty"`  // 终止原因
}