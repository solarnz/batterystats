// +build linux

package batterystats

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const BATTERY_PATH = "/sys/class/power_supply"

func GetBatteries() ([]BatteryStatus, error) {
	dirs, err := ioutil.ReadDir(BATTERY_PATH)
	if err != nil {
		return nil, err
	}

	batteries := make([]BatteryStatus, 0)
	for _, dir := range dirs {
		if strings.ToLower(dir.Name()) == "ac" {
			continue
		}

		battery, err := readBattery(dir.Name())
		if err != nil {
			return nil, err
		}
		batteries = append(batteries, *battery)
	}

	return batteries, nil
}

func readBattery(name string) (*BatteryStatus, error) {
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/uevent", BATTERY_PATH, name))
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(file), "\n")
	vars := make(map[string]string)

	for _, line := range lines {
		data := strings.Split(line, "=")
		if len(data) == 2 {
			vars[data[0]] = data[1]
		}
	}

	battery := &BatteryStatus{}
	battery.Timestamp = time.Now()
	battery.BatteryID = name
	battery.CycleCount, _ = strconv.ParseUint(vars["POWER_SUPPLY_CYCLE_COUNT"], 10, 64)
	battery.EnergyFull, _ = strconv.ParseUint(vars["POWER_SUPPLY_ENERGY_FULL"], 10, 64)
	battery.EnergyFullDesign, _ = strconv.ParseUint(vars["POWER_SUPPLY_ENERGY_FULL_DESIGN"], 10, 64)
	battery.EnergyNow, _ = strconv.ParseUint(vars["POWER_SUPPLY_ENERGY_NOW"], 10, 64)
	battery.Manufacturer = vars["POWER_SUPPLY_MANUFACTURER"]
	battery.ModelName = vars["POWER_SUPPLY_MODEL_NAME"]
	battery.PowerNow, _ = strconv.ParseUint(vars["POWER_SUPPLY_POWER_NOW"], 10, 64)
	battery.SerialNumber = vars["POWER_SUPPLY_SERIAL_NUMBER"]
	battery.Status = vars["POWER_SUPPLY_STATUS"]
	battery.Technology = vars["POWER_SUPPLY_TECHNOLOGY"]

	return battery, nil
}
