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

	cmdResponse := "response"

	mockedCmd := NewMockCmd(ctrl)
	mockedCmd.EXPECT().Output().Return([]byte(cmdResponse), nil).Times(1)

	ha := general.HostnameArgs{
		Hostname: "leberKleber-nb",
	}

	m.CommandContext = func(_ context.Context, name string, args ...string) utils.Cmd {
		require.Equal(t, "nmcli", name)
		require.EqualValues(t, []string{"general", "hostname", ha.Hostname}, args)

		return mockedCmd
	}

	output, err := m.Hostname(context.Background(), ha)
	require.NoError(t, err)
	require.Equal(t, cmdResponse, output)
}
