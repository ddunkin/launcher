package launcher

import (
	"testing"
	"time"
)

func TestTurn(t *testing.T) {
	launcher := Create()
	defer launcher.Destroy()
	launcher.SendCommand(Left)
	time.Sleep(200 * time.Millisecond)
	launcher.SendCommand(Stop)
	time.Sleep(500 * time.Millisecond)

	launcher.SendCommand(Right)
	time.Sleep(200 * time.Millisecond)
	launcher.SendCommand(Stop)
	time.Sleep(500 * time.Millisecond)

	launcher.SendCommand(Up)
	time.Sleep(200 * time.Millisecond)
	launcher.SendCommand(Stop)
	time.Sleep(500 * time.Millisecond)

	launcher.SendCommand(Down)
	time.Sleep(200 * time.Millisecond)
	launcher.SendCommand(Stop)
	time.Sleep(500 * time.Millisecond)

}

