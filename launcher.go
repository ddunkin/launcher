package launcher

import (
	"log"
	"time"
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

	const ifc = 0
	r, driver := dev.Driver(ifc)
	if r == 0 {
		log.Printf("Claimed by %s\n", driver)
		if dev.Detach(ifc) != 0 {
			log.Println("Cannot detach")
			return nil
		}
	}

	claimed := dev.Interface(ifc)
	if claimed != 0 {
		log.Println("Cannot claim device")
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

func (launcher *Launcher) SendCommandDuration(command byte, durationMillis int64) {
	launcher.SendCommand(command)
	if durationMillis != 0 {
		time.Sleep(time.Duration(durationMillis * int64(time.Millisecond)))
		launcher.SendCommand(Stop)
	}
}
