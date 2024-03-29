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
	"io"
)

// UploadFaceVerifyImage 无源人脸比对流程, 开发者后台通过调用此接口将基准图片上传到飞书后台, 做检测时的对比使用。
//
// ::: note
// 无源人脸比对接入需申请白名单, 接入前请联系飞书开放平台工作人员, 邮箱: open-platform@bytedance.com。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/upload-facial-reference-image
// new doc: https://open.feishu.cn/document/server-docs/human_authentication-v1/upload-facial-reference-image
func (r *HumanAuthService) UploadFaceVerifyImage(ctx context.Context, request *UploadFaceVerifyImageReq, options ...MethodOptionFunc) (*UploadFaceVerifyImageResp, *Response, error) {
	if r.cli.mock.mockHumanAuthUploadFaceVerifyImage != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] HumanAuth#UploadFaceVerifyImage mock enable")
		return r.cli.mock.mockHumanAuthUploadFaceVerifyImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "HumanAuth",
		API:                   "UploadFaceVerifyImage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/face_verify/v1/upload_face_image",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		IsFile:                true,
	}
	resp := new(uploadFaceVerifyImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHumanAuthUploadFaceVerifyImage mock HumanAuthUploadFaceVerifyImage method
func (r *Mock) MockHumanAuthUploadFaceVerifyImage(f func(ctx context.Context, request *UploadFaceVerifyImageReq, options ...MethodOptionFunc) (*UploadFaceVerifyImageResp, *Response, error)) {
	r.mockHumanAuthUploadFaceVerifyImage = f
}

// UnMockHumanAuthUploadFaceVerifyImage un-mock HumanAuthUploadFaceVerifyImage method
func (r *Mock) UnMockHumanAuthUploadFaceVerifyImage() {
	r.mockHumanAuthUploadFaceVerifyImage = nil
}

// UploadFaceVerifyImageReq ...
type UploadFaceVerifyImageReq struct {
	OpenID     *string   `query:"open_id" json:"-"`     // 用户应用标识, 与employee_id二选其一
	EmployeeID *string   `query:"employee_id" json:"-"` // 用户租户标识, 与open_id二选其一
	Image      io.Reader `json:"image,omitempty"`       // 带有头像的人脸照片
}

// UploadFaceVerifyImageResp ...
type UploadFaceVerifyImageResp struct {
	FaceUid string `json:"face_uid,omitempty"` // 人脸图片用户Uid, 需返回给应用小程序, 作为小程序调起人脸识别接口的uid参数
}

// uploadFaceVerifyImageResp ...
type uploadFaceVerifyImageResp struct {
	Code  int64                      `json:"code,omitempty"` // 返回码, 非0为失败
	Msg   string                     `json:"msg,omitempty"`  // 返回信息, 返回码的描述
	Data  *UploadFaceVerifyImageResp `json:"data,omitempty"` // 业务数据
	Error *ErrorDetail               `json:"error,omitempty"`
}
