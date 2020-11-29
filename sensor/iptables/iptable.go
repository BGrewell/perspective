package iptables

import (
	"fmt"
	"github.com/BGrewell/perspective/helpers"
)

func AddNatRules() (err error) {
	cmds := []string{
		fmt.Sprintf("iptables -t nat -I POSTROUTING -o eth0 -j MASQUERADE"),
		fmt.Sprintf("iptables -A FORWARD -i eth1 -o eth0 -j ACCEPT"),
		fmt.Sprintf("iptables -A FORWARD -i eth0 -o eth1 -m state --state RELATED,ESTABLISHED -j ACCEPT"),
	}
	_, err = helpers.ExecuteCommands(cmds)
	return err
}

func DelNatRules() (err error) {
	cmds := []string{
		fmt.Sprintf("iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE"),
		fmt.Sprintf("iptables -D FORWARD -i eth1 -o eth0 -j ACCEPT"),
		fmt.Sprintf("iptables -D FORWARD -i eth0 -o eth1 -m state --state RELATED,ESTABLISHED -j ACCEPT"),
	}
	_, err = helpers.ExecuteCommands(cmds)
	return err
}

func AddTProxyRule(proto string, port int, mark int) (err error) {
	//ip, err := helpers.GetPublicIp()
	//if err != nil {
	//	return err
	//}
	//fmt.Println(ip)
	cmd := fmt.Sprintf("iptables -t mangle -A PREROUTING -p %s ! --dport 22 -j TPROXY --on-port %d --on-ip 0.0.0.0 --tproxy-mark 0x%x/0x%x", proto, port, mark, mark)
	_, err = helpers.ExecuteCommand(cmd)
	return err
}

func DelTProxyRule(proto string, port int, mark int) (err error) {
	//ip, err := helpers.GetPublicIp()
	cmd := fmt.Sprintf("iptables -t mangle -D PREROUTING -p %s ! --dport 22 -j TPROXY --on-port %d --on-ip 0.0.0.0 --tproxy-mark 0x%x/0x%x", proto, port, mark, mark)
	_, err = helpers.ExecuteCommand(cmd)
	return err
}
