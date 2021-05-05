package lark

import (
	"context"
)

// GetCalendarFreeBusyList 查询用户主日历或会议室的忙闲信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/freebusy/list
func (r *CalendarAPI) GetCalendarFreeBusyList(ctx context.Context, request *GetCalendarFreeBusyListReq) (*GetCalendarFreeBusyListResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/freebusy/list",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getCalendarFreeBusyListResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "GetCalendarFreeBusyList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetCalendarFreeBusyListReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型,**示例值**："open_id",**可选值有**：,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值**：`open_id`,**当值为 `user_id`，字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	TimeMin    string  `json:"time_min,omitempty"`     // 查询时段开始时间,**示例值**："2020-10-28T12:00:00+08:00"
	TimeMax    string  `json:"time_max,omitempty"`     // 查询时段结束时间,**示例值**："2020-12-28T12:00:00+08:00"
	UserID     *string `json:"user_id,omitempty"`      // 用户user_id，输入时与 room_id 二选一,**示例值**："ou_xxxxxxxxxx"
	RoomID     *string `json:"room_id,omitempty"`      // 会议室room_id，输入时与 user_id 二选一,**示例值**："omm_xxxxxxxxxx"
}

type getCalendarFreeBusyListResp struct {
	Code int                          `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarFreeBusyListResp `json:"data,omitempty"` // \-
}

type GetCalendarFreeBusyListResp struct {
	FreebusyList []*GetCalendarFreeBusyListRespFreebusy `json:"freebusy_list,omitempty"` // 日历上请求时间区间内的忙闲信息
}

type GetCalendarFreeBusyListRespFreebusy struct {
	StartTime string `json:"start_time,omitempty"` // 忙闲信息开始时间，RFC3339 date_time 格式
	EndTime   string `json:"end_time,omitempty"`   // 忙闲信息结束时间，RFC3339 date_time 格式
}