package devices

import "fmt"

const (
	LowBatteryDischarging      = "%s is running out of battery!"
	LowBatteryCharging         = "%s has been plugged in"
	FullyCharged               = "%s has been fully charged. Please remove charger"
	ChargerRemoved             = "%s has been plugged out while partially charged. Current level %d"
	ChargerRemovedFullyCharged = "%s has been plugged out after a full charge"
	DefaultStatus              = "Status cannot be defined for device %s"
)

type Device struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	// True -> Charging || False -> Discharging
	ChargingStatus bool `json:"status"`
}

// Returns the device charging status in a string format
func (device Device) Status() string {
	if device.ChargingStatus {
		return "charging"
	} else {
		return "discharging"
	}
}

// Returns the device status in human readable message
func (device Device) EventMessage() string {
	switch {
	case !device.ChargingStatus && device.Level < 20:
		return fmt.Sprintf(LowBatteryDischarging, device.Name)
	case device.ChargingStatus && device.Level < 100:
		return fmt.Sprintf(LowBatteryCharging, device.Name)
	case device.ChargingStatus && device.Level == 100:
		return fmt.Sprintf(FullyCharged, device.Name)
	case !device.ChargingStatus && device.Level < 100:
		return fmt.Sprintf(ChargerRemoved, device.Name, device.Level)
	case !device.ChargingStatus && device.Level == 100:
		return fmt.Sprintf(ChargerRemovedFullyCharged, device.Name)
	default:
		return fmt.Sprintf(DefaultStatus, device.Name)
	}
}
