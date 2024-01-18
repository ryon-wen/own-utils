package code

const (
	Success         = 10000 + iota // 成功
	LoseArgument                   // 参数相关错误
	BadRequest                     // 请求相关错误
	InternalError                  // 服务层相关错误
	OperationFailed                // 代码执行相关错误
)
