package managers

import (
	"strings"

	"strconv"

	"github.com/hjhcode/deploy-web/models"
)

func IsProjectCreator(projectId int64, accountId int64) bool {
	project, err := models.Project{}.GetById(projectId)
	if err != nil {
		panic(err)
	}
	if project.AccountId != accountId {
		return false
	}
	return true
}

func IsProjectMember(projectId int64, accountId int64) bool {
	project, err := models.Project{}.GetById(projectId)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	return true
}

func AddNewProject() {

}

func UpdateProject() {

}
