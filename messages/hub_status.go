package messages

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidResetFlag = errors.New("messages: invalid reset flag")
)

type ResetFlag int

const (
	ResetFlagUndefined ResetFlag = iota
	ResetFlagBrownoutReset
	ResetFlagPINReset
	ResetFlagPowerReset
	ResetFlagSoftwareReset
	ResetFlagWatchdogReset
	ResetFlagWindowWatchdogReset
	ResetFlagLowPowerReset
	ResetFlagHardFaultDetect
)

func ParseResetFlag(f string) (ResetFlag, error) {
	switch f {
	case "BOR":
		return ResetFlagBrownoutReset, nil
	case "PIN":
		return ResetFlagPINReset, nil
	case "POR":
		return ResetFlagLowPowerReset, nil
	case "SFT":
		return ResetFlagSoftwareReset, nil
	case "WDG":
		return ResetFlagWatchdogReset, nil
	case "WWD":
		return ResetFlagWindowWatchdogReset, nil
	case "LPW":
		return ResetFlagLowPowerReset, nil
	case "HRDFLT":
		return ResetFlagHardFaultDetect, nil

	default:
		return ResetFlagUndefined, ErrInvalidResetFlag
	}
}

type HubStatus struct {
	Message
	Time             time.Time
	Uptime           time.Duration
	FirmwareRevision string
	RadioStats       RadioStats
	ResetFlags       []ResetFlag
	RSSI             float64
}

func (m *HubStatus) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &m.Message); err != nil {
		return err
	}

	type Alias struct {
		Timestamp        int64   `json:"timestamp"`
		Uptime           int64   `json:"uptime"`
		FirmwareRevision string  `json:"firmware_revision"`
		ResetFlags       string  `json:"reset_flags"`
		RSSI             float64 `json:"rssi"`
	}

	a := new(Alias)
	if err := json.Unmarshal(data, a); err != nil {
		return err
	}

	var fs []ResetFlag
	for _, f := range strings.Split(a.ResetFlags, ",") {
		tmp, err := ParseResetFlag(f)
		if err != nil {
			return err
		}

		fs = append(fs, tmp)
	}

	m.Time = time.Unix(a.Timestamp, 0)
	m.Uptime = time.Duration(a.Uptime * int64(time.Second))
	m.FirmwareRevision = a.FirmwareRevision

	if err := json.Unmarshal(data, &m.RadioStats); err != nil {
		return err
	}

	m.ResetFlags = fs
	m.RSSI = a.RSSI

	return nil
}
