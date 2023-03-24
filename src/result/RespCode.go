package result

var (
	OK              = response(0, "ok") // 通用成功
	ParamError      = response(1000001, "参数错误")
	Err             = response(1000002, "")           // 通用错误
	SystemError     = response(1000003, "系统异常，请稍候重试") //
	TokenInvalid    = response(1000004, "token失效")
	PermissionError = response(1000005, "权限不足")

	// UserNotFound 用户模块错误吗
	UserNotFound      = response(2000001, "用户不存在")
	UserPasswordError = response(2000002, "用户名或密码错误")
	UserLoginError    = response(2000003, "登录失败")

	// 节点
	NodeNotFound = response(3000001, "节点不存在")

	//钱包
	TransferError = response(4000001, "转账失败")

	// 服务商
	SpNotFound = response(5000001, "服务商不存在")
)
