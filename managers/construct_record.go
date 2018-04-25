package managers

import (
	"time"

	"github.com/hjhcode/deploy-web/models"
)

func GetAllConstructRecord(size int, requestPage int) ([]map[string]interface{}, int) {
	constrcuctList, err := models.ConstructRecord{}.QueryAllConstructByPage(size, (requestPage-1)*size)
	if err != nil {
		panic(err.Error())
	}

	if constrcuctList == nil {
		return nil, 0
	}

	var constrcuctLists []map[string]interface{}
	for i := 0; i < len(constrcuctList); i++ {
		startTime := time.Unix(constrcuctList[i].ConstructStart, 0).Format("2006-01-02 15:04:05")
		endTime := time.Unix(constrcuctList[i].ConstructEnd, 0).Format("2006-01-02 15:04:05")
		records := make(map[string]interface{})
		records["id"] = constrcuctList[i].Id
		records["account_id"] = getCreator(constrcuctList[i].AccountId)
		records["project_id"] = getProjectName(constrcuctList[i].ProjectId)
		records["mirror_id"] = getMirrorNameById(constrcuctList[i].MirrorId)
		records["construct_start"] = startTime
		records["construct_end"] = endTime
		records["construct_statu"] = constrcuctList[i].ConstructStatu
		constrcuctLists = append(constrcuctLists, records)
	}

	count, err := models.ConstructRecord{}.CountAllConstructByPage()
	if err != nil {
		panic(err.Error())
	}

	return constrcuctLists, int(count)
}
