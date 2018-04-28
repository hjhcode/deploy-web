package managers

import (
	"strings"

	"strconv"

	"time"

	"regexp"

	"github.com/hjhcode/deploy-web/common/components"
	"github.com/hjhcode/deploy-web/models"
)

func DelProject(projectId int64, accountId int64) (bool, string) {
	result := isProjectCreator(projectId, accountId)
	if !result {
		return false, "You have no authority"
	}
	project := &models.Project{Id: projectId, IsDel: 1}
	err := models.Project{}.Update(project)
	if err != nil {
		panic(err.Error())
	}
	return true, ""
}

func AddNewProject(accountId int64, projectName string, projectDescribe string, gitDockerPath string,
	projectMember string) (bool, string) {
	result := checkProjectName(projectName)
	if !result {
		return false, "project name is exist"
	}
	project := &models.Project{
		AccountId:       accountId,
		ProjectName:     projectName,
		ProjectDescribe: projectDescribe,
		GitDockerPath:   gitDockerPath,
	}

	if projectMember != "" {
		project.ProjectMember = projectMember
	}

	createTime := time.Now().Unix()
	project.CreateDate = createTime
	project.UpdateDate = createTime
	_, err := models.Project{}.Add(project)
	if err != nil {
		panic(err.Error())
	}

	return true, ""
}

func UpdateProject(accountId int64, projectId int64, projectName string, projectDescribe string, gitDockerPath string,
	projectMember string) (bool, string) {
	result := isProjectMember(projectId, accountId)
	if !result {
		return false, "You have no authority"
	}
	project := &models.Project{
		Id:              projectId,
		ProjectName:     projectName,
		ProjectDescribe: projectDescribe,
		GitDockerPath:   gitDockerPath,
	}
	if projectMember != "" {
		project.ProjectMember = projectMember
	}
	project.UpdateDate = time.Now().Unix()
	err := models.Project{}.Update(project)
	if err != nil {
		panic(err.Error())
	}

	return true, ""
}

func ConstructProject(accountId int64, projectID int64) (bool, string) {
	result := isProjectMember(projectID, accountId)
	if !result {
		return false, "You have no authority"
	}
	record, err := models.ConstructRecord{}.GetByProjectId(projectID)
	if err != nil {
		panic(err.Error())
	}

	if record != nil && record.ConstructStatu == 1 {
		return false, "Project is building"
	}

	constructRecord := &models.ConstructRecord{
		AccountId: accountId,
		ProjectId: projectID,
	}
	constructRecord.ConstructStatu = 1 //构建中
	constructRecord.ConstructStart = time.Now().Unix()
	constructRecord.ConstructEnd = 0
	id, err := models.ConstructRecord{}.Add(constructRecord)
	if err != nil {
		panic(err.Error())
	}
	mess := &components.SendMess{SubmitType: "construct", SubmitId: id}
	components.Send("project", mess)
	return true, ""
}

//func GetAllProjectByPage(size int, requestPage int) ([]map[string]interface{}, int) {
//	projectList, err := models.Project{}.QueryAllProjectByPage(size, (requestPage-1)*size)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	if projectList == nil {
//		return nil, 0
//	}
//
//	var projectLists []map[string]interface{}
//	for i := 0; i < len(projectList); i++ {
//		createtime := time.Unix(projectList[i].CreateDate, 0).Format("2006-01-02 15:04:05")
//		updatetime := time.Unix(projectList[i].UpdateDate, 0).Format("2006-01-02 15:04:05")
//		projects := make(map[string]interface{})
//		projects["id"] = projectList[i].Id
//		projects["account_id"] = getCreator(projectList[i].AccountId)
//		projects["project_name"] = projectList[i].ProjectName
//		projects["project_describe"] = projectList[i].ProjectDescribe
//		projects["git_docker_path"] = projectList[i].GitDockerPath
//		projects["create_date"] = createtime
//		projects["update_date"] = updatetime
//		projectLists = append(projectLists, projects)
//	}
//
//	count, err := models.Project{}.CountAllProject()
//	if err != nil {
//		panic(err.Error())
//	}
//
//	return projectLists, int(count)
//}

func GetAllProject() ([]map[string]interface{}, int) {
	projectList, err := models.Project{}.QueryAllProject()
	if err != nil {
		panic(err.Error())
	}

	if projectList == nil {
		return nil, 0
	}

	var projectLists []map[string]interface{}
	for i := 0; i < len(projectList); i++ {
		createtime := time.Unix(projectList[i].CreateDate, 0).Format("2006-01-02 15:04:05")
		updatetime := time.Unix(projectList[i].UpdateDate, 0).Format("2006-01-02 15:04:05")
		projects := make(map[string]interface{})
		projects["id"] = projectList[i].Id
		projects["account_id"] = getCreator(projectList[i].AccountId)
		projects["project_name"] = projectList[i].ProjectName
		projects["project_describe"] = projectList[i].ProjectDescribe
		projects["git_docker_path"] = projectList[i].GitDockerPath
		projects["create_date"] = createtime
		projects["update_date"] = updatetime
		projectLists = append(projectLists, projects)
	}

	count, err := models.Project{}.CountAllProject()
	if err != nil {
		panic(err.Error())
	}

	return projectLists, int(count)
}

func GetProjectByParam(projectName string) ([]map[string]interface{}, int) {
	project := &models.Project{IsDel: 0}
	projectList, err := models.Project{}.QueryProjectBySearch(projectName, project)
	if err != nil {
		panic(err.Error())
	}

	if projectList == nil {
		return nil, 0
	}

	var projectLists []map[string]interface{}
	for i := 0; i < len(projectList); i++ {
		createtime := time.Unix(projectList[i].CreateDate, 0).Format("2006-01-02 15:04:05")
		updatetime := time.Unix(projectList[i].UpdateDate, 0).Format("2006-01-02 15:04:05")
		projects := make(map[string]interface{})
		projects["id"] = projectList[i].Id
		projects["account_id"] = getCreator(projectList[i].AccountId)
		projects["project_name"] = projectList[i].ProjectName
		projects["project_describe"] = projectList[i].ProjectDescribe
		projects["git_docker_path"] = projectList[i].GitDockerPath
		projects["create_date"] = createtime
		projects["update_date"] = updatetime
		projectLists = append(projectLists, projects)
	}

	count, err := models.Project{}.CountBySearch(projectName, project)
	if err != nil {
		panic(err.Error())
	}

	return projectLists, int(count)
}

func GetOneProject(projectId int64) map[string]interface{} {
	project, err := models.Project{}.GetById(projectId)
	if err != nil {
		panic(err.Error())
	}
	createtime := time.Unix(project.CreateDate, 0).Format("2006-01-02 15:04:05")
	updatetime := time.Unix(project.UpdateDate, 0).Format("2006-01-02 15:04:05")
	projects := make(map[string]interface{})
	projects["id"] = projectId
	projects["account_id"] = getCreator(project.AccountId)
	projects["project_name"] = project.ProjectName
	projects["project_describe"] = project.ProjectDescribe
	projects["git_docker_path"] = project.GitDockerPath
	projects["create_date"] = createtime
	projects["update_date"] = updatetime
	projects["project_member"] = getProjectMember(project.ProjectMember)

	return projects
}

func checkProjectName(projectName string) bool {
	project, err := models.Project{}.GetByProjectName(projectName)
	if err != nil {
		panic(err.Error())
	}
	if project != nil {
		return false
	}
	return true
}

func isProjectCreator(projectId int64, accountId int64) bool {
	project, err := models.Project{}.GetById(projectId)
	if err != nil {
		panic(err.Error())
	}
	if project.AccountId != accountId {
		return false
	}
	return true
}

func isProjectMember(projectId int64, accountId int64) bool {
	project, err := models.Project{}.GetById(projectId)
	if err != nil {
		panic(err.Error())
	}
	if project.AccountId == accountId {
		return true
	}

	if project.ProjectMember != "" {
		var i int
		array := strings.Split(project.ProjectMember, ";")
		userId := strconv.FormatInt(accountId, 10)
		for i = 0; i < len(array); i++ {
			if array[i] == userId {
				return true
			}
		}
	}

	return false
}

func getProjectMember(projectMember string) string {
	if projectMember == "" {
		return ""
	}
	var member string
	array := strings.Split(projectMember, ";")
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

func getProjectName(projectId int64) string {
	project, err := models.Project{}.GetById(projectId)
	if err != nil {
		panic(err.Error())
	}

	if project == nil {
		return ""
	}

	return project.ProjectName
}

func MatchGitDockerPath(out string) bool {
	paths := `^((https|http)?:\/\/github.com\/)[^\s]+`
	reg, _ := regexp.Compile(paths)
	loss := reg.FindString(out)
	if loss == "" {
		return false
	}
	return true
}
