package messages

import (
	"encoding/json"
	"strconv"
)

type Observation struct {
	StationMessage
	FirmwareRevision string
	Reports          []Report
}

func (o *Observation) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.StationMessage); err != nil {
		return err
	}

	type Alias struct {
		Reports          [][]float64 `json:"obs"`
		FirmwareRevision int64       `json:"firmware_revision"`
	}

	a := new(Alias)
	if err := json.Unmarshal(data, a); err != nil {
		return err
	}

	var rs []Report
	for _, r := range a.Reports {
		tmp, err := NewReport(r)
		if err != nil {
			return err
		}

		rs = append(rs, tmp)
	}

	o.FirmwareRevision = strconv.Itoa(int(a.FirmwareRevision))
	o.Reports = rs

	return nil
}
