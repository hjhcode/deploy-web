package managers

import (
	"encoding/json"
	"strconv"
	"strings"

	"time"

	"github.com/hjhcode/deploy-web/models"
)

func DelService(serviceId int64, accountId int64) (bool, string) {
	result := isServiceCreator(serviceId, accountId)
	if !result {
		return false, "您没有权限"
	}
	service := &models.Service{Id: serviceId, IsDel: 1}
	err := models.Service{}.Update(service)
	if err != nil {
		panic(err.Error())
	}
	//mess := &components.SendMess{OrderType: 3, DataId: serviceId}
	//components.Send("deploy", mess)

	return true, ""
}

func AddNewService(accountId int64, serviceName string, serviceDescribe string, hostList string,
	mirrorList int64, dockerConfig string, serviceMember string) (bool, string) {
	result := checkServiceName(serviceName)
	if !result {
		return false, "服务名已存在"
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
		var member string
		array := strings.Split(serviceMember, ",")
		for i := 0; i < len(array); i++ {
			account, err := models.Account{}.GetByName(array[i])
			if err != nil {
				panic(err.Error())
			}
			if account == nil {
				return false, "用户成员不存在"
			}

			member += strconv.FormatInt(account.Id, 10)
			if i < len(array)-1 {
				member += ";"
			}
		}
		service.ServiceMember = member
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
	mirrorList int64, dockerConfig string, serviceMember string) (bool, string) {
	result := isServiceMember(serviceId, accountId)
	if !result {
		return false, "您没有权限"
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
		var member string
		array := strings.Split(serviceMember, ",")
		for i := 0; i < len(array); i++ {
			account, err := models.Account{}.GetByName(array[i])
			if err != nil {
				panic(err.Error())
			}
			if account == nil {
				return false, "用户成员不存在"
			}

			member += strconv.FormatInt(account.Id, 10)
			if i < len(array)-1 {
				member += ";"
			}
		}
		service.ServiceMember = member
	}

	service.UpdateDate = time.Now().Unix()
	err := models.Service{}.Update(service)
	if err != nil {
		panic(err.Error())
	}

	return true, ""
}

//func DeployService(serviceId int64, accountId int64) (bool, string, int64) {
//	result := isServiceMember(serviceId, accountId)
//	if !result {
//		return false, "您没有权限", 0
//	}
//
//	deployResult, err := models.Deploy{}.GetByServiceId(serviceId)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	if deployResult != nil && (deployResult.DeployStatu == 0 || deployResult.DeployStatu == 1) {
//		return false, "服务正在部署中", 0
//	}
//
//	service, _ := models.Service{}.GetById(serviceId)
//
//	deploy := &models.Deploy{
//		ServiceId:    serviceId,
//		AccountId:    accountId,
//		HostList:     service.HostList,
//		MirrorList:   service.MirrorList,
//		DockerConfig: service.DockerConfig,
//	}
//	deploy.DeployStart = time.Now().Unix()
//	deploy.DeployEnd = time.Now().Unix()
//	deploy.DeployStatu = 0 //待部署
//	id, err := models.Deploy{}.Add(deploy)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	return true, "", id
//}

func DeployService(serviceId int64, accountId int64) (bool, string, int64) {
	result := isServiceMember(serviceId, accountId)
	if !result {
		return false, "您没有权限", 0
	}

	deployResult, err := models.Deploy{}.GetByServiceId(serviceId)
	if err != nil {
		panic(err.Error())
	}

	if deployResult != nil && (deployResult.DeployStatu == 0 || deployResult.DeployStatu == 1) {
		return false, "服务正在部署中", 0
	}

	service, _ := models.Service{}.GetById(serviceId)
	data := changeJsonToServiceData(service.HostList)

	var deploys models.DeployData
	for _, value := range data.Stage {
		deployStage := models.DeployStage{}
		deployMachine := models.DeployMachine{}
		if value.Machine != nil {
			for _, values := range value.Machine {
				deployMachine.Id = values.Id
				deployMachine.Name = getHostName(values.Id)
				deployMachine.Machine_status = 0
				deployStage.Machine = append(deployStage.Machine, deployMachine)
				deployStage.Stage_status = 0
			}
		} else {
			deployStage.Machine = nil
			deployStage.Stage_status = 0
		}

		deploys.Stage = append(deploys.Stage, deployStage)
	}

	deployDatas := changeDeployDataToJson(&deploys)

	deploy := &models.Deploy{
		ServiceId:    serviceId,
		AccountId:    accountId,
		HostList:     deployDatas,
		MirrorList:   service.MirrorList,
		DockerConfig: service.DockerConfig,
	}

	deploy.DeployStart = time.Now().Unix()
	deploy.DeployEnd = time.Now().Unix()
	deploy.DeployStatu = 0 //待部署
	id, err := models.Deploy{}.Add(deploy)
	if err != nil {
		panic(err.Error())
	}

	return true, "", id
}

//func GetAllServiceByPage(size int, requestPage int) ([]map[string]interface{}, int) {
//	serviceList, err := models.Service{}.QueryAllServiceByPage(size, (requestPage-1)*size)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	if serviceList == nil {
//		return nil, 0
//	}
//
//	var serviceLists []map[string]interface{}
//	for i := 0; i < len(serviceList); i++ {
//		createtime := time.Unix(serviceList[i].CreateDate, 0).Format("2006-01-02 15:04:05")
//		updatetime := time.Unix(serviceList[i].UpdateDate, 0).Format("2006-01-02 15:04:05")
//		services := make(map[string]interface{})
//		services["id"] = serviceList[i].Id
//		services["account_id"] = getCreator(serviceList[i].AccountId)
//		services["service_name"] = serviceList[i].ServiceName
//		services["service_describe"] = serviceList[i].ServiceDescribe
//		services["create_date"] = createtime
//		services["update_date"] = updatetime
//		serviceLists = append(serviceLists, services)
//	}
//
//	count, err := models.Service{}.CountAllService()
//	if err != nil {
//		panic(err.Error())
//	}
//
//	return serviceLists, int(count)
//}

func GetAllService() ([]map[string]interface{}, int) {
	serviceList, err := models.Service{}.QueryAllService()
	if err != nil {
		panic(err.Error())
	}

	if serviceList == nil {
		return nil, 0
	}

	var serviceLists []map[string]interface{}
	for i := 0; i < len(serviceList); i++ {
		createtime := time.Unix(serviceList[i].CreateDate, 0).Format("2006-01-02 15:04:05")
		updatetime := time.Unix(serviceList[i].UpdateDate, 0).Format("2006-01-02 15:04:05")
		services := make(map[string]interface{})
		services["id"] = serviceList[i].Id
		services["account_name"] = getCreator(serviceList[i].AccountId)
		services["service_name"] = serviceList[i].ServiceName
		services["service_describe"] = serviceList[i].ServiceDescribe
		services["create_date"] = createtime
		services["update_date"] = updatetime
		services["service_member"] = getServiceMember(serviceList[i].ServiceMember)
		services["service_statu"] = serviceList[i].ServiceStatu
		serviceLists = append(serviceLists, services)
	}

	count, err := models.Service{}.CountAllService()
	if err != nil {
		panic(err.Error())
	}

	return serviceLists, int(count)
}

func GetServiceByParam(serviceName string) ([]map[string]interface{}, int) {
	service := &models.Service{IsDel: 0}
	serviceList, err := models.Service{}.QueryServiceBySearch(serviceName, service)
	if err != nil {
		panic(err.Error())
	}

	if serviceList == nil {
		return nil, 0
	}

	var serviceLists []map[string]interface{}
	for i := 0; i < len(serviceList); i++ {
		createtime := time.Unix(serviceList[i].CreateDate, 0).Format("2006-01-02 15:04:05")
		updatetime := time.Unix(serviceList[i].UpdateDate, 0).Format("2006-01-02 15:04:05")
		services := make(map[string]interface{})
		services["id"] = serviceList[i].Id
		services["account_name"] = getCreator(serviceList[i].AccountId)
		services["service_name"] = serviceList[i].ServiceName
		services["service_describe"] = serviceList[i].ServiceDescribe
		services["create_date"] = createtime
		services["update_date"] = updatetime
		services["service_statu"] = serviceList[i].ServiceStatu
		serviceLists = append(serviceLists, services)
	}

	count, err := models.Service{}.CountServiceBySearch(serviceName, service)
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

	data := changeJsonToServiceData(service.HostList)

	var host models.ServiceData
	for _, value := range data.Stage {
		hostStage := models.ServiceStage{}
		hostMachine := models.ServiceMachine{}
		for _, values := range value.Machine {
			hostMachine.Id = values.Id
			hostMachine.Name = getHostName(values.Id)
			hostStage.Machine = append(hostStage.Machine, hostMachine)
		}
		host.Stage = append(host.Stage, hostStage)
	}

	jsonData := changeServiceDataToJson(host)

	services := make(map[string]interface{})
	services["id"] = serviceId
	services["service_name"] = service.ServiceName
	services["service_describe"] = service.ServiceDescribe
	services["host_list"] = jsonData
	services["mirror_list"] = getMirrorName(service.MirrorList)
	services["docker_config"] = service.DockerConfig
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
			member += ","
		}
	}
	return member
}

func getServiceName(serviceId int64) string {
	service, err := models.Service{}.GetById(serviceId)
	if err != nil {
		panic(err.Error())
	}

	if service == nil {
		return ""
	}

	return service.ServiceName
}

func changeJsonToServiceData(hostList string) *models.ServiceData {
	data := &models.ServiceData{}
	if err := json.Unmarshal([]byte(hostList), &data); err != nil {
		panic(err.Error())
	}

	return data
}

func changeServiceDataToJson(serviceData models.ServiceData) string {
	data, err := json.Marshal(serviceData)
	if err != nil {
		panic(err.Error())
	}
	return string(data)
}
