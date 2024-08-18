package idasen

import (
	"fmt"
	"math"
)

func (i *Idasen) command(cmd byte) error {
	return i.write(ControlCommandCharacteristic, []byte{cmd, 0x00})
}

func (i *Idasen) referenceInput(cmd byte) error {
	return i.write(ReferenceInputOne, []byte{cmd, 0x80})
}

func (i *Idasen) Ascend() error {
	return i.command(CommandUp)
}

func (i *Idasen) Descend() error {
	return i.command(CommandDown)
}

func (i *Idasen) Stop() error {
	err := i.command(CommandStop)
	if err != nil {
		return err
	}
	return i.referenceInput(ReferenceInputStop)
}

func (i *Idasen) SetHeight(targetHeight float64) error {
	return i.SetHeightWithUpdateChannel(targetHeight, nil)
}

// Sets the desk height to the target height (in meters).
func (i *Idasen) SetHeightWithUpdateChannel(targetHeight float64, ch chan float64) error {
	assertInBounds(targetHeight, MinHeight, MaxHeight)

	previousHeight, err := i.Height()
	if err != nil {
		return err
	}

	isDescend := math.Signbit(targetHeight - previousHeight)

	for true {
		currentHeight, err := i.Height()
		if err != nil {
			return err
		}

		if ch != nil {
			ch <- currentHeight
		}

		if math.Signbit(currentHeight-previousHeight) != isDescend {
			_ = i.Stop()
			return fmt.Errorf("safety stop was trigged")
		}
		previousHeight = currentHeight

		if math.Abs(targetHeight-currentHeight) < MinMovement {
			return i.Stop()
		} else if isDescend {
			return i.Descend()
		} else {
			return i.Ascend()
		}

	}

	return nil
}
