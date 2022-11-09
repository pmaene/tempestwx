package messages

import (
	"encoding/json"
	"strconv"
	"time"
)

type DeviceStatus struct {
	StationMessage
	Timestamp        time.Time
	Uptime           time.Duration
	BatteryVoltage   float64
	FirmwareRevision string
	RSSI             float64
	HubRSSI          float64
	SensorStatus     SensorStatus
	Debug            bool
}

func (s *DeviceStatus) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &s.StationMessage); err != nil {
		return err
	}

	type Alias struct {
		Timestamp        int64   `json:"timestamp"`
		Uptime           int64   `json:"uptime"`
		BatteryVoltage   float64 `json:"voltage"`
		FirmwareRevision int64   `json:"firmware_revision"`
		RSSI             float64 `json:"rssi"`
		HubRSSI          float64 `json:"hub_rssi"`
		Debug            int64   `json:"debug"`
	}

	a := new(Alias)
	if err := json.Unmarshal(data, a); err != nil {
		return err
	}

	s.Timestamp = time.Unix(a.Timestamp, 0)
	s.Uptime = time.Duration(a.Uptime * int64(time.Second))
	s.BatteryVoltage = a.BatteryVoltage
	s.FirmwareRevision = strconv.Itoa(int(a.FirmwareRevision))
	s.RSSI = a.RSSI
	s.HubRSSI = a.HubRSSI

	if err := json.Unmarshal(data, &s.SensorStatus); err != nil {
		return err
	}

	s.Debug = a.Debug == 1
	return nil
}
