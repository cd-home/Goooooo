package errno

import "errors"

// System Error
var (
	ErrorNotFoundService    = errors.New("未找到服务应用与环境信息")
	ErrorDataBase           = errors.New("数据库操作失败, 请重试")
	ErrorUserRecordNotExist = errors.New("该用户记录不存在, 请先注册")
	ErrorRedisEmpty         = errors.New("缓存记录不存在")
)

// Response Message
const (
	Success               = "成功"
	Failure               = "失败"
	LoginSuccess          = "登录成功"
	LogOutSuccess         = "登出成功"
	RegisterSuccess       = "注册成功"
	ModifyPasswordSuccess = "修改密码成功, 请重新登录"
	UploadSuccess         = "文件上传成功"
	NoPermission          = "无操作权限"
	RoleCreateSuccess     = "角色创建成功"
)

// API Error
var (
	ErrorUserRecordExist = errors.New("该用户名已经被注册")
	ErrorParamsParse     = errors.New("请求参数错误")
	ErrorUserRegister    = errors.New("注册失败, 请重试")
	ErrorUserNotLogin    = errors.New("未登录, 请重试")
	ErrorUserPassword    = errors.New("密码错误, 请重试")
	ErrorOriginPassword  = errors.New("原始密码不正确")
	ErrorUploadFile      = errors.New("上传文件失败")
	ErrorNotEnoughParam  = errors.New("参数不完整")
	ErrorLogOutForced    = errors.New("修改密码成功, 退出登录失败, 请强制退出")
	ErrorSessionsInvalid = errors.New("会话失效, 请重新登录")
)
