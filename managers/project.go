package managers

import (
	"strings"

	"strconv"

	"time"

	"github.com/hjhcode/deploy-web/models"
)

func IsProjectCreator(projectId int64, accountId int64) bool {
	project, err := models.Project{}.GetById(projectId)
	if err != nil {
		panic(err.Error())
	}
	if project.AccountId != accountId {
		return false
	}
	return true
}

func IsProjectMember(projectId int64, accountId int64) bool {
	project, err := models.Project{}.GetById(projectId)
	if err != nil {
		panic(err.Error())
	}
	if project.AccountId != accountId || project.ProjectMember == "" {
		return false
	}

	var i int
	array := strings.Split(project.ProjectMember, ";")
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

func DelProject(projectId int64) bool {
	project := &models.Project{Id: projectId, IsDel: 1}
	err := models.Project{}.Update(project)
	if err != nil {
		panic(err.Error())
	}
	return true
}

//还需判断之前数据库中有没有这个工程名，如果有则提示用户不能起这个名字
func AddNewProject(project *models.Project) bool {
	createTime := time.Now()
	project.CreateDate = createTime
	project.UpdateDate = createTime
	_, err := models.Project{}.Add(project)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func UpdateProject(project *models.Project) bool {
	updateTime := time.Now()
	project.UpdateDate = updateTime
	err := models.Project{}.Update(project)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func CheckProjectName(projectName string) bool {
	return true
}
