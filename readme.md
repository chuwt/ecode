# 统一错误码

## 说明
根据`bilibili`的微服务框架[kratos](https://github.com/go-kratos/kratos)的[错误码](https://github.com/go-kratos/kratos/blob/master/doc/wiki-cn/ecode.md)修改而来，加入对信息的format功能

## quick start
1. 定义错误
```
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

	// 错误组，根据http状态码分类
	// 创建 http状态码为400的组
	group := NewGroup(400)
	// 创建错误，错误的http状态码都会是400
	BindErr := group.New(4000002, "参数解析错误")
	// 打印方法
	fmt.Println(BindErr.Code(), BindErr.HttpCode(), BindErr.Error())
```
2. 使用
```
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
```

