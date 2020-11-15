// +build linux

package main

import (
	"fmt"
	"errors"
	"strings"
	"regexp"
	"io/ioutil"

	"github.com/Davidc999/joycon/prog4/jcpc"
	"github.com/Davidc999/joycon/prog4/output"
)

func getOutputFactory() jcpc.OutputFactory {
	return func(t jcpc.JoyConType, playerNum int, remap jcpc.InputRemappingOptions) (jcpc.Output, error) {
		switch t {
		case jcpc.TypeLeft:
			return output.NewUInput(output.MappingL, fmt.Sprintf("Half Joy-Con %d", playerNum), remap)
		case jcpc.TypeRight:
			return output.NewUInput(output.MappingR, fmt.Sprintf("Half Joy-Con %d", playerNum), remap)
		case jcpc.TypeBoth:
			return output.NewUInput(output.MappingDual, fmt.Sprintf("Full Joy-Con %d", playerNum), remap)
		}
		panic("bad joycon type")
	}
}


func deleteEventNode(jc jcpc.JoyCon) error {
    dat, err := ioutil.ReadFile("/proc/bus/input/devices")
    if err != nil {
        return err
    }
	file_cont := string(dat)

	found_ind := strings.Index(file_cont, jc.Serial())
	if found_ind == -1 {
	    return errors.New("Node not found for the controller")
    }
	re := regexp.MustCompile(`event\d+`)
	event_used := re.Find([]byte(file_cont[found_ind:]))
	fmt.Printf("Must delete %q for serial %q\n", event_used, jc.Serial())
	return nil
}