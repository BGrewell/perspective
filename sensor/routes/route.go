package routes

import "github.com/BGrewell/perspective/helpers"

func AddTProxyRoute() (err error) {
	cmds := []string{
		"ip rule add fwmark 1 lookup 100",
		"ip route add local 0.0.0.0/0 dev lo table 100",
	}
	_, err = helpers.ExecuteCommands(cmds)
	return err
}

func DelTProxyRoute() (err error) {
	cmds := []string{
		"ip rule del fwmark 1 lookup 100",
		"ip route del local 0.0.0.0/0 dev lo table 100",
	}
	_, err = helpers.ExecuteCommands(cmds)
	return err
}
