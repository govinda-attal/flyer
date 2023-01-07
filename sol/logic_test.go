package sol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverall(t *testing.T) {
	type test struct {
		journey Journey
		want    *SourceDest
		err     error
	}
	testcases := []test{
		{
			journey: Journey{{"SFO", "EWR"}},
			want:    &SourceDest{"SFO", "EWR"},
			err:     nil,
		},
		{
			journey: Journey{{"ATL", "EWR"}, {"SFO", "ATL"}},
			want:    &SourceDest{"SFO", "EWR"},
			err:     nil,
		},
		{
			journey: Journey{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
			want:    &SourceDest{"SFO", "EWR"},
			err:     nil,
		},
	}
	assert := assert.New(t)

	for i, tc := range testcases {
		sd, err := Overall(tc.journey)
		assert.ErrorIs(err, tc.err, "case: %d", i)
		assert.Equal(tc.want, sd, "case: %d", i)
	}
}

func TestItinerary(t *testing.T) {
	type test struct {
		journey Journey
		want    *Journey
		err     error
	}
	testcases := []test{
		{
			journey: Journey{{"SFO", "EWR"}},
			want:    &Journey{{"SFO", "EWR"}},
			err:     nil,
		},
		{
			journey: Journey{{"ATL", "EWR"}, {"SFO", "ATL"}},
			want:    &Journey{{"SFO", "ATL"}, {"ATL", "EWR"}},
			err:     nil,
		},
		{
			journey: Journey{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
			want:    &Journey{{"SFO", "ATL"}, {"ATL", "GSO"}, {"GSO", "IND"}, {"IND", "EWR"}},
			err:     nil,
		},
	}
	assert := assert.New(t)

	for i, tc := range testcases {
		it, err := Itinerary(tc.journey)
		assert.ErrorIs(err, tc.err, "case: %d", i)
		assert.Equal(tc.want, it, "case: %d", i)
	}
}
