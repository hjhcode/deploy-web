package managers

import (
	"strconv"
	"strings"

	"time"

	"github.com/hjhcode/deploy-web/models"
)

func DelService(serviceId int64, accountId int64) (bool, string) {
	result := isServiceCreator(serviceId, accountId)
	if !result {
		return false, "You have no authority"
	}
	err := models.Service{}.Remove(serviceId)
	if err != nil {
		panic(err.Error())
	}

	return true, ""
}

func AddNewService(accountId int64, serviceName string, serviceDescribe string, hostList string,
	mirrorList string, dockerConfig string, serviceMember string) (bool, string) {
	result := checkServiceName(serviceName)
	if !result {
		return false, "service name is exist"
	}
	service := &models.Service{
		AccountId:       accountId,
		ServiceName:     serviceName,
		ServiceDescribe: serviceDescribe,
		HostList:        hostList,
		MirrorList:      mirrorList,
		DockerConfig:    dockerConfig,
	}

	if serviceMember != "" {
		service.ServiceMember = serviceMember
	}
	createTime := time.Now().Unix()
	service.CreateDate = createTime
	service.UpdateDate = createTime
	_, err := models.Service{}.Add(service)
	if err != nil {
		panic(err.Error())
	}

	return true, ""
}

func UpdateService(accountId int64, serviceId int64, serviceName string, serviceDescribe string, hostList string,
	mirrorList string, dockerConfig string, serviceMember string) (bool, string) {
	result := isServiceMember(serviceId, accountId)
	if !result {
		return false, "You have no authority"
	}
	service := &models.Service{
		Id:              serviceId,
		ServiceName:     serviceName,
		ServiceDescribe: serviceDescribe,
		HostList:        hostList,
		MirrorList:      mirrorList,
		DockerConfig:    dockerConfig,
	}
	if serviceMember != "" {
		service.ServiceMember = serviceMember
	}
	service.UpdateDate = time.Now().Unix()
	err := models.Service{}.Update(service)
	if err != nil {
		panic(err.Error())
	}

	return true, ""
}

func DeployService(serviceId int64, accountId int64) (bool, string) {
	result := isServiceMember(serviceId, accountId)
	if !result {
		return false, "You have no authority"
	}
	service, _ := models.Service{}.GetById(serviceId)
	deploy := &models.Deploy{
		ServiceId:    serviceId,
		AccountId:    accountId,
		HostList:     service.HostList,
		MirrorList:   service.MirrorList,
		DockerConfig: service.DockerConfig,
	}
	deploy.DeployStart = time.Now().Unix()
	deploy.DeployStatu = 1 //部署中
	_, err := models.Deploy{}.Add(deploy)
	if err != nil {
		panic(err.Error())
	}

	return true, ""
}

func GetAllService(size int, requestPage int) ([]map[string]interface{}, int) {
	serviceList, err := models.Service{}.QueryAllServiceByPage(size, (requestPage-1)*size)
	if err != nil {
		panic(err.Error())
	}

	var serviceLists []map[string]interface{}
	for i := 0; i < len(serviceList); i++ {
		createtime := time.Unix(serviceList[i].CreateDate, 0).Format("2006-01-02 15:04:05")
		updatetime := time.Unix(serviceList[i].UpdateDate, 0).Format("2006-01-02 15:04:05")
		services := make(map[string]interface{})
		services["id"] = serviceList[i].Id
		services["account_id"] = getCreator(serviceList[i].AccountId)
		services["service_name"] = serviceList[i].ServiceName
		services["service_describe"] = serviceList[i].ServiceDescribe
		services["create_date"] = createtime
		services["update_date"] = updatetime
		serviceLists = append(serviceLists, services)
	}

	count, err := models.Service{}.CountAllServiceByPage()
	if err != nil {
		panic(err.Error())
	}

	return serviceLists, int(count)
}

func GetServiceByParam(serviceName string, size int, requestPage int) ([]map[string]interface{}, int) {
	serviceList, err := models.Service{}.QueryServiceBySearch(serviceName, size, (requestPage-1)*size)
	if err != nil {
		panic(err.Error())
	}

	var serviceLists []map[string]interface{}
	for i := 0; i < len(serviceList); i++ {
		createtime := time.Unix(serviceList[i].CreateDate, 0).Format("2006-01-02 15:04:05")
		updatetime := time.Unix(serviceList[i].UpdateDate, 0).Format("2006-01-02 15:04:05")
		services := make(map[string]interface{})
		services["id"] = serviceList[i].Id
		services["account_id"] = getCreator(serviceList[i].AccountId)
		services["service_name"] = serviceList[i].ServiceName
		services["service_describe"] = serviceList[i].ServiceDescribe
		services["create_date"] = createtime
		services["update_date"] = updatetime
		serviceLists = append(serviceLists, services)
	}

	count, err := models.Service{}.CountServiceBySearch(serviceName)
	if err != nil {
		panic(err.Error())
	}

	return serviceLists, int(count)
}

func GetOneService(serviceId int64) map[string]interface{} {
	service, err := models.Service{}.GetById(serviceId)
	if err != nil {
		panic(err.Error())
	}
	createtime := time.Unix(service.CreateDate, 0).Format("2006-01-02 15:04:05")
	updatetime := time.Unix(service.UpdateDate, 0).Format("2006-01-02 15:04:05")
	services := make(map[string]interface{})
	services["id"] = serviceId
	services["account_id"] = getCreator(service.AccountId)
	services["service_name"] = service.ServiceName
	services["service_describe"] = service.ServiceDescribe
	services[" host_list"] = service.HostList
	services["mirror_list"] = getMirrorName(service.MirrorList)
	services["docker_config"] = service.DockerConfig
	services["create_date"] = createtime
	services["update_date"] = updatetime
	services["service_member"] = getServiceMember(service.ServiceMember)

	return services
}

func checkServiceName(serviceName string) bool {
	service, err := models.Service{}.GetByServiceName(serviceName)
	if err != nil {
		panic(err.Error())
	}
	if service != nil {
		return false
	}

	return true
}

func isServiceCreator(serviceId int64, accountId int64) bool {
	service, err := models.Service{}.GetById(serviceId)
	if err != nil {
		panic(err.Error())
	}

	if service.AccountId != accountId {
		return false
	}
	return true
}

func isServiceMember(serviceId int64, accountId int64) bool {
	service, err := models.Service{}.GetById(serviceId)
	if err != nil {
		panic(err.Error())
	}

	if service.AccountId == accountId {
		return true
	}

	if service.ServiceMember != "" {
		var i int
		array := strings.Split(service.ServiceMember, ";")
		userId := strconv.FormatInt(accountId, 10)
		for i = 0; i < len(array); i++ {
			if array[i] == userId {
				return true
			}
		}
	}

	return false
}

func getServiceMember(serviceMember string) string {
	if serviceMember == "" {
		return ""
	}
	var member string
	array := strings.Split(serviceMember, ";")
	for i := 0; i < len(array); i++ {
		id, _ := strconv.ParseInt(array[i], 10, 64)
		account, err := models.Account{}.GetById(id)
		if err != nil {
			panic(err.Error())
		}
		member += account.Name
		if i < len(array)-1 {
			member += ";"
		}
	}
	return member
}
