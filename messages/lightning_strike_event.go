package messages

import (
	"encoding/json"
	"time"
)

type LightningStrikeEvent struct {
	StationMessage
	Timestamp time.Time
	Distance  float64
	Energy    float64
}

func (m *LightningStrikeEvent) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &m.StationMessage); err != nil {
		return err
	}

	type Alias struct {
		Event []float64 `json:"evt"`
	}

	a := new(Alias)
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	m.Timestamp = time.Unix(int64(a.Event[0]), 0)
	m.Distance = a.Event[1] * 1000
	m.Energy = a.Event[2]

	return nil
}
