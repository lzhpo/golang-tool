package response_vo

// Response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseFunc func(*Response)

// WithCode
func WithCode(code int) ResponseFunc {
	return func(response *Response) {
		response.Code = code
	}
}

// WithMsg
func WithMsg(msg string) ResponseFunc {
	return func(response *Response) {
		response.Msg = msg
	}
}

// WithData
func WithData(data interface{}) ResponseFunc {
	return func(response *Response) {
		response.Data = data
	}
}

// default Response
var (
	response = Response{
		Code: 200,
		Msg:  "success",
		Data: "",
	}
)

func NewResponse1(opts ...ResponseFunc) *Response {
	temp := response

	for _, opt := range opts {
		opt(&temp)
	}

	return &temp
}

func NewResponse2(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
