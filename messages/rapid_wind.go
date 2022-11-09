package messages

import (
	"encoding/json"
	"time"
)

type RapidWind struct {
	StationMessage
	Timestamp time.Time
	Direction float64
	Speed     float64
}

func (m *RapidWind) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &m.StationMessage); err != nil {
		return err
	}

	type Alias struct {
		Observations []float64 `json:"ob"`
	}

	a := new(Alias)
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	m.Timestamp = time.Unix(int64(a.Observations[0]), 0)
	m.Direction = a.Observations[2]
	m.Speed = a.Observations[1]

	return nil
}
