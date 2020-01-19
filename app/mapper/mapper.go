package mapper

// AccountMapper 数据库交互接口
type AccountMapper interface {
	SelectAccount(string, string) ([]map[string]interface{}, error)
	InsertAccount(map[string]interface{}) (int8, error)
	UpdateAccount(map[string]interface{}) (int8, error)
	DeleteAccount(map[string]interface{}) (int8, error)
}
