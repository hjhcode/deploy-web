package managers

import (
	"strconv"
	"strings"

	"github.com/hjhcode/deploy-web/models"
)

func IsServiceCreator(serviceId int64, accountId int64) bool {
	service, err := models.Service{}.GetById(serviceId)
	if err != nil {
		panic(err.Error())
	}

	if service.AccountId != accountId {
		return false
	}
	return true
}

func IsServiceMember(serviceId int64, accountId int64) bool {
	service, err := models.Service{}.GetById(serviceId)
	if err != nil {
		panic(err.Error())
	}

	if service.AccountId != accountId || service.ServiceMember == "" {
		return false
	}
	var i int
	array := strings.Split(service.ServiceMember, ";")
	userId := strconv.FormatInt(accountId, 10)
	for i = 0; i < len(array); i++ {
		if array[i] == userId {
			return true
		}
	}
	if i == len(array) {
		return false
	}
	return true
}

func DelService(serviceId int64) bool {
	err := models.Service{}.Remove(serviceId)
	if err != nil {
		panic(err.Error())
	}
	return true
}

//进行名字查重检验
func AddNewService() bool {
	return true
}
