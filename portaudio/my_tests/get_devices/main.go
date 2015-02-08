package main

import (
	"code.google.com/p/portaudio-go/portaudio"
	"fmt"
)

func main() {

	var err = portaudio.Initialize()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer portaudio.Terminate()

	//-----------------------------
	ExampleDevices()
	ExampleDefaultInputDevice()
	ExampleDefaultOutputDevice()
}

func str_DeviceInfo(di *portaudio.DeviceInfo) string {

	return fmt.Sprintf("%v", *di)
}

func ExampleDevices() {

	ds, err := portaudio.Devices()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("devices:")

	for i, di := range ds {
		fmt.Printf("d[%d]: %s\n", i, str_DeviceInfo(di))
	}
}

func ExampleDefaultInputDevice() error {

	di, err := portaudio.DefaultInputDevice()
	if err != nil {
		return err
	}

	fmt.Printf("DefaultInputDevice: %s\n", str_DeviceInfo(di))

	return nil
}

func ExampleDefaultOutputDevice() error {
	di, err := portaudio.DefaultOutputDevice()
	if err != nil {
		return err
	}

	fmt.Printf("DefaultOutputDevice: %s\n", str_DeviceInfo(di))

	return nil
}
