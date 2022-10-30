package messages

import (
	"encoding/json"
	"time"
)

type RainStartEvent struct {
	StationMessage
	Time time.Time
}

func (m *RainStartEvent) UnmarshalJSON(data []byte) error {
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

	m.Time = time.Unix(int64(a.Event[0]), 0)
	return nil
}
