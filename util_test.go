package footballdata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_durationToTimeFrame_Zero(t *testing.T) {
	assert.Equal(t, "0", durationToTimeFrame(0))
}

func Test_durationToTimeFrame_Negative7Days(t *testing.T) {
	assert.Equal(t, "p7", durationToTimeFrame(-7*day))
}

func Test_durationToTimeFrame_Positive7Days(t *testing.T) {
	assert.Equal(t, "n7", durationToTimeFrame(7*day))
}

func Test_durationToTimeFrame_PositiveHalfDay(t *testing.T) {
	assert.Equal(t, "n0.5", durationToTimeFrame(12*time.Hour))
}
