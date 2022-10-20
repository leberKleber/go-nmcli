package general_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leberKleber/go-nmcli/general"
	"github.com/leberKleber/go-nmcli/utils"
	"github.com/stretchr/testify/require"
)

func TestManager_Hostname(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := general.Manager{}

	persistedHostname := "leberKleber-nb"
	mockedCmd := NewMockCmd(ctrl)
	mockedCmd.EXPECT().Output().Times(1).Return([]byte(persistedHostname), nil)

	m.CommandContext = func(ctx context.Context, name string, args ...string) utils.Cmd {
		require.Equal(t, "nmcli", name)
		require.EqualValues(t, []string{"general", "hostname"}, args)

		return mockedCmd
	}

	hostname, err := m.Hostname(context.Background())
	require.NoError(t, err)

	require.Equal(t, persistedHostname, hostname)
}

func TestManager_SetHostname(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := general.Manager{}

	mockedCmd := NewMockCmd(ctrl)
	mockedCmd.EXPECT().Run().Times(1)

	newHostname := "leberKleber-nb"

	m.CommandContext = func(ctx context.Context, name string, args ...string) utils.Cmd {
		require.Equal(t, "nmcli", name)
		require.EqualValues(t, []string{"general", "hostname", newHostname}, args)

		return mockedCmd
	}

	err := m.SetHostname(context.Background(), newHostname)
	require.NoError(t, err)
}
