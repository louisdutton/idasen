package idasen

import (
	"encoding/binary"
	"fmt"
)

func assertInBounds(value float64, min float64, max float64) error {
	if value < min || value > max {
		return fmt.Errorf("value %.2f out of bounds [%.2f, %.2f]", value, min, max)
	}
	return nil
}

func rawToMeters(b []byte) float64 {
	if len(b) < 2 {
		return -1
	}

	raw := binary.LittleEndian.Uint16(b[0:2])
	return float64(float64(raw) / 10000)
}

func metersToRaw(m float64) []byte {
	raw := uint16((m) * 10000)
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, raw)
	return b
}
