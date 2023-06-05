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

// EventV2CorehrDepartmentCreatedV1 飞书人事中「部门被创建」时将触发此事件。注意: 1. 触发时间为部门实际生效时间, 如在 2022-01-01 创建部门, 部门生效时间设置为 2022-05-01, 事件将在 2022-05-01 进行推送。2. 现在创建部门也会同时触发「部门更新」事件{使用示例}(url=/api/tools/api_explore/api_explore_config?project=corehr&version=v1&resource=department&event=created)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/department/events/created
func (r *EventCallbackService) HandlerEventV2CorehrDepartmentCreatedV1(f EventV2CorehrDepartmentCreatedV1Handler) {
	r.cli.eventHandler.eventV2CorehrDepartmentCreatedV1Handler = f
}

// EventV2CorehrDepartmentCreatedV1Handler event EventV2CorehrDepartmentCreatedV1 handler
type EventV2CorehrDepartmentCreatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2CorehrDepartmentCreatedV1) (string, error)

// EventV2CorehrDepartmentCreatedV1 ...
type EventV2CorehrDepartmentCreatedV1 struct {
	DepartmentID string `json:"department_id,omitempty"` // 新建部门的 ID
}
