package service

// AccountService 账号服务接口
type AccountService interface {
	SelectAccount(email, password string) bool
	InsertAccount(email, password, password0 string) bool
}
