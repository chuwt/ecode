package ecode

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

var (
	_codes = map[int]int{} // register codes. stored code:httpStatusCode
)

// Codes ecode error interface which has a code & message.
type Codes interface {
	// Error get error msg.
	Error() string
	// Code get error code.
	Code() int
	// Get httpStatus code.
	HttpCode() int
}

type Code struct {
	httpStatusCode int
	errCode        int
	msg            string
}

func (e Code) Error() string {
	return e.msg
}

func (e Code) HttpCode() int { return e.httpStatusCode }

// Code return error code
func (e Code) Code() int { return e.errCode }

func (e Code) SetArgs(args ...interface{}) Code {
	copyCode := CopyCode(e)
	msg := fmt.Sprintf(e.msg, args...)
	copyCode.msg = msg
	return copyCode
}

type Group struct {
	httpStatusCode int
}

func (g *Group) New(e int, msg ...string) Code {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return NewCode(e, g.httpStatusCode, msg...)
}

func CopyCode(code Code) Code {
	return Code{
		errCode:        code.errCode,
		httpStatusCode: code.httpStatusCode,
	}
}

// return a group that httpStatusCode is httpStatusCode
func NewGroup(httpStatusCode int) Group {
	return Group{
		httpStatusCode: httpStatusCode,
	}
}

// return a single eCode
func NewCode(e, httpStatusCode int, msg ...string) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = httpStatusCode
	return Code{
		errCode:        e,
		httpStatusCode: httpStatusCode,
		msg:            strings.Join(msg, ""),
	}
}

// Cause cause from error to ecode.
func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return UnDefinedErr.SetArgs(e.Error())
}
