package errno

import "errors"

// System Error
var (
	ErrorNotFoundService    = errors.New("未找到服务应用与环境信息")
	ErrorDataBase           = errors.New("数据库操作失败, 请重试")
	ErrorUserRecordNotExist = errors.New("该用户记录不存在")
	ErrorRedisEmpty         = errors.New("缓存记录不存在")
)

// Response Message
const (
	Success = "成功"
	Failure = "失败"
)

// API Error
var (
	ErrorUserRecordExist = errors.New("该用户名已经被注册")
	ErrorParamsParse     = errors.New("请求参数错误")
	ErrorUserRegister    = errors.New("注册失败, 请重试")
	ErrorUserNotLogin    = errors.New("未登录, 请重试")
)
