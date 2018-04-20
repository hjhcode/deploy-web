package models

import (
	"testing"
)

func TestProjectMemberAdd(t *testing.T) {
	InitAllInTest()
	projectMember := &ProjectMember{ProjectId: 1, AccountId: 1}
	if _, err := projectMember.Add(projectMember); err != nil {
		t.Error("Add() failed.Error:", err)
	}
}

func TestProjectMemberUpdate(t *testing.T) {
	InitAllInTest()
	projectMember := &ProjectMember{Id: 1, ProjectId: 1, AccountId: 1}
	if err := projectMember.Update(projectMember); err != nil {
		t.Error("Update() failed.Error:", err)
	}
}

func TestProjectMemberRemove(t *testing.T) {
	InitAllInTest()
	var projectMember ProjectMember
	if err := projectMember.DelByProjectId(1); err != nil {
		t.Error("Remove() failed.Error:", err)
	}
}

func TestProjectMemberGetById(t *testing.T) {
	InitAllInTest()
	projectMember := &ProjectMember{ProjectId: 1, AccountId: 1}
	projectMember.Add(projectMember)

	getProjectMember, err := projectMember.GetById(projectMember.Id)
	if err != nil {
		t.Error("GetById() failed.Error:", err)
	}

	if *getProjectMember != *projectMember {
		t.Error("GetById() failed", getProjectMember, "!=", projectMember)
	}
}
