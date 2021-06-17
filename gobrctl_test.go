package gobrctl

import (
	"fmt"
	"testing"
)

func TestGetInterfaceNames(t *testing.T) {
	interfaces := GetInterfaceNames()
	fmt.Print(interfaces)
	if len(interfaces) == 0 {
		t.Fatalf("get interfaces failed")
	}
}

func TestGetBridges(t *testing.T) {
	bridges := GetAllBridges()
	fmt.Print(bridges)
	if len(bridges) == 0 {
		t.Fatalf("get bridges failed")
	}
}
