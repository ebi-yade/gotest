package cases

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type ErrorCheck func(*testing.T, error)

func NoError(t *testing.T, err error) {
	require.NoError(t, err)
}

func ErrorIs(target error) ErrorCheck {
	return func(t *testing.T, err error) {
		require.ErrorIs(t, err, target)
	}
}

func ErrorAs(target error) ErrorCheck {
	return func(t *testing.T, err error) {
		require.ErrorAs(t, err, target)
	}
}
