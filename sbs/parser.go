package sbs

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Message is one decoded SBS message from dump1090
//
// We have pointers to most fields except the hex code which always has a value, because the messages
// arrive in random fashion. We do not want miscommunications where a not supplied int (for example, altitude)
// defaults to 0 which gives an impression of plane on ground which it isn't. We simply need to wait till
// we get the data. Pointers solve this easily
type Message struct {
	CapturedAt    time.Time
	Hex           string
	Altitude      *int
	Speed         *int
	Heading       *int
	Callsign      *string
	Latitude      *float64
	Longitude     *float64
	VerticalSpeed *int
	Squawk        *int
}

func Parse(line string, receivedAt time.Time) (Message, error) {
	splitLines := strings.SplitN(strings.TrimSpace(line), ",", -1)

	if len(splitLines) != 22 {
		return Message{}, fmt.Errorf("sbs: want 22 fields, got %d", len(splitLines))
	}

	if splitLines[0] != "MSG" {
		return Message{}, fmt.Errorf("sbs: Didn't receive a MSG type from dump1090")
	}

	// Hex is emitted every message so it could never be null
	m := Message{CapturedAt: receivedAt, Hex: splitLines[4]}

	// Parse all other fields whose values arrive at irregular intervals

	alt, err := optInt(splitLines[11])
	if err != nil {
		return Message{}, fmt.Errorf("sbs: altitude: %w", err)
	}
	m.Altitude = alt

	speed, err := optInt(splitLines[12])
	if err != nil {
		return Message{}, fmt.Errorf("sbs: speed: %w", err)
	}
	m.Speed = speed

	heading, err := optInt(splitLines[13])
	if err != nil {
		return Message{}, fmt.Errorf("sbs: heading: %w", err)
	}
	m.Heading = heading

	if callsign := strings.TrimSpace(splitLines[10]); callsign != "" {
		m.Callsign = &callsign
	}

	lat, err := optFloat64(splitLines[14])
	if err != nil {
		return Message{}, fmt.Errorf("sbs: latitude: %w", err)
	}
	m.Latitude = lat

	long, err := optFloat64(splitLines[15])
	if err != nil {
		return Message{}, fmt.Errorf("sbs: longitude: %w", err)
	}
	m.Longitude = long

	// TODO Do not parse squawk code as int as it would remove the leading 0. Squawk 0602 -> 602
	// Squawk is always 4 digits, on panel on the aircraft and also in ATC communications
	sq, err := optInt(splitLines[17])
	if err != nil {
		return Message{}, fmt.Errorf("sbs: squawk: %w", err)
	}
	m.Squawk = sq

	vs, err := optInt(splitLines[16])
	if err != nil {
		return Message{}, fmt.Errorf("sbs: vertical speed: %w", err)
	}
	m.VerticalSpeed = vs

	return m, nil

}

func optInt(s string) (*int, error) {
	if s == "" {
		return nil, nil
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return nil, fmt.Errorf("parse int %q: %w", s, err)
	}
	return &n, nil
}

func optFloat64(s string) (*float64, error) {
	if s == "" {
		return nil, nil
	}
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, fmt.Errorf("parse float %q: %w", s, err)
	}
	return &n, nil
}

func fmtOptInt(p *int) string {
	if p == nil {
		return "nil"
	}
	return strconv.Itoa(*p)
}

func fmtOptStr(s *string) string {
	if s == nil {
		return "nil"
	}

	return *s
}

func fmtOptFloat64(p *float64) string {
	if p == nil {
		return "nil"
	}
	return strconv.FormatFloat(*p, 'f', 5, 64)
}

func (m Message) String() string {
	return fmt.Sprintf("{CapturedAt: %s, Callsign: %s, Hex: %s, Altitude: %s, Speed: %s, Heading: %s, Lat: %s, Lon: %s, Vertical Speed: %s,Squawk: %s}",
		m.CapturedAt.Format("02.01.2006 15:04:05,000Z0700"),
		fmtOptStr(m.Callsign),
		m.Hex,
		fmtOptInt(m.Altitude),
		fmtOptInt(m.Speed),
		fmtOptInt(m.Heading),
		fmtOptFloat64(m.Latitude),
		fmtOptFloat64(m.Longitude),
		fmtOptInt(m.VerticalSpeed),
		fmtOptInt(m.Squawk),
	)
}
