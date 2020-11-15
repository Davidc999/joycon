// +build !linux

package main

import (
    "errors"

	"github.com/riking/joycon/prog4/jcpc"
	"github.com/riking/joycon/prog4/output"
)

func getOutputFactory() jcpc.OutputFactory {
	return func(t jcpc.JoyConType, playerNum int, remap InputRemappingOptions) (jcpc.Output, error) {
		return output.NewConsole(t, playerNum)
	}
}

func deleteEventNode(jc jcpc.JoyCon) error {
    return errors.New("Not implemented!")
}
