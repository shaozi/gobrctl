package gobrctl

import (
	"errors"
	"io/ioutil"
	"os"
)

type Bridge struct {
	Name      string
	Id        string
	Stp       bool
	Interface string
}

const sysClassNet = "/sys/class/net/"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func GetInterfaceNames() []string {
	files, err := ioutil.ReadDir(sysClassNet)
	checkError(err)
	ret := []string{}
	for _, file := range files {
		ret = append(ret, file.Name())
	}
	return ret
}

func GetBridgeByName(name string) (Bridge, error) {
	bridge := Bridge{Name: name}
	bridgeFolder := sysClassNet + name + "/bridge"
	if exists(bridgeFolder) {
		id, err := ioutil.ReadFile(bridgeFolder + "/bridge_id")
		checkError(err)
		bridge.Id = string(id)
		stp, err := ioutil.ReadFile(bridgeFolder + "/stp_state")
		checkError(err)
		bridge.Stp = stp[0] == '1'
		files, err := ioutil.ReadDir(sysClassNet + name + "/brif")
		if err == nil && len(files) > 0 {
			bridge.Interface = files[0].Name()
		}
		return bridge, nil
	}
	return bridge, errors.New("not a bridge")
}

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
