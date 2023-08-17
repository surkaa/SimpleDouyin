// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    response, err := UnmarshalResponse(bytes)
//    bytes, err = response.Marshal()

package response

import "encoding/json"

func (r *Response) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Response struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

const (
	// SuccessCode 表示成功的状态码
	SuccessCode int64 = 0

	// ErrorCode 表示失败的状态码
	ErrorCode int64 = 1
)

// NewBaseResponse 生成一个新的 Response 对象
func NewBaseResponse(statusCode int64, statusMsg string) Response {
	return Response{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}

// Fail 失败返回 StatusCode= ErrorCode(1) StatusMsg=""
func Fail() Response {
	return NewBaseResponse(ErrorCode, "")
}

// FailWithMsg 失败返回 StatusCode=ErrorCode StatusMsg=msg
func FailWithMsg(msg string) Response {
	return NewBaseResponse(ErrorCode, msg)
}

// Success 成功返回 StatusCode=SuccessCode StatusMsg=""
func Success() Response {
	return NewBaseResponse(SuccessCode, "")
}
