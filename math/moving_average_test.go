package math_test

import (
	"testing"

	"github.com/elgohr/go-math-tools/math"
	"github.com/stretchr/testify/require"
)

func TestMovingAverage(t *testing.T) {
	input := make(chan float64, 6)
	input <- 1
	input <- 2
	input <- 3
	input <- 4
	input <- 5
	input <- 0
	out := math.MovingAverage(t.Context(), input, 3)
	require.Equal(t, 2.0, <-out)
	require.Equal(t, 3.0, <-out)
	require.Equal(t, 4.0, <-out)
	require.Equal(t, 3.0, <-out)
}
