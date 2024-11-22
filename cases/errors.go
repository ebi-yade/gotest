package cases

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type ErrorCheck func(*testing.T, error)

func NoError(t *testing.T, err error) {
	t.Helper()
	require.NoError(t, err)
}

func Error(t *testing.T, err error) {
	t.Helper()
	require.Error(t, err)
}

func ErrorIs(target error) ErrorCheck {
	return func(t *testing.T, err error) {
		t.Helper()
		require.ErrorIs(t, err, target)
	}
}

func ErrorAs(target error) ErrorCheck {
	return func(t *testing.T, err error) {
		t.Helper()
		require.ErrorAs(t, err, target)
	}
}
