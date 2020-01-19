package service

import (
	"fmt"

	"github.com/Scfy-Code/US/app/account/mapper"
)

type accountService struct {
	accountMapper mapper.AccountMapper
}

// NewAccountService 新建账号服务
func NewAccountService() AccountService {
	return accountService{
		mapper.NewAccountMapper(),
	}
}
func (as accountService) SelectAccount(email, password string) bool {
	account, err := as.accountMapper.SelectAccount(email, password)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Println(account)
	return true
}
func (as accountService) InsertAccount(email, password, password0 string) bool {
	// var (
	// 	account entity.Account = entity.Account{
	// 		time.Now().UnixNano(), email, password, "", "", "",
	// 	}
	// )
	//result := as.accountMapper.InsertAccount(account)
	// if result == 1 {
	// 	return true
	// }
	return false
}
