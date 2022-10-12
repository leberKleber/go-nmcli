package go_nmcli

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNMCli_GeneralPermissions(t *testing.T) {
	cli := NewNMCli()
	cli.CommandContext = func(ctx context.Context, name string, args ...string) Cmd {
		require.Equal(t, "nmcli", name)
		require.EqualValues(t, []string{"-g", "PERMISSION,VALUE", "general", "permissions"}, args)

		return testCmd{
			output: []byte(`org.freedesktop.NetworkManager.checkpoint-rollback:auth
org.freedesktop.NetworkManager.enable-disable-connectivity-check:yes
org.freedesktop.NetworkManager.sleep-wake:no`),
		}
	}

	permissions, err := cli.GeneralPermissions(context.Background())
	require.NoError(t, err)

	require.EqualValues(t, []GeneralPermission{
		{Permission: "org.freedesktop.NetworkManager.checkpoint-rollback", Value: GeneralPermissionValueAuth},
		{Permission: "org.freedesktop.NetworkManager.enable-disable-connectivity-check", Value: GeneralPermissionValueYes},
		{Permission: "org.freedesktop.NetworkManager.sleep-wake", Value: GeneralPermissionValueNo},
	}, permissions)
}
