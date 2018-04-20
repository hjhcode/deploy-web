package models

import (
	"testing"
)

func TestDeployAdd(t *testing.T) {
	InitAllInTest()
	//tm2, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	deploy := &Deploy{ServiceId: 1, AccountId: 1}
	if _, err := deploy.Add(deploy); err != nil {
		t.Error("Add() failed.Error:", err)
	}
}

func TestDeployRemove(t *testing.T) {
	InitAllInTest()
	var deploy Deploy
	if err := deploy.Remove(1); err != nil {
		t.Error("Remove() failed.Error:", err)
	}
}

func TestDeployGetById(t *testing.T) {
	InitAllInTest()
	deploy := &Deploy{ServiceId: 1, AccountId: 1}
	deploy.Add(deploy)

	getDeploy, err := deploy.GetById(deploy.Id)
	if err != nil {
		t.Error("GetById() failed.Error:", err)
	}

	if *getDeploy != *deploy {
		t.Error("GetById() failed", getDeploy, "!=", deploy)
	}
}
