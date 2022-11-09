package messages

import (
	"errors"
	"time"
)

var (
	ErrInvalidPrecipitationType = errors.New("messages: invalid precipitation type")
)

type PrecipitationType int

const (
	PrecipitationTypeUndefined = iota
	PrecipitationTypeNone
	PrecipitationTypeRain
	PrecipitationTypeHail
	PrecipitationTypeBoth
)

func ParsePrecipitationType(t int) (PrecipitationType, error) {
	switch t {
	case 0:
		return PrecipitationTypeNone, nil
	case 1:
		return PrecipitationTypeRain, nil
	case 2:
		return PrecipitationTypeHail, nil
	case 3:
		return PrecipitationTypeBoth, nil

	default:
		return PrecipitationTypeUndefined, ErrInvalidPrecipitationType
	}
}

type Report struct {
	Timestamp                      time.Time
	RelativeHumidity               float64
	StationPressure                float64
	AirTemperature                 float64
	Illuminance                    float64
	SolarRadiation                 float64
	UVIndex                        float64
	LightningStrikeAverageDistance float64
	LightningStrikeCount           float64
	PrecipitationAmount            float64
	PrecipitationType              PrecipitationType
	WindDirection                  float64
	WindSpeedAverage               float64
	WindSpeedGust                  float64
	WindSpeedLull                  float64
	WindSampleInterval             float64
	BatteryVoltage                 float64
	Interval                       float64
}

func NewReport(r []float64) (Report, error) {
	t, err := ParsePrecipitationType(int(r[13]))
	if err != nil {
		return Report{}, err
	}

	return Report{
		Timestamp:                      time.Unix(int64(r[0]), 0),
		RelativeHumidity:               r[8],
		StationPressure:                r[6] / 1000,
		AirTemperature:                 r[7],
		Illuminance:                    r[9],
		SolarRadiation:                 r[11],
		UVIndex:                        r[10],
		LightningStrikeAverageDistance: r[14] * 1000,
		LightningStrikeCount:           r[15],
		PrecipitationAmount:            r[12] / 1000,
		PrecipitationType:              t,
		WindDirection:                  r[4],
		WindSpeedAverage:               r[2],
		WindSpeedGust:                  r[3],
		WindSpeedLull:                  r[1],
		WindSampleInterval:             r[5],
		BatteryVoltage:                 r[16],
		Interval:                       r[17],
	}, nil
}
