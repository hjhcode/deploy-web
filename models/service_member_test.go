package models

import (
	"testing"
)

func TestServiceMemberAdd(t *testing.T) {
	InitAllInTest()
	serviceMember := &ServiceMember{ServiceId: 1, AccountId: 1}
	if _, err := serviceMember.Add(serviceMember); err != nil {
		t.Error("Add() failed.Error:", err)
	}
}

func TestServiceMemberUpdate(t *testing.T) {
	InitAllInTest()
	serviceMember := &ServiceMember{Id: 1, ServiceId: 1, AccountId: 1}
	if err := serviceMember.Update(serviceMember); err != nil {
		t.Error("Update() failed.Error:", err)
	}
}

func TestServiceMemberRemove(t *testing.T) {
	InitAllInTest()
	var serviceMember ServiceMember
	if err := serviceMember.Remove(1); err != nil {
		t.Error("Remove() failed.Error:", err)
	}
}

func TestServiceMemberGetById(t *testing.T) {
	InitAllInTest()
	serviceMember := &ServiceMember{ServiceId: 1, AccountId: 1}
	serviceMember.Add(serviceMember)

	getServiceMember, err := serviceMember.GetById(serviceMember.Id)
	if err != nil {
		t.Error("GetById() failed.Error:", err)
	}

	if *getServiceMember != *serviceMember {
		t.Error("GetById() failed", getServiceMember, "!=", serviceMember)
	}
}
