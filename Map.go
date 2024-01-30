package main

import (
	"fmt"
	"strings"
)

func main() {

	var a = map[string][3]string{"SR": {"7750", "7450", "7950"},
		"IP-Edge": {"7210", "7705", "WBX"}}
	//fmt.Printf("a\t%v\n", a)

	//a = map[string][3]string{"IP-Edge": {"7210", "7705", "WBX"}}
	fmt.Printf("a\t%v\n", a)

	//a = map ["NOKIA" : {"SR-MD", "SR-L", "SR-M"}}

	var b = map[string]string{"SR": "7750,7450,7950"}

	b["IP-Edge"] = "7210,7705"
	fmt.Printf("a\t%v\n", b)

	b["OMNI"] = "WBX,DC-TOR,MD-SR"

	fmt.Printf("a\t%v\n", b)

	var node string = b["SR"]

	var nodes []string = strings.Split(node, ",")
	var exist bool = false
	for j := 0; j < len(nodes); j++ {
		if nodes[j] == "7250" {
			fmt.Println("Node 7250 is not present in the SR family")
			exist = true
		}

		if !exist {
			var result string = ""
			result = node + ",7250"
			b["SR"] = result
		}

	}
	fmt.Printf("a\t%v\n", b)

	var node2 string = b["OMNI"]

	var nodes2 []string = strings.Split(node2, ",")
	//var exist2 bool = false
	var result2 string = ""
	for j := 0; j < len(nodes2); j++ {
		if nodes2[j] == "WBX" {
			fmt.Println("Node WBX is  present in the OMNI family")
			//exist2 = true
			nodes2[j] = ""

		}
		result2 = result2 + nodes2[j] + ","

	}
	b["OMNI"] = result2
	fmt.Printf("a\t%v\n", b)
}
