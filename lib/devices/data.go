package devices

import (
	"context"
	"fmt"
	"time"

	db "devices/lib/mongo"
	notify "devices/lib/slack"
)

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
	Timestamp      time.Time
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
	case !device.ChargingStatus && device.Level <= 20:
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

// Adds current timestamp to record
func (device Device) AutoTimestamp() Device {
	device.Timestamp = time.Now()
	return device
}

// Saves current configuration to db
func (device Device) Save() error {
	ctx := context.Background()
	collection := db.MongoCollection("battery")

	_, err := collection.InsertOne(ctx, device)

	return err
}

// Send slack notification for the given device status
func (device Device) Notify() {
	var msg string

	if device.ChargingStatus {
		msg = "🔌 " + device.EventMessage()
	} else {
		msg = "🔋 " + device.EventMessage()
	}
	notify.SendSimpleMessage(msg)
}
