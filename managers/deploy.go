package managers

import (
	"time"

	"github.com/hjhcode/deploy-web/models"
)

func GetAllDeploy(size int, requestPage int) ([]map[string]interface{}, int) {
	deployList, err := models.Deploy{}.QueryAllDeployByPage(size, (requestPage-1)*size)
	if err != nil {
		panic(err.Error())
	}

	if deployList == nil {
		return nil, 0
	}

	var deployLists []map[string]interface{}
	for i := 0; i < len(deployList); i++ {
		startTime := time.Unix(deployList[i].DeployStart, 0).Format("2006-01-02 15:04:05")
		endTime := time.Unix(deployList[i].DeployEnd, 0).Format("2006-01-02 15:04:05")
		deploy := make(map[string]interface{})
		deploy["id"] = deployList[i].Id
		deploy["account_id"] = getCreator(deployList[i].AccountId)
		deploy["service_id"] = getServiceName(deployList[i].ServiceId)
		deploy["deploy_start"] = startTime
		deploy["deploy_end"] = endTime
		deploy["deploy_statu"] = deployList[i].DeployStatu
		deployLists = append(deployLists, deploy)
	}

	count, err := models.Deploy{}.CountAllDeployByPage()
	if err != nil {
		panic(err.Error())
	}

	return deployLists, int(count)
}
