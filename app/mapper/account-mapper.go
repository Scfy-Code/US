package mapper

import "github.com/Scfy-Code/US/sys"

type accountMapper struct {
}

// NewAccountMapper 新建数据库交互对象
func NewAccountMapper() AccountMapper {
	return accountMapper{}
}
func (am accountMapper) SelectAccount(email, password string) ([]map[string]interface{}, error) {
	return sys.Select("US", "select * from account where email=? and password=?", email, password)
}
func (am accountMapper) InsertAccount(map[string]interface{}) (int8, error) {
	return 0, nil
}
func (am accountMapper) UpdateAccount(map[string]interface{}) (int8, error) {
	return 0, nil
}
func (am accountMapper) DeleteAccount(map[string]interface{}) (int8, error) {
	return 0, nil
}
