package main

import (
	"encoding/json"
	"io/ioutil"
)

type Device struct {
	Hostname    string `json:"hostname"`
	Ipaddress   string `json:"ipaddress"`
	SnmpVersion int    `json:"snmpVersion"`
	SnmpV1cred  struct {
		Community string `json:"community"`
	} `json:"snmpV1"`
	SnmpV3cred struct {
		SecName      string `json:"secName"`
		AuthPassword string `json:"authPassword"`
		AuthProto    string `json:"authProtocol"`
		PrivPassword string `json:"privPassword"`
		PrivProto    string `json:"priveProtocol"`
		SecLevel     string `json:"secLevel"`
	} `json:"snmpV3"`
}

func ImportDevices(deviceFile string) []Device {
	deviceList := []Device{}
	buffer, err := ioutil.ReadFile(deviceFile)
	if err != nil {
		logger.Fatal(err)
	}

	err = json.Unmarshal(buffer, &deviceList)
	if err != nil {
		logger.Fatal("Failed to load device. ", err)
	}

	return deviceList
}
