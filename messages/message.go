package messages

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidMessageType = errors.New("messages: invalid message type")
)

type MessageType int

const (
	MessageTypeUndefined MessageType = iota
	MessageTypeLightningStrikeEvent
	MessageTypeRainStartEvent
	MessageTypeRapidWind
	MessageTypeObservation
	MessageTypeDeviceStatus
	MessageTypeHubStatus
)

func ParseMessageType(s string) (MessageType, error) {
	switch s {
	case "evt_strike":
		return MessageTypeLightningStrikeEvent, nil
	case "evt_precip":
		return MessageTypeRainStartEvent, nil
	case "rapid_wind":
		return MessageTypeRapidWind, nil
	case "obs_st":
		return MessageTypeObservation, nil

	case "device_status":
		return MessageTypeDeviceStatus, nil
	case "hub_status":
		return MessageTypeHubStatus, nil

	default:
		return MessageTypeUndefined, ErrInvalidMessageType
	}
}

type Message struct {
	Type         MessageType
	SerialNumber string
}

func (m *Message) UnmarshalJSON(data []byte) error {
	type Alias struct {
		Type         string `json:"type"`
		SerialNumber string `json:"serial_number"`
	}

	a := new(Alias)
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	t, err := ParseMessageType(a.Type)
	if err != nil {
		return err
	}

	m.Type = t
	m.SerialNumber = a.SerialNumber

	return nil
}
