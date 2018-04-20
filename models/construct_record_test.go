package models

import (
	"testing"
)

func TestConstructRecordAdd(t *testing.T) {
	InitAllInTest()
	constructRecord := &ConstructRecord{AccountId: 1, ProjectId: 1}
	if _, err := constructRecord.Add(constructRecord); err != nil {
		t.Error("Add() failed.Error:", err)
	}
}

func TestConstructRecordRemove(t *testing.T) {
	InitAllInTest()
	var constructRecord ConstructRecord
	if err := constructRecord.Remove(1); err != nil {
		t.Error("Remove() failed.Error:", err)
	}
}

func TestConstructRecordGetById(t *testing.T) {
	InitAllInTest()
	constructRecord := &ConstructRecord{AccountId: 1, ProjectId: 1}
	constructRecord.Add(constructRecord)

	getConstructRecord, err := constructRecord.GetById(constructRecord.Id)
	if err != nil {
		t.Error("GetById() failed.Error:", err)
	}

	if *getConstructRecord != *constructRecord {
		t.Error("GetById() failed", getConstructRecord, "!=", constructRecord)
	}
}
