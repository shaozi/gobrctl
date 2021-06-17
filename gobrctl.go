// Package gobrctl gets list of bridges in Linux the same way as `brctl show`.
package gobrctl

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

// Bridge is a linux bridge.
type Bridge struct {
	Name       string
	Id         string
	Stp        bool
	Interfaces []string
}

const sysClassNet = "/sys/class/net/"

// checkError is a function that will panic if error is not nil.
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// exists is a simple check if a path exists.
// Returns bool.
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// GetInterfaceNames gets all interface names.
// It returns an array of strings, each is an interface's name.
func GetInterfaceNames() []string {
	files, err := ioutil.ReadDir(sysClassNet)
	checkError(err)
	ret := []string{}
	for _, file := range files {
		ret = append(ret, file.Name())
	}
	return ret
}

// GetBridgeByName gets the details of a bridge by its name.
// It returns the Bridge if the name is a bridge
func GetBridgeByName(name string) (Bridge, error) {
	bridge := Bridge{Name: name}
	bridgeFolder := sysClassNet + name + "/bridge"
	if exists(bridgeFolder) {
		id, err := ioutil.ReadFile(bridgeFolder + "/bridge_id")
		checkError(err)
		bridge.Id = strings.TrimSpace(string(id))
		stp, err := ioutil.ReadFile(bridgeFolder + "/stp_state")
		checkError(err)
		bridge.Stp = stp[0] == '1'
		files, err := ioutil.ReadDir(sysClassNet + name + "/brif")
		if err == nil {
			interfaces := []string{}
			for _, file := range files {
				interfaces = append(interfaces, file.Name())
			}
			bridge.Interfaces = interfaces
		}
		return bridge, nil
	}
	return bridge, errors.New("not a bridge")
}

// GetAllBridges gets all bridges in the Linux.
// It returns an array of Bridges.
func GetAllBridges() []Bridge {
	interfaces := GetInterfaceNames()
	ret := []Bridge{}
	for _, intf := range interfaces {
		bridge, err := GetBridgeByName(intf)
		if err == nil {
			ret = append(ret, bridge)
		}
	}
	return ret
}
