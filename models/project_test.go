package models

import (
	"testing"
)

func TestProjectAdd(t *testing.T) {
	InitAllInTest()
	project := &Project{AccountId: 1, ProjectName: "testProject"}
	if _, err := project.Add(project); err != nil {
		t.Error("Add() failed.Error:", err)
	}
}

func TestProjectUpdate(t *testing.T) {
	InitAllInTest()
	project := &Project{Id: 1, AccountId: 1, ProjectName: "ttest"}
	if err := project.Update(project); err != nil {
		t.Error("Update() failed.Error:", err)
	}
}

func TestProjectRemove(t *testing.T) {
	InitAllInTest()
	var project Project
	if err := project.Remove(1); err != nil {
		t.Error("Delete() failed.Error:", err)
	}
}

func TestProjectGetById(t *testing.T) {
	InitAllInTest()

	project := &Project{AccountId: 12, ProjectName: "testProject"}
	project.Add(project)

	getProject, err := project.GetById(project.Id)

	if err != nil {
		t.Error("GetById() failed.Error:", err)
	}

	if *getProject != *project {
		t.Error("GetById() failed:", getProject, "!=", project)
	}
}