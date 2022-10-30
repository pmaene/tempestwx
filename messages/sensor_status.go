package messages

import (
	"encoding/json"
)

const (
	LightningFailedMask        = 0x00001
	LightningNoiseMask         = 0x00002
	LightningDisturberMask     = 0x00004
	PressureFailedMask         = 0x00008
	TemperatureFailedMask      = 0x00010
	RelativeHumidityFailedMask = 0x00020
	WindFailedMask             = 0x00040
	PrecipitationFailedMask    = 0x00080
	LightFailedMask            = 0x00100
)

type SensorStatus struct {
	RelativeHumidityFailed bool
	PressureFailed         bool
	TemperatureFailed      bool
	LightFailed            bool
	LightningDisturber     bool
	LightningFailed        bool
	LightningNoise         bool
	PrecipitationFailed    bool
	WindFailed             bool
}

func (s *SensorStatus) UnmarshalJSON(data []byte) error {
	type Alias struct {
		SensorStatus uint64 `json:"sensor_status"`
	}

	a := new(Alias)
	if err := json.Unmarshal(data, a); err != nil {
		return err
	}

	s.RelativeHumidityFailed = maskSensorStatus(a.SensorStatus, RelativeHumidityFailedMask)
	s.PressureFailed = maskSensorStatus(a.SensorStatus, PressureFailedMask)
	s.TemperatureFailed = maskSensorStatus(a.SensorStatus, TemperatureFailedMask)
	s.LightFailed = maskSensorStatus(a.SensorStatus, LightFailedMask)
	s.LightningFailed = maskSensorStatus(a.SensorStatus, LightningFailedMask)
	s.LightningNoise = maskSensorStatus(a.SensorStatus, LightningNoiseMask)
	s.LightningDisturber = maskSensorStatus(a.SensorStatus, LightningDisturberMask)
	s.PrecipitationFailed = maskSensorStatus(a.SensorStatus, PrecipitationFailedMask)
	s.WindFailed = maskSensorStatus(a.SensorStatus, WindFailedMask)

	return nil
}

func maskSensorStatus(v, m uint64) bool {
	return v&m != 0
}
