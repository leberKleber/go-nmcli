package general_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leberKleber/go-nmcli/general"
	"github.com/leberKleber/go-nmcli/utils"
	"github.com/stretchr/testify/require"
)

func TestManager_Permissions(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := general.Manager{}

	mockedCmd := NewMockCmd(ctrl)
	mockedCmd.EXPECT().Output().Times(1).Return([]byte(`org.freedesktop.NetworkManager.checkpoint-rollback:auth
org.freedesktop.NetworkManager.enable-disable-connectivity-check:yes
org.freedesktop.NetworkManager.sleep-wake:no`), nil)

	m.CommandContext = func(ctx context.Context, name string, args ...string) utils.Cmd {
		require.Equal(t, "nmcli", name)
		require.EqualValues(t, []string{"-g", "PERMISSION,VALUE", "general", "permissions"}, args)

		return mockedCmd
	}

	permissions, err := m.Permissions(context.Background())
	require.NoError(t, err)

	require.EqualValues(t, []general.Permission{
		{Permission: "org.freedesktop.NetworkManager.checkpoint-rollback", Value: general.PermissionValueAuth},
		{Permission: "org.freedesktop.NetworkManager.enable-disable-connectivity-check", Value: general.PermissionValueYes},
		{Permission: "org.freedesktop.NetworkManager.sleep-wake", Value: general.PermissionValueNo},
	}, permissions)
}
