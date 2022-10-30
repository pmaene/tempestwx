package tempestwx

import (
	"context"
	"encoding/json"
	"net"

	"github.com/pmaene/tempestwx/messages"
)

type TempestWX struct {
	RainStartEvents       chan *messages.RainStartEvent
	LightningStrikeEvents chan *messages.LightningStrikeEvent
	RapidWindMessages     chan *messages.RapidWind
	Observations          chan *messages.Observation
	DeviceStatusMessages  chan *messages.DeviceStatus
	HubStatusMessages     chan *messages.HubStatus
	Errors                chan error

	conn net.PacketConn
}

func (t *TempestWX) Start(ctx context.Context) {
	go t.readMessages(ctx)
}

func (t *TempestWX) readMessages(ctx context.Context) {
	buf := make([]byte, 1024)
	for {
		n, _, err := t.conn.ReadFrom(buf)
		if err != nil {
			t.Errors <- err
		}

		select {
		case <-ctx.Done():
			return

		default:
			m := new(messages.Message)
			if err := json.Unmarshal(buf[:n], m); err != nil {
				t.Errors <- err
			}

			switch m.Type {
			case messages.MessageTypeRainStartEvent:
				m := new(messages.RainStartEvent)
				if err := json.Unmarshal(buf[:n], m); err != nil {
					t.Errors <- err
				}

				t.RainStartEvents <- m

			case messages.MessageTypeLightningStrikeEvent:
				m := new(messages.LightningStrikeEvent)
				if err := json.Unmarshal(buf[:n], m); err != nil {
					t.Errors <- err
				}

				t.LightningStrikeEvents <- m

			case messages.MessageTypeRapidWind:
				m := new(messages.RapidWind)
				if err := json.Unmarshal(buf[:n], m); err != nil {
					t.Errors <- err
				}

				t.RapidWindMessages <- m

			case messages.MessageTypeObservation:
				m := new(messages.Observation)
				if err := json.Unmarshal(buf[:n], m); err != nil {
					t.Errors <- err
				}

				t.Observations <- m

			case messages.MessageTypeDeviceStatus:
				m := new(messages.DeviceStatus)
				if err := json.Unmarshal(buf[:n], m); err != nil {
					t.Errors <- err
				}

				t.DeviceStatusMessages <- m

			case messages.MessageTypeHubStatus:
				m := new(messages.HubStatus)
				if err := json.Unmarshal(buf[:n], m); err != nil {
					t.Errors <- err
				}

				t.HubStatusMessages <- m

			default:
				continue
			}
		}
	}
}

func New(addr string) (*TempestWX, error) {
	c, err := net.ListenPacket("udp4", addr)
	if err != nil {
		return nil, err
	}

	return &TempestWX{
		RainStartEvents:       make(chan *messages.RainStartEvent),
		LightningStrikeEvents: make(chan *messages.LightningStrikeEvent),
		RapidWindMessages:     make(chan *messages.RapidWind),
		Observations:          make(chan *messages.Observation),
		DeviceStatusMessages:  make(chan *messages.DeviceStatus),
		HubStatusMessages:     make(chan *messages.HubStatus),
		Errors:                make(chan error),

		conn: c,
	}, nil
}
