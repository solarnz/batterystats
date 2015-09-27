package batterystats

import (
	"time"
)

type BatteryStatus struct {
	Timestamp        time.Time
	BatteryID        string
	CycleCount       uint64
	EnergyFull       uint64
	EnergyFullDesign uint64
	EnergyNow        uint64
	Manufacturer     string
	ModelName        string
	PowerNow         uint64
	SerialNumber     string
	Status           string
	Technology       string
}
