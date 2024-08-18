package idasen

import (
	"tinygo.org/x/bluetooth"
)

const (
	MinHeight   float64 = 0.62  // Minimum desk height in meters.
	MaxHeight   float64 = 1.27  // Maximum desk height in meters.
	MinMovement float64 = 0.005 // Minimum movement interval in meters.

	CommandDown byte = 0x46
	CommandUp   byte = 0x47
	CommandWake byte = 0xFE
	CommandStop byte = 0xFF

	ReferenceInputStop byte = 0x01
)

var (
	Adapter = bluetooth.DefaultAdapter
	// spec: https://github.com/anson-vandoren/linak-desk-spec/blob/master/dpg_commands.md

	// Control
	ControlService               = uuid(0x01)
	ControlCommandCharacteristic = uuid(0x02)
	ControlErrorCharacteristic   = uuid(0x03)

	// DPG
	DPGService        = uuid(0x10)
	DPGCharacteristic = uuid(0x11)
	DPGGetSetup       = 0x0a
	DPGCurrentTime    = 0x16
	DPGFactoryReset   = 0x0b
	// TODO: finish rest

	// Reference Output
	ReferenceOutputService                  = uuid(0x20)
	ReferenceOutputOneCharacteristic        = uuid(0x21) // Height
	ReferenceOutputTwoCharacteristic        = uuid(0x22)
	ReferenceOutputThreeCharacteristic      = uuid(0x23)
	ReferenceOutputFourCharacteristic       = uuid(0x24)
	ReferenceOutputFiveCharacteristic       = uuid(0x25)
	ReferenceOutputSixCharacteristic        = uuid(0x26)
	ReferenceOutputSevenCharacteristic      = uuid(0x27)
	ReferenceOutputEightCharacteristic      = uuid(0x28)
	ReferenceOutputMaskCharacteristic       = uuid(0x29)
	ReferenceOutputDetectMaskCharacteristic = uuid(0x2a)

	ReferenceInputService = uuid(0x30)
	ReferenceInputOne     = uuid(0x31)
	ReferenceInputTwo     = uuid(0x32)
	ReferenceInputThree   = uuid(0x33)
	ReferenceInputFour    = uuid(0x34)

	CmdReferenceInputStop = []byte{0x01, 0x80}
)

func uuid(b byte) bluetooth.UUID {
	return bluetooth.NewUUID([16]byte{0x99, 0xfa, 0x00, b, 0x33, 0x8a, 0x10, 0x24, 0x8a, 0x49, 0x00, 0x9c, 0x02, 0x15, 0xf7, 0x8a})
}
