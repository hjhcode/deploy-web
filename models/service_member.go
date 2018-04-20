package models

import (
	. "github.com/hjhcode/deploy-web/common/store"
)

type ServiceMember struct {
	Id        int64
	ServiceId int64
	AccountId int64
}

//增加
func (this ServiceMember) Add(serviceMember *ServiceMember) (int64, error) {
	_, err := OrmWeb.Insert(serviceMember)
	if err != nil {
		return 0, err
	}
	return serviceMember.Id, nil
}

//删除
func (this ServiceMember) Remove(id int64) error {
	serviceMember := new(ServiceMember)
	_, err := OrmWeb.Id(id).Delete(serviceMember)
	return err
}

//修改
func (this ServiceMember) Update(serviceMember *ServiceMember) error {
	_, err := OrmWeb.Id(serviceMember.Id).Update(serviceMember)
	return err
}

//查询(根据服务成员id查询）
func (this ServiceMember) GetById(id int64) (*ServiceMember, error) {
	serviceMember := new(ServiceMember)
	has, err := OrmWeb.Id(id).Get(serviceMember)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return serviceMember, nil
}

//查询（根据服务id查询）
func (this ServiceMember) QueryByServiceId(serviceId int64) ([]*ServiceMember, error) {
	serviceMemberList := make([]*ServiceMember, 0)
	err := OrmWeb.Where("service_id = ?", serviceId).Find(&serviceMemberList)
	if err != nil {
		return nil, err
	}
	return serviceMemberList, nil
}

//查询(根据用户id查询）
func (this ServiceMember) QueryByAccountId(accountId int64) ([]*ServiceMember, error) {
	serviceMemberList := make([]*ServiceMember, 0)
	err := OrmWeb.Where("account_id = ?", accountId).Find(&serviceMemberList)
	if err != nil {
		return nil, err
	}
	return serviceMemberList, nil
}

//删除服务内所有成员
func (this ServiceMember) DelByServiceId(serviceId int64) error {
	sql := "delete from service_member where service_id = ?"
	_, err := OrmWeb.Exec(sql, serviceId)
	return err
}

//根据服务id和用户id查找某一个人
func (this ServiceMember) SearchServiceMember(serviceId int64, accountId int64) (*ServiceMember, error) {
	serviceMember := new(ServiceMember)
	has, err := OrmWeb.Where(" service_id = ? and account_id = ?", serviceId, accountId).Get(serviceMember)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return serviceMember, nil
}
