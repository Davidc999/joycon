//+build linux,!nobluez,never

package main

import (
	"github.com/Davidc999//joycon/prog4/bluez"
	"github.com/Davidc999//joycon/prog4/jcpc"
)

func getBluetoothManager() (jcpc.BluetoothManager, error) {
	return bluez.New()
}
