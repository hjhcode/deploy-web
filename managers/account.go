package managers

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/hjhcode/deploy-web/common/components"
	"github.com/hjhcode/deploy-web/models"
)

func AccountLogin(name, password string) (bool, string, string) {
	account := getAccountByName(name)
	if account == nil {
		return false, "", "user is not exist"
	}
	if account.Password != md5Encode(password) {
		return false, "", "Password is wrong"
	} else {
		userId := account.Id
		if token, err := components.CreateToken(userId); err != nil {
			panic(err.Error())
		} else {
			return true, token, ""
		}
	}
}

func AccountRegister(name, password string) (bool, int64, string) {
	account := getAccountByName(name)
	if account != nil {
		return false, 0, "user is exist"
	}
	account = &models.Account{Name: name, Password: md5Encode(password)}
	if insertId, err := (models.Account{}).Add(account); err != nil {
		panic(err.Error())
	} else {
		return true, insertId, ""
	}
}

func getAccountByName(name string) *models.Account {
	account, err := models.Account{}.GetByName(name)
	if err != nil {
		panic(err.Error())
	}
	return account
}

func md5Encode(password string) string {
	w := md5.New()
	io.WriteString(w, password)
	md5str := string(fmt.Sprintf("%x", w.Sum(nil)))
	return md5str
}

func getCreator(accountId int64) string {
	account, err := models.Account{}.GetById(accountId)
	if err != nil {
		panic(err.Error())
	}
	return account.Name
}
