package code

import (
	"fmt"
	"net/http"
	"sync"
)

const (
	Success           = 200
	BadRequest        = 400
	Unauthorized      = 401
	Forbidden         = 403
	NotFound          = 404
	Internal          = 500
	LoginFailed       = 1001
	LoginUserNotExist = 1002
	ChatWsNewErr      = 2001
	RoomUserJoinErr   = 3001
	RoomCreateErr     = 3002
)

// 定义常用的状态码和，错误信息，可以根据需要自行添加
func init() {
	register(Success, 200, "success")
	register(BadRequest, 400, "bad request")
	register(Unauthorized, 401, "unauthorized")
	register(Forbidden, 403, "forbidden")
	register(NotFound, 404, "not found")
	register(Internal, 500, "internal server error")
	register(LoginFailed, 200, "login failed")
	register(LoginUserNotExist, 200, "username or password error")
	register(ChatWsNewErr, 200, "chat ws new error")
	register(RoomUserJoinErr, 200, "room user join error")
	register(RoomCreateErr, 200, "room create error")
}

var (
	codes   = map[int]Coder{}
	codeMux = &sync.Mutex{}
)

func register(code int, httpStatus int, message string) {
	coder := &ErrCode{
		C:          code,
		HttpStatus: httpStatus,
		Message:    message,
	}

	MustRegister(coder)
}

// MustRegister 注册错误码
func MustRegister(coder Coder) {
	if coder.Code() == 0 {
		panic("ErrUnknown error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}

	codes[coder.Code()] = coder
}

func GetCoder(code int) Coder {
	return codes[code]
}

// ParseCoder 解析错误状态码
func ParseCoder(err error) Coder {
	if err == nil {
		return nil
	}

	if v, ok := err.(*withCode); ok {
		if coder, ok := codes[v.code]; ok {
			msg := v.Error() // 如果有错误信息，则返回错误信息
			if msg != "" {
				return ErrCode{
					C:          coder.Code(),
					HttpStatus: coder.HTTPStatus(),
					Message:    msg,
				}
			}
			return coder
		}
	}

	return ErrCode{
		C:          http.StatusBadRequest,
		HttpStatus: http.StatusBadRequest,
		Message:    err.Error(),
	}
}

// Coder 返回状态码接口定义
type Coder interface {
	Code() int
	HTTPStatus() int
	String() string
}

type ErrCode struct {
	C          int
	HttpStatus int
	Message    string
}

func (coder ErrCode) Code() int {
	return coder.C
}

func (coder ErrCode) String() string {
	return coder.Message
}

func (coder ErrCode) HTTPStatus() int {
	if coder.HttpStatus == 0 {
		return http.StatusInternalServerError
	}

	return coder.HttpStatus
}

type withCode struct {
	err  string
	code int
}

func (w *withCode) Error() string {
	if w.err == "" {
		return ""
	}
	return w.err
}

func WithCodeMsg(code int, msg ...any) error {
	var errMsg string
	switch len(msg) {
	case 0:
	case 1:
		errMsg = msg[0].(string)
	default:
		errMsg = fmt.Sprintf(msg[0].(string), msg[1:]...)
	}
	return &withCode{
		err:  errMsg,
		code: code,
	}
}
