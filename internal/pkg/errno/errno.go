package errno

import "errors"

// System Error
var (
	NotFoundServiceError    = errors.New("未找到服务应用与环境信息")
	DataBaseError           = errors.New("数据库操作失败, 请重试")
	UserRecordExistError    = errors.New("该用户名已经被注册")
	UserRecordNotExistError = errors.New("该用户名记录不存在")
)

// Response Message
const (
	Success = "成功"
	Failure = "失败"
)

// API Error
var (
	ParamsParseError  = errors.New("请求参数错误")
	UserRegisterError = errors.New("注册失败, 请重试")
)
