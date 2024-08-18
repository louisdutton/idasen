package idasen

import (
	"fmt"

	"tinygo.org/x/bluetooth"
)

type Idasen struct {
	device *bluetooth.Device
	height bluetooth.DeviceCharacteristic
}

type desk struct {
	Name    string
	Address string
}

func New(macAddress string, timeout int64) (*Idasen, error) {
	desk, err := GetDesk(macAddress, timeout)
	if err != nil {
		return nil, fmt.Errorf("Could not find device %s: %s", macAddress, err)
	}

	d, err := Adapter.Connect(desk.Address, bluetooth.ConnectionParams{})
	if err != nil {
		return nil, fmt.Errorf("Could not connect to device %s: %s", macAddress, err)
	}

	dref := &d

	heightCharacteristic, err := GetDeviceCharacteristic(dref, ReferenceOutputOneCharacteristic)
	if err != nil {
		return nil, fmt.Errorf("Could not get height characteristic: %s", err)
	}

	return &Idasen{device: dref, height: heightCharacteristic}, nil
}

func (i *Idasen) Close() {
	defer i.device.Disconnect()
}
