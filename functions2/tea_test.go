package main

import (
	"testing"
)

func TestNoSugar(t *testing.T) {
	var resources = teaResources{cleanMugs: 3, teabagsLeft: 1, isMilkFresh: true, sugarLeft: 0}
	var expectedErr = "There are 0 spoonfuls of sugar remaining, but 2 were requested"
	var _, err = addSugar(&resources, 2)
	if err == nil {
		t.Errorf("Should have thrown error for no remaining sugar")
	}
	if err.Error() != expectedErr {
		t.Errorf("Should have thrown error: \n %v \nActually Threw: \n %v", expectedErr, err.Error())
	}
}