package models

import (
	"testing"
)

func TestServiceAdd(t *testing.T) {
	InitAllInTest()
	service := &Service{AccountId: 1, ServiceName: "aaa"}
	if _, err := service.Add(service); err != nil {
		t.Error("Add() failed.Error:", err)
	}
}

func TestServiceUpdate(t *testing.T) {
	InitAllInTest()
	service := &Service{Id: 1, AccountId: 1, ServiceName: "shiyiisapig"}
	if err := service.Update(service); err != nil {
		t.Error("Update() failed.Error:", err)
	}
}

func TestServiceRemove(t *testing.T) {
	InitAllInTest()
	var service Service
	if err := service.Remove(1); err != nil {
		t.Error("Remove() failed.Error:", err)
	}
}

func TestServiceGetById(t *testing.T) {
	InitAllInTest()
	service := &Service{AccountId: 1, ServiceName: "aaa"}
	service.Add(service)

	getService, err := service.GetById(service.Id)
	if err != nil {
		t.Error("GetById() failed.Error:", err)
	}

	if *getService != *service {
		t.Error("GetById() failed", getService, "!=", service)
	}
}
