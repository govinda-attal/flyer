package sol

import (
	"errors"
)

type SourceDest [2]string

func (sd SourceDest) Values() (src, dst string) {
	return sd[0], sd[1]
}

type Journey []SourceDest

// Overall function returns starting port and final departure port.
// Complexity: O(n) + O(n).
func Overall(journey Journey) (*SourceDest, error) {
	if len(journey) == 0 {
		return nil, ErrInvalidFlightData
	}

	var (
		rs     SourceDest
		counts = map[string]int{}
	)

	for _, leg := range journey {
		src, dst := leg.Values()
		if src == "" || dst == "" {
			return nil, ErrInvalidFlightData
		}

		counts[src] = counts[src] - 1
		counts[dst] = counts[dst] + 1
	}

	for ap, count := range counts {
		switch count {
		case -1:
			if rs[0] != "" {
				return nil, ErrInvalidFlightData
			}
			rs[0] = ap
		case 1:
			if rs[1] != "" {
				return nil, ErrInvalidFlightData
			}
			rs[1] = ap
		case 0:
			continue
		default:
			return nil, ErrInvalidFlightData
		}
	}

	return &rs, nil
}

// Itinerary function returns sorted joruney with clear path.
// Runtime Complexity: O(n) + O(n) + O(n).
func Itinerary(journey Journey) (*Journey, error) {
	if len(journey) == 0 {
		return nil, ErrInvalidFlightData
	}

	destinations := map[string]string{}
	sources := map[string]string{}

	for _, leg := range journey {
		src, dst := leg.Values()

		if _, ok := sources[dst]; ok {
			return nil, ErrInvalidFlightData
		}
		destinations[src] = dst
		sources[dst] = src
	}
	var startLeg SourceDest
	for _, leg := range journey {
		src, _ := leg.Values()
		if _, ok := sources[src]; !ok {
			startLeg = leg
			break
		}
	}
	var (
		itinerary = Journey{startLeg}
		_, next   = startLeg.Values()
	)

	for destinations[next] != "" {
		src := next
		next = destinations[src]
		itinerary = append(itinerary, SourceDest{src, next})
	}

	return &itinerary, nil
}

var (
	ErrInvalidFlightData = errors.New("invalid flights data")
)

const (
	unknown = "unknown"
)
