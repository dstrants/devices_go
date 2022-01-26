package devices

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestEventMessage(t *testing.T) {
	name := "Test Device"
	now := time.Now()
	level4 := rand.Intn(79) + 20

	devices := []Device{
		{name, 18, false, now},
		{name, rand.Intn(79) + 20, true, now},
		{name, 100, true, now},
		{name, level4, false, now},
		{name, 100, false, now},
		{name, 123, true, now},
	}

	msgs := []string{
		"Test Device is running out of battery!",
		"Test Device has been plugged in",
		"Test Device has been fully charged. Please remove charger",
		"Test Device has been plugged out while partially charged. Current level " + fmt.Sprint(level4),
		"Test Device has been plugged out after a full charge",
		"Status cannot be defined for device Test Device",
	}

	for i, device := range devices {
		if device.EventMessage() != msgs[i] {
			t.Errorf("Event message is not the expected one for device %v", device)
		}
	}
}

func TestDeviceStatus(t *testing.T) {
	name := "Test Device"
	now := time.Now()
	devices := []Device{
		{name, 18, false, now},
		{name, rand.Intn(79) + 20, true, now},
	}

	statuses := []string{"discharging", "charging"}

	for i, device := range devices {
		if device.Status() != statuses[i] {
			t.Errorf("Device status is the expected on for %v", device)
		}
	}
}
