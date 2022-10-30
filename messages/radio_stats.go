package messages

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidRadioStatus = errors.New("messages: invalid radio status")
)

type RadioStatus int

const (
	RadioStatusUndefined = iota
	RadioStatusOff
	RadioStatusOn
	RadioStatusActive
	RadioStatusBLEConnected
)

func ParseRadioStatus(s int) (RadioStatus, error) {
	switch s {
	case 0:
		return RadioStatusOff, nil
	case 1:
		return RadioStatusOn, nil
	case 3:
		return RadioStatusActive, nil
	case 7:
		return RadioStatusBLEConnected, nil

	default:
		return RadioStatusUndefined, ErrInvalidRadioStatus
	}
}

type RadioStats struct {
	Version          int
	RebootCount      int
	I2CBusErrorCount int
	Status           RadioStatus
	NetworkID        int
}

func (s *RadioStats) UnmarshalJSON(data []byte) error {
	type Alias struct {
		RadioStats []int64 `json:"radio_stats"`
	}

	a := new(Alias)
	if err := json.Unmarshal(data, a); err != nil {
		return err
	}

	sts, err := ParseRadioStatus(int(a.RadioStats[3]))
	if err != nil {
		return err
	}

	s.Version = int(a.RadioStats[0])
	s.RebootCount = int(a.RadioStats[1])
	s.I2CBusErrorCount = int(a.RadioStats[2])
	s.Status = sts
	s.NetworkID = int(a.RadioStats[4])

	return nil
}
