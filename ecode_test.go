package ecode

import (
	"errors"
	"fmt"
	"testing"
)

var (
	testErr = NewCode(4100008, 401, "你好%s%s")
)

func TestECode(t *testing.T) {
	var (
		code  Code
		err   error
		codes Codes
	)

	t.Log(OK.Code(), OK.HttpCode(), OK.Error())

	code = NewCode(4000001, 401, "你好%s%s")
	t.Log(code.Code(), code.HttpCode(), code.SetArgs("ok", "ok2").Error())

	group := NewGroup(400)
	code = group.New(4000002, "你好")
	t.Log(code.Code(), code.HttpCode(), code.Error())

	group = NewGroup(400)
	code = group.New(4000003, "你好")
	t.Log(code.Code(), code.HttpCode(), code.Error())

	// catch error and using Cause converting to Codes
	err = testErrReturn()
	codes = Cause(err)
	t.Log(codes.Code(), codes.HttpCode(), codes.Error())

	// if catch an unKnow err that doesn't register in _codes will returns UnDefinedErr
	// you can add log to get the err
	err = testUnKnowErr()
	codes = Cause(err)
	t.Log(codes.Code(), codes.HttpCode(), codes.Error())

	// test cause
	err = testErrMsg()
	codes = Cause(err)
	t.Log(codes.Code(), code.HttpCode(), codes.Error())

	t.Log(testErr.Code(), code.HttpCode(), testErr.SetArgs("1", "2").Error())

	// will panic due to the same code
	code = group.New(4000003, "你好")
	t.Log(code.Code(), code.HttpCode(), code.Error())
}

func testErrReturn() error {
	return ServerErr
}

func testErrMsg() error {
	return testErr.SetArgs("用户名", "密码错误")
}

func testUnKnowErr() error {
	// un set error redirect to ServerErr
	return errors.New("test error")
}

func TestOk(t *testing.T) {
	// 第一参数为自定义错误吗
	// 第二个参数为http状态码
	// 第三个错误为错误信息
	InvalidParamsErr := NewCode(4000001, 400, "参数错误")
	// 打印方法
	fmt.Println(InvalidParamsErr.Code(), InvalidParamsErr.HttpCode(), InvalidParamsErr.Error())

	// 错误信息可以format的错误
	UserNotFound := NewCode(4040001, 404, "用户%s不存在")
	// 打印方法
	fmt.Println(UserNotFound.Code(), UserNotFound.HttpCode(), UserNotFound.SetArgs("123").Error())

	// 推荐返回为error，更通用
	err := func() error {
		return InvalidParamsErr
	}()
	// Cause方法根据err 反推回自定义code
	// 如果未找到自定义code，则返回服务器错误
	code := Cause(err)
	fmt.Println(code.Code(), code.HttpCode(), code.Error())

	// format的错误信息
	err = func() error {
		return UserNotFound.SetArgs("123")
	}()
	code = Cause(err)
	fmt.Println(code.Code(), code.HttpCode(), code.Error())

	// format的错误信息
	err = func() error {
		return UserNotFound.SetArgs("123")
	}()
	code = Cause(err)
	fmt.Println(code.Code(), code.HttpCode(), code.Error())
}
