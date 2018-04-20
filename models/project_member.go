package models

import (
	. "github.com/hjhcode/deploy-web/common/store"
)

type ProjectMember struct {
	Id        int64
	ProjectId int64
	AccountId int64
}

//增加
func (this ProjectMember) Add(projectMember *ProjectMember) (int64, error) {
	_, err := OrmWeb.Insert(projectMember)
	if err != nil {
		return 0, err
	}
	return projectMember.Id, nil
}

//删除
func (this ProjectMember) Remove(id int64) error {
	projectMember := new(ProjectMember)
	_, err := OrmWeb.Id(id).Delete(projectMember)
	return err
}

//修改
func (this ProjectMember) Update(projectMember *ProjectMember) error {
	_, err := OrmWeb.Id(projectMember.Id).Update(projectMember)
	return err
}

//查询(根据工程成员id查询）
func (this ProjectMember) GetById(id int64) (*ProjectMember, error) {
	projectMember := new(ProjectMember)
	has, err := OrmWeb.Id(id).Get(projectMember)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return projectMember, nil
}

//查询(根据工程id查询）
func (this ProjectMember) QueryByProjectId(projectId int64) ([]*ProjectMember, error) {
	projectMember := make([]*ProjectMember, 0)
	err := OrmWeb.Where("project_id = ?", projectId).Find(&projectMember)
	if err != nil {
		return nil, err
	}
	return projectMember, nil
}

//查询(根据用户id查询）
func (this ProjectMember) QueryByAccountId(accountId int64) ([]*ProjectMember, error) {
	projectMember := make([]*ProjectMember, 0)
	err := OrmWeb.Where("account_id = ?", accountId).Find(&projectMember)
	if err != nil {
		return nil, err
	}
	return projectMember, nil
}

//删除工程内所有成员
func (this ProjectMember) DelByProjectId(projectId int64) error {
	sql := "delete from project_member where project_id = ?"
	_, err := OrmWeb.Exec(sql, projectId)
	return err
}

//根据工程id和用户id查找某一个人
func (this ProjectMember) SearchProjectMember(projectId int64, accountId int64) (*ProjectMember, error) {
	projectMember := new(ProjectMember)
	has, err := OrmWeb.Where(" project_id = ? and account_id = ?", projectId, accountId).Get(projectMember)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return projectMember, nil
}
