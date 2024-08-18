package idasen

import (
	"fmt"
)

// Returns the current height of the desk relative to the ground.
func (i *Idasen) Height() (float64, error) {
	raw, err := i.read(i.height)
	if err != nil {
		return 0, fmt.Errorf("Failed to read height: %s", err)
	}

	return rawToMeters(raw) + MinHeight, nil
}
