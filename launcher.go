package launcher

import (
	"log"
	"github.com/ddunkin/go-libusb"
)

type Launcher struct {
	device *libusb.Device
}

const (
	Fire = 0x10
	Left = 0x4
	Right = 0x8
	Up = 0x2
	Down = 0x1
	Stop = 0x0
)

func Create() *Launcher {
	libusb.Init()

	dev := libusb.Open(0x2123, 0x1010)
	if dev == nil {
		log.Println("Device not found")
		return nil
	}

	launcher := new(Launcher)
	launcher.device = dev
	return launcher
}

func (launcher *Launcher) Destroy() {
	launcher.device.Close()
}

func (launcher *Launcher) SendCommand(command byte) {
	var msg [8]byte
	msg[0] = 0x02
	msg[1] = command

	launcher.device.ControlMsg(0x21, 0x09, 0x0200, 0, msg[:])
}

