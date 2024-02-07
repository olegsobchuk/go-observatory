package pinger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLatency(t *testing.T) {
	t.Run("Should return difference between finish and start time", func(t *testing.T) {
		timeNow := time.Now()
		result := time.Duration(2 * time.Second)
		pinger := Pinger{StartedAt: timeNow, FinishedAt: timeNow.Add(2 * time.Second)}
		assert.Equal(t, pinger.Latency(), result)
	})
}

func TestIsStarted(t *testing.T) {
	t.Run("Should return true on changed time", func(t *testing.T) {
		pinger := Pinger{StartedAt: time.Now()}
		assert.Equal(t, pinger.IsStarted(), true)
	})

	t.Run("Should return false on not changed time", func(t *testing.T) {
		pinger := Pinger{}
		assert.Equal(t, pinger.IsStarted(), false)
	})
}

func TestIsFinished(t *testing.T) {
	t.Run("Should return true on changed time", func(t *testing.T) {
		pinger := Pinger{StartedAt: time.Now(), FinishedAt: time.Now()}
		assert.Equal(t, pinger.IsFinished(), true)
	})

	t.Run("Should return false on not changed start time", func(t *testing.T) {
		pinger := Pinger{FinishedAt: time.Now()}
		assert.Equal(t, pinger.IsFinished(), false)
	})

	t.Run("Should return false on not changed time", func(t *testing.T) {
		pinger := Pinger{}
		assert.Equal(t, pinger.IsFinished(), false)
	})
}
