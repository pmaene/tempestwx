package messages

import "encoding/json"

type StationMessage struct {
	Message
	HubSerialNumber string
}

func (m *StationMessage) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &m.Message); err != nil {
		return err
	}

	type Alias struct {
		HubSerialNumber string `json:"hub_sn"`
	}

	a := new(Alias)
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	m.HubSerialNumber = a.HubSerialNumber
	return nil
}
