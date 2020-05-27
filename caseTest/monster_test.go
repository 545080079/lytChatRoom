package caseTest

import "testing"

func TestMonster_Store(t *testing.T) {

	monster := monster{
		Name: "luoyutao",
		Age:12,
		Skill: "play",
	}

	res := monster.Store()
	if !res{
		t.Fatalf("Store() fail ! hope = %v, but = %v",true, res)
	}
	t.Logf("Store() ok !")
}

func TestMonster_ReStore(t *testing.T) {
	monster := &monster{}
	res := monster.ReStore()
	if !res{
		t.Fatalf("Restore() fail ! hope = %v, but = %v", "luoyutao", res)
	}
	t.Logf("ReStore() ok !")
}
