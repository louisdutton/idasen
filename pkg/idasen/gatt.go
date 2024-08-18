package idasen

import "tinygo.org/x/bluetooth"

func (i *Idasen) read(c bluetooth.DeviceCharacteristic) ([]byte, error) {
	raw := make([]byte, 255)
	_, err := c.Read(raw)
	return raw, err
}

func GetDeviceCharacteristic(d *bluetooth.Device, uuid bluetooth.UUID) (bluetooth.DeviceCharacteristic, error) {
	c := bluetooth.DeviceCharacteristic{}
	services, _ := d.DiscoverServices(nil)
	for _, s := range services {
		characteristics, err := s.DiscoverCharacteristics([]bluetooth.UUID{uuid})
		if err != nil {
			// println(err)
		}
		if len(characteristics) > 0 {
			return characteristics[0], nil
		}
	}
	return c, nil
}

func (i *Idasen) write(uuid bluetooth.UUID, value []byte) error {
	c, err := GetDeviceCharacteristic(i.device, uuid)
	if err != nil {
		return err
	}
	_, err = c.WriteWithoutResponse(value)
	return err
}
