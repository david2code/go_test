package main

import (
	"fmt"
	"math/big"
	"net"
)

type route_item struct {
	network        uint32
	network_masked uint32
	mask           uint32 //掩码长度
	eth_name       string
}

type tbl_map_t map[uint32]*route_item
type tbl_list_t []tbl_map_t

var table_list tbl_list_t = make(tbl_list_t, 33)

func inet_aton(ip string) uint32 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return uint32(ret.Uint64())
}
func inet_ntoa(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func add_route_item(item route_item) (err error) {
	if item.mask > 32 {
		err = fmt.Errorf("mask error")
		return
	}
	item.network_masked = item.network >> (32 - item.mask)

	var map_item tbl_map_t = table_list[item.mask]
	if map_item == nil {
		map_item = tbl_map_t{}
		table_list[item.mask] = map_item
	}

	var oldItem *route_item
	oldItem, ok := map_item[item.network_masked]
	if ok {
		err = fmt.Errorf("add failed! eth:%s, network:%s, mask:%d conflict with eth:%s, network:%s, mask:%d", item.eth_name, inet_ntoa(item.network), item.mask, oldItem.eth_name, inet_ntoa(oldItem.network), oldItem.mask)
		return
	}
	map_item[item.network_masked] = &item
	return
}

func rule_match(ip uint32) (item *route_item, err error) {
	for i := 32; i >= 0; i-- {
		var map_item = table_list[i]
		if map_item == nil {
			continue
		}

		ip_masked := ip
		ip_masked >>= (32 - i)
		fmt.Printf("mask: %d, ip: %s, ip_masked: %s\n", i, inet_ntoa(ip), inet_ntoa(ip_masked))

		item, ok := map_item[ip_masked]
		if ok {
			return item, nil
		}
	}
	err = fmt.Errorf("miss match!")
	return
}

var route_table = []route_item{
	{
		inet_aton("192.168.0.0"),
		0,
		16,
		"eth1",
	},

	{
		inet_aton("10.10.12.0"),
		0,
		24,
		"ensp0s1",
	},

	{
		inet_aton("10.10.12.3"),
		0,
		24,
		"ensp0s3",
	},

	{
		inet_aton("1.2.0.0"),
		0,
		8,
		"wan",
	},

	{
		inet_aton("10.10.12.1"),
		0,
		32,
		"all_match_eth",
	},

	{
		inet_aton("127.0.0.1"),
		0,
		0,
		"default",
	},
}
var test_ip_list = []uint32{
	inet_aton("192.168.3.3"),
	inet_aton("192.168.0.3"),
	inet_aton("192.16.6.7"),
	inet_aton("10.10.3.7"),
	inet_aton("10.10.6.6"),
	inet_aton("172.10.3.3"),
}

func test_route_match() {
	for i := 0; i < len(route_table); i++ {
		err := add_route_item(route_table[i])
		if err != nil {
			fmt.Printf("err:%s\n", err)
		} else {
			item := &route_table[i]
			fmt.Printf("add success! eth:%s, network:%s, mask:%d\n", item.eth_name, inet_ntoa(item.network), item.mask)
		}
	}

	for i := 0; i < len(test_ip_list); i++ {
		item, err := rule_match(test_ip_list[i])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("  >%s match! eth:%s, network:%s, mask:%d\n", inet_ntoa(test_ip_list[i]), item.eth_name, inet_ntoa(item.network), item.mask)
		}
	}
}

func main() {
	test_route_match()
	return
}
