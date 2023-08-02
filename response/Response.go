package response

// Response 通用返回体
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

const (
	// SuccessCode 表示成功的状态码
	SuccessCode int32 = 0

	// ErrorCode 表示失败的状态码
	ErrorCode int32 = 1
)

// NewResponse 生成一个新的 Response 对象
func NewResponse(statusCode int32, statusMsg string) Response {
	return Response{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}

// Fail 失败返回 StatusCode=ErrorCode StatusMsg=""
func Fail() Response {
	return NewResponse(ErrorCode, "")
}

// Success 成功返回 StatusCode=SuccessCode StatusMsg=""
func Success() Response {
	return NewResponse(SuccessCode, "")
}
