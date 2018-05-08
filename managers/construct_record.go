package managers

import (
	"time"

	"github.com/hjhcode/deploy-web/common/components"
	"github.com/hjhcode/deploy-web/models"
)

func GetAllConstructRecord() ([]map[string]interface{}, int) {
	constrcuctList, err := models.ConstructRecord{}.QueryAllConstructRecord()
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
		records["account_name"] = getCreator(constrcuctList[i].AccountId)
		records["project_name"] = getProjectName(constrcuctList[i].ProjectId)
		records["mirror_name"] = getMirrorNameById(constrcuctList[i].MirrorId)
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

//构建页获取详情数据
func GetConstructProjectData(constructId int64) map[string]interface{} {

	constructs, err := models.ConstructRecord{}.GetById(constructId)
	if err != nil {
		panic(err)
	}

	if constructs == nil {
		return nil
	}

	project, err := models.Project{}.GetById(constructs.ProjectId)
	if err != nil {
		panic(err.Error())
	}

	if project == nil {
		return nil
	}

	lists := make(map[string]interface{})
	createtime := time.Unix(project.CreateDate, 0).Format("2006-01-02 15:04:05")
	updatetime := time.Unix(project.UpdateDate, 0).Format("2006-01-02 15:04:05")
	lists["project_name"] = project.ProjectName
	lists["project_describe"] = project.ProjectDescribe
	lists["git_docker_path"] = project.GitDockerPath
	lists["create_date"] = createtime
	lists["update_date"] = updatetime
	lists["account_name"] = getCreator(project.AccountId)
	lists["construct_statu"] = constructs.ConstructStatu
	lists["construct_log"] = constructs.ConstructLog

	return lists
}

func StartConstructProject(accountId int64, constructId int64) (bool, string) {

	record, err := models.ConstructRecord{}.GetById(constructId)
	if err != nil {
		panic(err.Error())
	}

	result := isProjectMember(record.ProjectId, accountId)
	if !result {
		return false, "您没有权限"
	}

	constructRecord := &models.ConstructRecord{
		Id:             constructId,
		ConstructStatu: 1, //构建中
		ConstructEnd:   time.Now().Unix(),
	}

	errs := models.ConstructRecord{}.Update(constructRecord)
	if errs != nil {
		panic(err.Error())
	}

	mess := &components.SendMess{OrderType: 0, DataId: constructId}
	components.Send("deploy", mess)

	return true, ""
}
